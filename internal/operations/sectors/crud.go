package sectors

import (
	"context"
	"database/sql"
	"fmt"
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
		return fmt.Errorf("create sector: %w", err)
	}

	log.Println("Sector created successfully")
	return nil
}

func ReadSectors(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Sector, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query sectors: %w", err)
	}
	defer rows.Close()

	var sectors []Sector

	for rows.Next() {
		var s Sector

		if err := rows.Scan(
			&s.SectorId,
			&s.SectorName,
			&s.SectorStatus,
			&s.SectorDescription,
		); err != nil {
			return nil, fmt.Errorf("scan sector: %w", err)
		}

		sectors = append(sectors, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sectors: %w", err)
	}

	log.Println("Sectors fetched successfully")

	return sectors, nil
}

func UpdateSectorStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)
	if err != nil {
		return fmt.Errorf("update sector status: %w", err)
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	log.Println("Sector status updated successfully")

	return nil
}
