package models

type Sector struct{
	SectorID int `json:"sector_id"`
	SectorName string `json:"sector_name"`
	SectorDescription string `json:"sector_description"`
}

type User struct{
	UserID int `json:"user_id"`
	SectorID int `json:"sector_id"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	AccStatus string `json:"acc_status"`
	RegStatus string `json:"reg_status"`
}