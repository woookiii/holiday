package entity

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	Id          uuid.UUID `db:"id" binding:"required"`
	Name        string    `db:"name" binding:"required"`
	Email       string    `db:"email" binding:"required"`
	Password    string    `db:"password" binding:"required"`
	Role        string    `db:"role" binding:"required"`
	CreatedTime time.Time `db:"created_time" binding:"required"`
	UpdatedTime time.Time `db:"updated_time"`
	DeletedTime time.Time `db:"deleted_time"`
}
