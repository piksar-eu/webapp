package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

func NewPgAuthRoleRepository(db *sql.DB) auth.RoleRepository {
	return &pgAuthRoleRepository{
		db: db,
	}
}

type pgAuthRoleRepository struct {
	db *sql.DB
}

func (r *pgAuthRoleRepository) Get(id string) (*auth.Role, error) {
	row := r.db.QueryRow("SELECT id, name, permissions FROM auth__roles WHERE id = $1 LIMIT 1", id)

	role, err := r.scanToRole(row)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *pgAuthRoleRepository) List() (map[string]*auth.Role, error) {
	rows, err := r.db.Query("SELECT id, name, permissions FROM auth__roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := make(map[string]*auth.Role)

	for rows.Next() {
		role, err := r.scanToRole(rows)

		if err != nil {
			return nil, err
		}

		roles[role.Id] = role
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *pgAuthRoleRepository) Save(role *auth.Role) error {
	permissionsJSON, err := json.Marshal(role.Permissions)
	if err != nil {
		return fmt.Errorf("permissions can not be marschal")
	}
	query := `
		INSERT INTO auth__roles (id, name, permissions)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
			permissions = EXCLUDED.permissions;
	`
	_, err = r.db.Exec(query, role.Id, role.Name, permissionsJSON)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgAuthRoleRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM auth__roles WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *pgAuthRoleRepository) scanToRole(row S) (*auth.Role, error) {
	var id, name, permissionsRaw string
	var permissions []string

	err := row.Scan(&id, &name, &permissionsRaw)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(permissionsRaw), &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal permissions: %v", err)
	}

	return &auth.Role{
		Id:          id,
		Name:        name,
		Permissions: permissions,
	}, nil
}

type S interface {
	Scan(dest ...any) error
}

func (r *pgAuthRoleRepository) NewId() string {
	for {
		id, _ := shared.RandId("ro", 5)
		user, _ := r.Get(id)
		if user != nil {
			continue
		}

		return id
	}
}
