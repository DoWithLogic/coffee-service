package repository

const (
	insertUsers       = `INSERT INTO coffee_service.users (username, email, password, gender, birthday, created_at) VALUES (:username, :email, :password, :gender, :birthday, :created_at)`
	userDetail        = `SELECT id, username, email, password, gender, birthday, points, created_at, updated_at FROM coffee_service.users where id = ?;`
	userDetailByEmail = `SELECT id, username, email, password, gender, birthday, points, created_at, updated_at FROM coffee_service.users where email = ?;`
	updateUserProfile = `UPDATE coffee_service.users
						 SET
						   username = CASE WHEN COALESCE(:username, '') = '' THEN username ELSE :username END,
						   email = CASE WHEN COALESCE(:email, '') = '' THEN email ELSE :email END,
						   password = CASE WHEN COALESCE(:password, '') = '' THEN password ELSE :password END,
						   gender = CASE WHEN COALESCE(:gender, '') = '' THEN gender ELSE :gender END,
						   birthday = CASE WHEN COALESCE(:birthday, '') = '' THEN birthday ELSE :birthday END,
						   updated_at = :updated_at
						 WHERE id = :id;`
	updateUserPoint = `UPDATE coffee_service.users SET points = ?, updated_at = ? WHERE id = ?;`
)
