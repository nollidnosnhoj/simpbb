package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Thread struct {
	bun.BaseModel `bun:"table:threads,alias:t"`

	ID            int       `bun:"id,pk,autoincrement"`
	Subject       string    `bun:"subject,notnull"`
	Slug          string    `bun:"slug,notnull"`
	Body          string    `bun:"body,notnull"`
	IsPinned      bool      `bun:"pinned,type:integer,nullzero,notnull,default:false"`
	IsLocked      bool      `bun:"locked,type:integer,nullzero,notnull,default:false"`
	CreatedAt     time.Time `bun:"created_at,type:text,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,type:text,nullzero"`
	UpdatedReason string    `bun:"updated_reason,nullzero"`
	BoardID       int       `bun:"board_id,notnull"`
	UserID        int       `bun:"user_id,notnull"`

	Board *Board `bun:"rel:belongs-to,join:board_id=id"`
	User  *User  `bun:"rel:belongs-to,join:user_id=id"`
}
