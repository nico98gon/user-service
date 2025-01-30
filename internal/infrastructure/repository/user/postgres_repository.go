package repository

import (
	"database/sql"
	"errors"
	"nilus-challenge-backend/internal/domain/user"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindAll() ([]user.User, error) {
	query := `SELECT id, name, email, opt_out, locality_id, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.OptOut, &u.LocalityID, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}		
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresUserRepository) FindByID(id int) (*user.User, error) {
	query := `SELECT id, name, email, opt_out, locality_id, created_at, updated_at FROM users WHERE id = $1`
	var u user.User
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.OptOut, &u.LocalityID, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresUserRepository) Create(u *user.User) error {
	query := `INSERT INTO users (name, email, opt_out, locality_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, u.Name, u.Email, u.OptOut, u.LocalityID).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

func (r *PostgresUserRepository) Update(u *user.User) error {
	query := `UPDATE users SET name = $1, email = $2, opt_out = $3, locality_id = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5`
	result, err := r.db.Exec(query, u.Name, u.Email, u.OptOut, u.LocalityID, u.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *PostgresUserRepository) OptOut(id int) error {
	query := `UPDATE users SET opt_out = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
