package queries

type SectorQueries struct {
	CreateSector       string
	UpdateSectorName   string
	UpdateSectorStatus string
}

var Sector_Queries SectorQueries = SectorQueries{
	CreateSector: `INSERT INTO sectors 
				   (sector_name,sector_status,sector_description)
				   VALUES ($1,$2,$3)
				   `,
	UpdateSectorName:   `UPDATE sectors SET sector_name = $1 WHERE id = $2`,
	UpdateSectorStatus: `UPDATE sectors SET sector_status = $1 WHERE id = $2`,
}
