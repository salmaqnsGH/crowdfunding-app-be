package transaction

import (
	"time"

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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
