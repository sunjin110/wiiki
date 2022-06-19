package repomodel

import "time"

type User struct {
	ID        string    `xorm:"id"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"email"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type UpdateUser struct {
	Name      *string
	Email     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
