package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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

	roles, err := r.roles(user.Id)
	if err != nil {
		return nil, err
	}
	if roles[user.Id] != nil {
		user.Roles = roles[user.Id]
	}

	return user, nil
}

func (r *pgAuthUserRepository) GetByEmail(email string) (*auth.User, error) {
	row := r.db.QueryRow("SELECT id, email, name, auth_methods, created_at FROM auth__users WHERE email = $1 LIMIT 1", email)

	user, err := r.scanToUser(row)
	if err != nil {
		return nil, err
	}

	roles, err := r.roles(user.Id)
	if err != nil {
		return nil, err
	}
	if roles[user.Id] != nil {
		user.Roles = roles[user.Id]
	}

	return user, nil
}

func (r *pgAuthUserRepository) List() (map[string]*auth.User, error) {
	rows, err := r.db.Query("SELECT id, email, name, auth_methods, created_at FROM auth__users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make(map[string]*auth.User)

	roles, err := r.roles("")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := r.scanToUser(rows)
		if roles[user.Id] != nil {
			user.Roles = roles[user.Id]
		}

		if err != nil {
			return nil, err
		}

		users[user.Id] = user
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *pgAuthUserRepository) Save(user *auth.User) error {
	authMethodsJSON, err := json.Marshal(user.AuthMethods)
	if err != nil {
		return fmt.Errorf("authMethods can not be marschal")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

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
		return fmt.Errorf("failed to insert/update user: %v", err)
	}

	query = "DELETE FROM auth__user_roles WHERE user_id = $1"
	_, err = tx.Exec(query, user.Id)
	if err != nil {
		return fmt.Errorf("failed to delete existing roles: %v", err)
	}

	if len(user.Roles) > 0 {
		insertRolesQuery := "INSERT INTO auth__user_roles (user_id, role_id) VALUES "
		args := []interface{}{}
		values := []string{}

		for i, role := range user.Roles {
			values = append(values, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
			args = append(args, user.Id, role)
		}

		insertRolesQuery += fmt.Sprintf("%s;", strings.Join(values, ","))
		_, err = tx.Exec(insertRolesQuery, args...)
		if err != nil {
			return fmt.Errorf("failed to insert new roles: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *pgAuthUserRepository) scanToUser(row S) (*auth.User, error) {
	user := &auth.User{
		Roles: []string{},
	}
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

func (r *pgAuthUserRepository) roles(id ...string) (map[string][]string, error) {
	userRoles := make(map[string][]string)

	var query string
	var args []interface{}

	if len(id) == 1 && id[0] == "" {
		query = "SELECT user_id, role_id FROM auth__user_roles"
	} else {
		placeholders := make([]string, len(id))
		for i, v := range id {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args = append(args, v)
		}
		query = fmt.Sprintf("SELECT user_id, role_id FROM auth__user_roles WHERE user_id IN (%s)", strings.Join(placeholders, ","))
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userId, roleId string

	for rows.Next() {
		if err := rows.Scan(&userId, &roleId); err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		userRoles[userId] = append(userRoles[userId], roleId)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userRoles, nil
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
