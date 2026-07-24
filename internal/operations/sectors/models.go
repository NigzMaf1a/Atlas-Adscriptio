package sectors

type Sector struct {
	SectorId          int    `json:"sector_id"`
	SectorName        string `json:"sector_name"`
	SectorStatus      string `json:"sector_status"`
	SectorDescription string `json:"sector_description"`
}
