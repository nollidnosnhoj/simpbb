package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories,alias:cat"`

	ID        int       `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,nullzero,notnull"`
	Slug      string    `bun:"slug,nullzero,notnull"`
	Position  int       `bun:"position,notnull"`
	CreatedAt time.Time `bun:"created_at,type:text,nullzero,notnull,default:current_timestamp"`
}
