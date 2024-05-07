package resolvers

import (
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph"
)

type Resolver struct {
	OrgClient orgs.OrgServiceClient
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
