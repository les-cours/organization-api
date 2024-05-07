package grpcToGraph

import (
	"github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
)

func Department(department *orgs.Department) *models.Department {
	return &models.Department{
		DepartmentID:  department.DepartmentID,
		Title:         department.Title,
		ArabicTitle:   department.Description,
		Schools:       department.Description,
		Description:   department.Description,
		DescriptionAr: department.Description,
		Grades:        Grads(department.Grades),
	}

}

func Departments(cs *orgs.Departments) []*models.Department {
	var departments []*models.Department

	for _, d := range cs.Departments {
		departments = append(departments, Department(d))
	}
	return departments
}

func Grads(gs *orgs.Grads) []*models.Grad {
	var grads []*models.Grad

	for _, d := range gs.Grads {
		grads = append(grads, Grad(d))
	}
	return grads
}

func Grad(g *orgs.Grad) *models.Grad {
	return &models.Grad{
		GradID:      g.GradID,
		Title:       g.Title,
		ArabicTitle: g.ArabicTitle,
		Subjects:    Subjects(g.Subjects),
	}

}

func Subjects(ss *orgs.Subjects) []*models.Subject {
	var departments []*models.Subject

	for _, d := range ss.Subjects {
		departments = append(departments, Subject(d))
	}
	return departments
}

func Subject(s *orgs.Subject) *models.Subject {
	return &models.Subject{
		SubjectID:   s.SubjectID,
		Title:       s.Title,
		ArabicTitle: s.ArabicTitle,
	}
}
