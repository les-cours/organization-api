package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
	gprcToGraph "github.com/les-cours/organization-api/grpcToGraph"
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

func (r *mutationResolver) Test(ctx context.Context, in *string) (*models.OperationStatus, error) {

	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
