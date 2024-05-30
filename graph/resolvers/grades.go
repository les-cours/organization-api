package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
	gprcToGraph "github.com/les-cours/organization-api/grpcToGraph"
)

func (r *queryResolver) Grads(ctx context.Context, departmentID string) ([]*models.Grad, error) {

	res, err := r.OrgClient.GetGrads(ctx, &orgs.GetGradsRequest{
		DepartmentID: departmentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Grads(res), nil
}

func (r *queryResolver) Grad(ctx context.Context, gradID string) (*models.Grad, error) {

	res, err := r.OrgClient.GetGrad(ctx, &orgs.GetGradRequest{
		GradID: gradID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Grad(res), nil
}
