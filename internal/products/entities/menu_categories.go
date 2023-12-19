package entities

import "time"

type MenuCategory struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdateAt  *time.Time `db:"updated_at"`
}

type UpdateMenuCategory struct {
	ID       int64      `db:"id"`
	Name     string     `db:"name"`
	UpdateAt *time.Time `db:"updated_at"`
}

type MenuCategories []MenuCategory
