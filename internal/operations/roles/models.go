package roles

type Role struct {
	RoleId     int    `json:"role_id"`
	SectorId   int    `json:"sector_id"`
	RoleTitle  string `json:"role_title"`
	RoleStatus string `json:"role_status"`
}
