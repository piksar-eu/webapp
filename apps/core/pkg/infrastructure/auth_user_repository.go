package infrastructure

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

func NewPgAuthUserRepository(db *sql.DB) auth.UserRepository {
	return &pgAuthUserRepository{
		db: db,
	}
}

type pgAuthUserRepository struct {
	db *sql.DB
}

func (r *pgAuthUserRepository) GetById(id string) (*auth.User, error) {
	row := r.db.QueryRow("SELECT id, email, name, auth_methods, created_at FROM auth__users WHERE id = $1 LIMIT 1", id)

	user, err := r.scanToUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *pgAuthUserRepository) GetByEmail(email string) (*auth.User, error) {
	row := r.db.QueryRow("SELECT id, email, name, auth_methods, created_at FROM auth__users WHERE email = $1 LIMIT 1", email)

	user, err := r.scanToUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *pgAuthUserRepository) Save(user *auth.User) error {

	authMethodsJSON, err := json.Marshal(user.AuthMethods)
	if err != nil {
		return fmt.Errorf("AuthMethods can not be marschal")
	}
	query := `
		INSERT INTO auth__users (id, email, name, auth_methods, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET email = EXCLUDED.email,
			name = EXCLUDED.name,
			auth_methods = EXCLUDED.auth_methods,
			created_at = EXCLUDED.created_at;
	`
	_, err = r.db.Exec(query, user.Id, user.Email, user.Name, authMethodsJSON, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgAuthUserRepository) NewId() string {
	for {
		id, _ := shared.RandId("us", 7)
		user, _ := r.GetById(id)
		if user != nil {
			continue
		}

		return id
	}
}

func (r *pgAuthUserRepository) scanToUser(row s) (*auth.User, error) {
	user := &auth.User{}
	var authMethodsRaw string

	err := row.Scan(&user.Id, &user.Email, &user.Name, &authMethodsRaw, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var rawMethods []json.RawMessage
	err = json.Unmarshal([]byte(authMethodsRaw), &rawMethods)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal auth_methods: %v", err)
	}

	for _, rawMethod := range rawMethods {
		var m auth.AuthMethod
		err := json.Unmarshal(rawMethod, &m)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal auth method: %v", err)
		}

		switch m.Method {
		case "srp":
			var srpData auth.SRPData
			if err := shared.MapToStruct(m.Data, &srpData); err != nil {
				return nil, fmt.Errorf("failed to unmarshal SRP data: %v", err)
			}
			m.Data = srpData
		default:
			return nil, fmt.Errorf("unsupported auth method: %s", m.Method)
		}

		user.AuthMethods = append(user.AuthMethods, m)
	}

	return user, nil
}

type s interface {
	Scan(dest ...any) error
}
