package auth

type User struct {
	UserId    int    `json:"user_id"`
	SectorId  int    `json:"sector_id"`
	RoleId    int    `json:"role_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AccStatus string `json:"acc_status"`
	RegStatus string `json:"reg_status"`
	Location  string `json:"location"`
}
