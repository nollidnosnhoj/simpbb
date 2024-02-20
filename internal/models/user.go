package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int       `bun:"id,pk,autoincrement"`
	Username      string    `bun:"username,notnull"`
	Password      string    `bun:"password,notnull"`
	Email         string    `bun:"email,notnull"`
	CreatedAt     time.Time `bun:"created_at,type:text,notnull,default:current_timestamp"`

	Threads []*Thread `bun:"rel:has-many,join:id=user_id"`
}
