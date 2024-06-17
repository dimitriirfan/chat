package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
)

type UsersRepository interface {
	Save(ctx context.Context, user entity.User) (lastInsertId int64, err error)
	Search(ctx context.Context, params entity.SearchUserParams) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
}

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Save(ctx context.Context, user entity.User) (lastInsertId int64, err error) {
	sql := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := r.db.Exec(sql, user.Username, user.Password)
	if err != nil {
		return -1, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastInsertId, nil
}

func (r *UsersRepo) Search(ctx context.Context, params entity.SearchUserParams) ([]entity.User, error) {
	sql := "SELECT * FROM users WHERE username = ?"
	rows, err := r.db.Query(sql, params.Username)
	if err != nil {
		return []entity.User{}, err
	}

	users := []entity.User{}
	for rows.Next() {
		var (
			user                 entity.User
			createdAt, updatedAt string
		)
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &createdAt, &updatedAt)
		if err != nil {
			return []entity.User{}, err
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		user.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

		users = append(users, user)
	}

	return users, nil
}

func (r *UsersRepo) GetByID(ctx context.Context, id int) (entity.User, error) {
	sql := "SELECT * FROM users WHERE id = ? LIMIT 1"
	row := r.db.QueryRow(sql, id)

	var (
		user                 entity.User
		createdAt, updatedAt string
	)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &createdAt, &updatedAt)
	if err != nil {
		return entity.User{}, err
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	user.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return user, nil
}
