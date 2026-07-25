package queries

type AuthQueries struct {
	CreateUser       string
	ReadUsers        string
	UpdateUserStatus string
	LoginUser        string
}

var Auth_Queries AuthQueries = AuthQueries{
	CreateUser:       ``,
	ReadUsers:        ``,
	UpdateUserStatus: ``,
	LoginUser:        ``,
}
