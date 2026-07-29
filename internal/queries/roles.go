package queries

type RoleQueries struct {
	CreateRole       string
	ReadRoles        string
	UpdateRoleTitle  string
	UpdateRoleStatus string
}

var Role_Queries = RoleQueries{
	CreateRole: `
		INSERT INTO roles
		(
			sector_id,
			role_title,
			role_status
		)
		VALUES ($1,$2,$3)
		RETURNING role_id`,

	ReadRoles: `
		SELECT
			role_id,
			sector_id,
			role_title,
			role_status
		FROM roles`,

	UpdateRoleTitle: `
		UPDATE roles
		SET role_title = $1
		WHERE role_id = $2`,

	UpdateRoleStatus: `
		UPDATE roles
		SET role_status = $1
		WHERE role_id = $2`,
}
