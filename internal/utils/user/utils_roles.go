package registration

import (
	"strconv"
	"database/sql"
	"context"
	"log"
)

func CreateRole(ctx context.Context, db *sql.DB, query string, r *Role) error {
	err := db.QueryRowContext(
		ctx,
		query,
		r.RoleTitle
	).Scan(r.RoleID)

	if err == nil return nil

	log.Printf("Error occurred while querying database:%v\n",err)
	return err
}

func ReadRoles(ctx context.Context, db *sql.DB, query string) ([]Role,error) {
	roles := []Role{}

	rows,err := db.QueryContext(ctx,query)

	if err != nil{
		log.Printf("Error occurred while querying the database")
		return nil,err
	}

	defer rows.Close()

	for rows.Next(){
		var r Role

		err := rows.Scan(
			&r.RoleID,
			&r.SectorID,
			&r.RoleTitle
		)

		if err != nil {
			log.Printf("Error while scanning record: %v\n",err)
			return nil,err
		}

		roles = append(roles,r)
	}

	if err = rows.Err(); err != nil{}

	return roles,nil
}

func UpdateRole(ctx context.Context, db *sql.DB, query string, r *Role) error {
	id,err := ConvertIDs(r.RoleID)

	if err != nil {
		log.Printf("Error occurred while converting ID: %v\n", err)
		return err
	}

	result,err := db.ExecContext(ctx,query,r.RoleTitle,id)

	if err != nil{
		log.Printf("Error while querying database:%v\n",err)
		return err
	}

	aff,err := result.RowsAffected()

	if err != nil{
		log.Printf("Error getting rows affected: %v\n", err)
		return err		
	}

	if aff == 0{
		log.Printf("Update failed: user with ID %d not found\n", id)
		return ErrUserNotFound		
	}

	log.Println("Record successfully updated")
	return nil
}