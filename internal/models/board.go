package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Board struct {
	bun.BaseModel `bun:"table:boards,alias:b"`

	ID            int       `bum:"id,pk,autoincrement"`
	Name          string    `bun:"name,notnull"`
	Description   string    `bun:"description"`
	Position      int       `bun:"position,notnull"`
	IsHidden      bool      `bun:"hidden,type:integer,notnull,default:false"`
	IsLocked      bool      `bun:"locked,type:integer,notnull,default:false"`
	RedirectUrl   string    `bun:"redirect,nullzero"`
	CreatedAt     time.Time `bun:"created_at,type:text,notnull,default:current_timestamp"`
	ParentBoardID int       `bun:"parent_board_id,nullzero"`
	CategoryID    int       `bun:"category_id,notnull"`

	ParentBoard *Board    `bun:"rel:belongs-to,join:parent_board_id=id"`
	Category    *Category `bun:"rel:belongs-to,join:category_id=id"`
}
