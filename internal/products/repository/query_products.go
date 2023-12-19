package repository

const (
	insertMenuCategory = `INSERT INTO coffee_service.menu_categories (name, created_at) VALUES (:name, :created_at);`
	detailMenuCategory = `SELECT mc.id, mc.name, mc.created_at, mc.updated_at FROM coffee_service.menu_categories mc WHERE mc.id = ?;`
	updateMenuCategory = `UPDATE coffee_service.menu_categories SET name =:name, updated_at=:updated_at WHERE id=:id;`
	menuCategories     = `SELECT mc.id, mc.name, mc.created_at, mc.updated_at FROM coffee_service.menu_categories mc;`
)
