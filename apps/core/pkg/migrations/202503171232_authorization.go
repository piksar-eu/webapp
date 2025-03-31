package migrations

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/di"
)

func m202503171232_authorization() Migration {
	return Migration{
		Id: "202503171232_authorization",
		Up: func() error {
			db := di.NewDb()

			if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS auth__roles (
				id VARCHAR PRIMARY KEY,
				name VARCHAR(255),
				permissions JSONB
			);`); err != nil {
				return err
			}

			if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS auth__user_roles (
				user_id VARCHAR REFERENCES auth__users(id) ON DELETE CASCADE,
				role_id VARCHAR REFERENCES auth__roles(id) ON DELETE CASCADE,
				PRIMARY KEY (user_id, role_id)
			);`); err != nil {
				return err
			}

			if _, err := db.Exec(`INSERT INTO auth__roles (id, name, permissions) VALUES ('roAdm', 'System Administrator', '[]');`); err != nil {
				return err
			}

			if _, err := db.Exec(`INSERT INTO auth__users (id, email, auth_methods) VALUES ('us5s81m', 'admin@pragmatyczny.dev', '[{"Data": {"salt": "2e34b112cc8ba31bba2fdab34286eb630d09d281babd69972e2c7028cf4b9d1f", "verifier": "2f38ef202a2be26c56283af925bcf186ad82928eb283884f93f9dbf9fd33f33deb4a651d5ad3227edb78ec0e5f18adbe561c027246d037e887baaa5900b483f54d644385ce829688ab792002d9ad9c25c7355a9b35d33bf49c7373a376c588585d2b0ffc042782adf82d9ca23e28664bdb172933d088f59b76798e7c5363ccc5cdd798a4205cfad70bbbe7b04f46d999c4795dbf6fd4cb4213f332791fc00bf96a19abfb697ef88ed960ffb4ebcf1b10330387f9a1e3519d2174621b7a3471702344238ae23be7c976f78ce7d50254b2346bf9ba08bb00e48aa92b91f4447a004a3380ce05d5eae768846100c9b3d4af41b370e2e2cde57096dadd83d761d60e42c3032f3967ab9fa144f500f497d24fa3f19025df0523628e442f6886a8234f7afb8c0c2d42523e2086c5796f15ecebaba3304c5bbfdd68c22b9d85be63d4a13e44048356968a5554a6a50c5d0fe0aff05f23144945dc72bad3e3a7cdfd5ba1a16de39d96813458fa40bacb180ab9f4aa4ee97042ce7e8d4f909e6d44548d2a"}, "Method": "srp"}]');`); err != nil {
				return err
			}

			if _, err := db.Exec(`INSERT INTO auth__user_roles (user_id, role_id) VALUES ('us5s81m', 'roAdm');`); err != nil {
				return err
			}

			return nil
		},
	}
}
