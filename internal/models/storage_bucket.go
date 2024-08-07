package models

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/uptrace/bun"
)

type StorageBucket struct {
	ID        string `bun:"id,pk"`
	Name      string `bun:"name"`
	Slug      string `bun:"slug,unique"`
	Host      string `bun:"host"`
	Region    string `bun:"region"`
	KeyId     string `bun:"key_id"`
	SecretKey string `bun:"secret_key"`

	Organization   *Organization `bun:"rel:belongs-to,join:organization_id=id"`
	OrganizationID string        `bun:"organization_id"`
}

var _ bun.BeforeAppendModelHook = (*StorageBucket)(nil)

func (bucket *StorageBucket) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		bucket.ID = xid.New().String()
	}
	return nil
}

type StorageFile struct {
	Size      float64
	Name      string
	UpdatedAt time.Time
	Type      string
}
