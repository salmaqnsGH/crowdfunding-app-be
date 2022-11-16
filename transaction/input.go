package transaction

import user "github.com/salmaqnsGH/crowdfunding-app/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
