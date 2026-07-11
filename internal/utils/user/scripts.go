package registration

import (
	"strconv"
)

func ConvertIDs(i any) (int,error) {
	id,err := strconv.ParseInt(i,10,64)

	if err == nil return id,nil
	return nil,err
}

func (user *User) UpdateUserStructProps(u models.User) {
	user.SectorID = u.SectorID
	user.RoleID = u.RoleID
	user.UserName = u.UserName
	user.Email = u.Email
	user.Password = u.Password
	user.AccStatus = u.AccStatus
	user.RegStatus = u.RegStatus
	user.Location = u.Location
}

func (u User) GetUserRole(roles []Role) (string,error) {
	the_err := errors.New("Failed to get user role")

	uid,err := ConvertIDs(u.RoleID)
	if err != nil {
		log.Printf("Error while converting user ID: %v\n")
		return nil,err
	}

	for _,v := range roles{
		if vid,_ := ConvertIDs(v.RoleID); vid == uid return v.RoleTitle,nil
	}

	return nil,the_err
}