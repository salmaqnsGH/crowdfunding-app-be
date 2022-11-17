package transaction

import (
	"time"

	"github.com/salmaqnsGH/crowdfunding-app/campaign"
	user "github.com/salmaqnsGH/crowdfunding-app/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
