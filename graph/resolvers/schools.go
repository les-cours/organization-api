package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
	gprcToGraph "github.com/les-cours/organization-api/grpcToGraph"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (r *queryResolver) School(ctx context.Context, schoolID string) ([]*models.Department, error) {

	res, err := r.OrgClient.GetSchool(ctx, &orgs.IDRequest{
		Id: schoolID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Departments(res), nil
}

func (r *queryResolver) Dashboard(ctx context.Context) (*models.Dashboard, error) {

	var apps = make([]*models.App, 0)
	for i := 0; i < 5; i++ {
		apps = append(apps, &models.App{
			ID:          strconv.Itoa(i),
			Name:        "test" + strconv.Itoa(i),
			Image:       "Image" + strconv.Itoa(i),
			Description: "Description Description Description" + strconv.Itoa(i),
		})
	}
	a, b := randomN()
	return &models.Dashboard{
		RecentApps: apps,
		ProductViews: &models.ProductViews{
			Categories: []*string{randomString(8), randomString(4)},
			Products:   []*string{randomString(8), randomString(4)},
			Views:      []*int{b, b, b},
		},
		ProductsBar: &models.ProductsBar{
			Labels:   []*string{randomString(8), randomString(4)},
			Products: []*string{randomString(8), randomString(4)},
		},
		CashFlow: &models.CashFlow{
			Categories: []*string{randomString(8), randomString(4)},
			Cash:       []*float64{a, a, a},
		},
	}, nil
}

func (r *mutationResolver) Test(ctx context.Context, in *string) (*models.OperationStatus, error) {

	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}

func randomString(length int) *string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var output strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		output.WriteByte(charset[randomIndex])
	}
	r := output.String()
	return &r
}

func randomN() (*float64, *int) {
	a := rand.Float64()
	b := rand.Int()
	return &a, &b
}
