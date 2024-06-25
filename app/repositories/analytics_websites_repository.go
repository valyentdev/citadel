package repositories

import (
	"citadel/app/models"
	"context"

	"github.com/caesar-rocks/orm"
)

type AnalyticsWebsitesRepository struct {
	*orm.Repository[models.AnalyticsWebsite]
}

func NewAnalyticsWebsitesRepository(db *orm.Database) *AnalyticsWebsitesRepository {
	return &AnalyticsWebsitesRepository{Repository: &orm.Repository[models.AnalyticsWebsite]{
		Database: db,
	}}
}

func (r AnalyticsWebsitesRepository) FindAllFromUser(ctx context.Context, userId string) ([]models.AnalyticsWebsite, error) {
	var items []models.AnalyticsWebsite = make([]models.AnalyticsWebsite, 0)

	err := r.NewSelect().Model((*models.AnalyticsWebsite)(nil)).Where("user_id = ?", userId).Scan(ctx, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
