package models

type ActionLog struct {
	Id           string	`bson:"_id,omitempty"`
	ActionName   string	`bson:"action_name"`
	ActionTarget string	`bson:"action_target"`
	UserId       string	`bson:"user_id"`
	GuestUserId  string	`bson:"guest_user_id"`
	Url          string	`bson:"url"`
	CreatedAt    int	`bson:"created_at"`
}

func CreateNewActionLog() *ActionLog {
	return &ActionLog{}
}