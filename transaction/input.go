package transaction

import "bwastartup/user"

type GetTransactionsByCampaignIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type GetTransactionByUserIDInput struct {
	ID int `uri:"id" binding:"required"`
}
