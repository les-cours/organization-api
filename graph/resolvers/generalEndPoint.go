package resolvers

import (
	"context"
	orgs "github.com/les-cours/organization-api/api/orgs"
	"github.com/les-cours/organization-api/graph/models"
)

func (r *queryResolver) Cities(ctx context.Context) ([]*models.City, error) {

	res, err := r.OrgClient.GetCities(ctx, &orgs.Empty{})
	if err != nil {
		return nil, ErrApi(err)
	}

	var cities = make([]*models.City, 0)
	for _, city := range res.Cities {
		cities = append(cities, &models.City{
			ID:     int(city.Id),
			Name:   city.Name,
			NameAr: city.ArabicName,
		})
	}
	return cities, nil
}
