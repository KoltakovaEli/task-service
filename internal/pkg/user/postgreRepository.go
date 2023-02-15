package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type postgreRepository struct {
	db *sql.DB
}

func createPointer(user User) *User {
	return &user
}

func NewPostgresRepository(db *sql.DB) *postgreRepository {
	return &postgreRepository{db}
}

func scanUsers(rows *sql.Rows) ([]*User, error) {
	var users []*User
	var u User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Login, &u.Email)
		users = append(users, createPointer(u))
		if err != nil {
			return []*User{}, fmt.Errorf("failed to scan users from db %w", err)
		}
	}
	return users, nil
}

func scanUser(row *sql.Row) (User, error) {
	var u User
	err := row.Scan(&u.ID, &u.Login, &u.Email)
	if err != nil {
		return User{}, fmt.Errorf("failed to scan users from db %w", err)
	}
	return u, nil
}

func (p postgreRepository) GetUsers(ctx context.Context) ([]*User, error) {
	sql := "SELECT id,login,email FROM users"
	rows, err := p.db.Query(sql)
	if err != nil {
		return []*User{}, err
	}
	users, err := scanUsers(rows)
	if err != nil {
		return []*User{}, err
	}
	return users, nil
}

func (p postgreRepository) Create(ctx context.Context, login string, email string) (User, error) {
	if login == "" {
		return User{}, errors.New("empty name")
	}
	sql := "INSERT INTO users (login, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	params := []interface{}{
		login,
		email,
		time.Now().UTC(),
		time.Now().UTC(),
	}
	row := p.db.QueryRow(sql, params...)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user into db %w", err)
	}
	return User{
		ID:    id,
		Login: login,
		Email: email,
	}, nil
}

func (p postgreRepository) GetByID(ctx context.Context, ID string) (User, error) {
	sql := "SELECT id,login,email FROM users WHERE id = $1"
	row := p.db.QueryRow(sql, ID)
	user, err := scanUser(row)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (p postgreRepository) Delete(ctx context.Context, ID string) error {
	sql := "DELETE FROM users WHERE id = $1;"
	_, err := p.db.Exec(sql, ID)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreRepository) Update(ctx context.Context, ID string, login string, email string) (User, error) {
	sql := "UPDATE users SET login = $1,email = $2  WHERE id = $3 RETURNING id"
	row := p.db.QueryRow(sql, login, email, ID)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user into db %w", err)
	}
	return User{
		ID:    id,
		Login: login,
		Email: email,
	}, nil
}
