package repomodel

import "time"

type User struct {
	ID        int       `xorm:"id"`
	Name      string    `xorm:"name"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}
