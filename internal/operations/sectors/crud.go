package sectors

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func CreateSector(
	ctx context.Context,
	db *sql.DB,
	query string,
	s Sector,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		s.SectorName,
		s.SectorStatus,
		s.SectorDescription,
	).Scan(&s.SectorId)

	if err != nil {
		log.Println("Error occurred while creating sector")
		return err
	}

	log.Println("Sector created successfully")
	return nil
}

func ReadSectors(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Sector, error) {
	sectors := []Sector{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		log.Fatal("Error occurred while querying database table sectors")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s Sector

		err := rows.Scan(
			s.SectorId,
			s.SectorName,
			s.SectorStatus,
			s.SectorDescription,
		)

		if err != nil {
			log.Fatal("Error occurred while scanning table rows")
			return nil, err
		}

		sectors = append(sectors, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error occurred while fetching rows")
		return nil, err
	}

	log.Println("Sectors fetched successfully")

	return sectors, nil
}

func UpdateSectorStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int,
) error {
	result, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		log.Fatal("Error occurred while querying the database")
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		log.Fatal("No rows affected")
		return err
	}

	if aff == 0 {
		log.Fatal("No rows affected")
		return errors.New("Record not found")
	}

	return nil
}
