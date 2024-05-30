package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
	gprcToGraph "github.com/les-cours/organization-api/grpcToGraph"
)

func (r *queryResolver) Subjects(ctx context.Context, gradID string) ([]*models.Subject, error) {

	res, err := r.OrgClient.GetSubjectsByGrad(ctx, &orgs.IDRequest{
		Id: gradID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Subjects(res), nil
}
