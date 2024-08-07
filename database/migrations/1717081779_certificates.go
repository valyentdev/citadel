package migrations

import (
	"citadel/internal/models"
	"context"

	"github.com/uptrace/bun"
)

func certificatesMigrationUp_1717081779(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*models.Certificate)(nil)).Exec(ctx)
	return err
}

func certificatesMigrationDown_1717081779(ctx context.Context, db *bun.DB) error {
	_, err := db.NewDropTable().Model((*models.Certificate)(nil)).Exec(ctx)
	return err
}

func init() {
	Migrations.MustRegister(certificatesMigrationUp_1717081779, certificatesMigrationDown_1717081779)
}
