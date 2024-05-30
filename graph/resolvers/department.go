package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
	gprcToGraph "github.com/les-cours/organization-api/grpcToGraph"
)

func (r *queryResolver) Departments(ctx context.Context, schoolID string) ([]*models.Department, error) {

	res, err := r.OrgClient.GetDepartments(ctx, &orgs.GetDepartmentsRequest{
		SchoolID: schoolID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Departments(res), nil
}

func (r *queryResolver) Department(ctx context.Context, departmentID string) (*models.Department, error) {

	res, err := r.OrgClient.GetDepartment(ctx, &orgs.GetDepartmentRequest{
		DepartmentID: departmentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Department(res), nil
}
