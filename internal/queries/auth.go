package queries

type AuthQueries struct {
	CreateUser       string
	ReadUsers        string
	UpdateUserStatus string
	LoginUser        string
}

var Auth_Queries AuthQueries = AuthQueries{
	CreateUser: `INSERT INTO users
				(sector_id,role_id,user_name,email,password,acc_status,reg_type,location)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
				`,
	ReadUsers:        `SELECT * FROM users`,
	UpdateUserStatus: `UPDATE users SET acc_status = $1 WHERE id = $2`,
	LoginUser: `SELECT 
					sector_id,
					role_id,
					user_name,
					email,
					acc_status,
					reg_type,
					location
				WHERE email = $1, password = $2`,
}
