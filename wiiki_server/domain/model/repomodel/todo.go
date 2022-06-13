package repomodel

import "time"

type Todo struct {
	ID        string    `xorm:"id"`
	Text      string    `xorm:"text"`
	Done      bool      `xorm:"done"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}
