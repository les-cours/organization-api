package service

import (
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/env"
	"github.com/les-cours/organization-api/graph"
	"github.com/les-cours/organization-api/graph/resolvers"
	"github.com/les-cours/organization-api/types"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registry       = prometheus.NewRegistry()
	requestCounter = prometheus.NewGauge(prometheus.GaugeOpts{Name: "request_counter", Help: "request counter"})
	memoryUsage    = prometheus.NewGauge(prometheus.GaugeOpts{Name: "memory_usage", Help: "memory usage"})
	goRoutineNum   = prometheus.NewGauge(prometheus.GaugeOpts{Name: "go_routines_num", Help: "the number of go routine "})
)

func monitoringMiddleware(originalHandler http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memoryUsage.Set(float64(m.Alloc))
		goRoutineNum.Set(float64(runtime.NumGoroutine()))
		originalHandler.ServeHTTP(w, r)
	})
}
func getUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := types.UserToken{}

		log.Printf("acccountId %s %s %s", r.Header.Get("Accountid"), r.Header.Get("Id"), r.Header.Get("Email"))

		accountID := r.Header.Get("Accountid")
		id := r.Header.Get("Id")
		email := r.Header.Get("Email")

		json.Unmarshal([]byte(r.Header.Get("Create")), &user.Create)
		json.Unmarshal([]byte(r.Header.Get("Read")), &user.Read)
		json.Unmarshal([]byte(r.Header.Get("Update")), &user.Update)
		json.Unmarshal([]byte(r.Header.Get("Delete")), &user.Delete)
		user.AccountID = accountID
		user.ID = id
		user.Email = email

		ctx := context.WithValue(r.Context(), "responseWriter", w)
		ctx = context.WithValue(ctx, "user", &user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		requestCounter.Inc()
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func Start() {
	registry.MustRegister(requestCounter, memoryUsage, goRoutineNum)
	conn, err := grpc.Dial(env.Conf.OrgsService.Host+":"+env.Conf.OrgsService.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to agent domain service: %v", err)
	}
	defer conn.Close()
	orgsClient := orgs.NewOrgServiceClient(conn)

	r := resolvers.Resolver{OrgClient: orgsClient}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &r}))

	http.Handle("/api", cors(getUser(srv)))
	promHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.Handle("/metrics", cors(monitoringMiddleware(promHandler)))
	http.Handle("/", playground.Handler("GraphQL playground", "/api"))

	log.Printf("Starting http server on port " + env.Conf.HttpPort)

	err = http.ListenAndServe(":"+env.Conf.HttpPort, nil)
	if err != nil {
		log.Fatalf("Error http server on port %v: %v", env.Conf.HttpPort, err)
	}
}
