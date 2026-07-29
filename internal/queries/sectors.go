package queries

type SectorQueries struct {
	CreateSector       string
	UpdateSectorName   string
	UpdateSectorStatus string
	ReadSectors        string
}

var Sector_Queries = SectorQueries{
	CreateSector: `
		INSERT INTO sectors
		(
			sector_name,
			sector_status,
			sector_description
		)
		VALUES ($1,$2,$3)
		RETURNING sector_id`,

	UpdateSectorName: `
		UPDATE sectors
		SET sector_name = $1
		WHERE sector_id = $2`,

	UpdateSectorStatus: `
		UPDATE sectors
		SET sector_status = $1
		WHERE sector_id = $2`,

	ReadSectors: `
		SELECT
			sector_id,
			sector_name,
			sector_status,
			sector_description
		FROM sectors`,
}
