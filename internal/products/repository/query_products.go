package repository

const (
	insertMenuCategory = `INSERT INTO coffee_service.menu_categories (name, created_at) VALUES (:name, :created_at);`
	detailMenuCategory = `SELECT mc.id, mc.name, mc.created_at, mc.updated_at FROM coffee_service.menu_categories mc WHERE mc.id = ?;`
	updateMenuCategory = `UPDATE coffee_service.menu_categories SET name =:name, updated_at=:updated_at WHERE id=:id;`
	menuCategories     = `SELECT mc.id, mc.name, mc.created_at, mc.updated_at FROM coffee_service.menu_categories mc;`

	insertMenu = `INSERT INTO coffee_service.menu (menu_categories_id, name, description, price, created_at) VALUES (:menu_categories_id, :name, :description, :price, :created_at);`
	detailMenu = `SELECT m.id, m.menu_categories_id, m.name, m.description, m.price, m.created_at, m.updated_at FROM coffee_service.menu m WHERE m.id = ?`
	updateMenu = `UPDATE coffee_service.menu SET 
					menu_categories_id = CASE WHEN :menu_categories_id != '' THEN :menu_categories_id ELSE menu_categories_id END,
					name = CASE WHEN :name != '' THEN :name ELSE name END,
					description = CASE WHEN :description != '' THEN :description ELSE description END,
					price = CASE WHEN :price != '' THEN :price ELSE price END,
					updated_at = :updated_at
				  WHERE id = :id;`
	listMenu = `SELECT m.id, m.menu_categories_id, m.name, m.description, m.price, m.created_at, m.updated_at FROM coffee_service.menu m`
)
