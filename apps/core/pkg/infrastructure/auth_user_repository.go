package infrastructure

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

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

func (r *pgAuthUserRepository) Get(email string) (*auth.User, error) {
	row := r.db.QueryRow("SELECT auth_methods, created_at FROM auth__users WHERE email = $1 LIMIT 1", email)

	var authMethodsRaw string
	var authMethods []auth.AuthMethod
	var createdAt time.Time

	err := row.Scan(&authMethodsRaw, &createdAt)
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

		authMethods = append(authMethods, m)
	}

	return &auth.User{
		Email:       email,
		AuthMethods: authMethods,
		CreatedAt:   createdAt,
	}, nil
}

func (r *pgAuthUserRepository) Save(user *auth.User) error {

	authMethodsJSON, err := json.Marshal(user.AuthMethods)
	if err != nil {
		return fmt.Errorf("AuthMethods can not be marschal")
	}
	query := `
		INSERT INTO auth__users (email, auth_methods, created_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO UPDATE
		SET auth_methods = EXCLUDED.auth_methods,
			created_at = EXCLUDED.created_at;
	`
	_, err = r.db.Exec(query, user.Email, authMethodsJSON, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
