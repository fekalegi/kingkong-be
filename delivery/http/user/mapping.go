package user

import (
	"kingkong-be/delivery/http/user/model"
	entity "kingkong-be/domain/user"
	"kingkong-be/helper"
)

func mapRequestAddUser(req *model.User, e *entity.User) {
	e.UserID = req.UserID
	e.Username = req.Username
	e.Password = helper.GetMD5String(req.Password)
	e.Role = req.Role
}
