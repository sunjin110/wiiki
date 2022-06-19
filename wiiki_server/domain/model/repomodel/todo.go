package repomodel

import "time"

type Todo struct {
	ID        string    `xorm:"id"`
	Text      string    `xorm:"text"`
	Done      bool      `xorm:"done"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type UpdateTodo struct {
	Text      *string
	Done      *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
