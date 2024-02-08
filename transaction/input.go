package transaction

import "bwastartup/user"

type ParamTransaction struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
