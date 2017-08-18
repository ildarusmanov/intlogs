package models

type ActionLog struct {
	Id           string				`bson:"_id,omitempty"`
	ActionName   string				`bson:"action_name",validate:"nonzero"`
	ActionTarget string				`bson:"action_target"`
	ActionCost   int				`bson:"action_cost"`
	UserId       string				`bson:"user_id"`
	GuestUserId  string				`bson:"guest_user_id",validate:"nonzero"`
	Url          string				`bson:"url"`
	CreatedAt    int				`bson:"created_at",validate:"nonzero"`
	Params		 map[string]string  `bson:"params"`
	Tags		 []string 			`bson:"tags"`
}

func CreateNewActionLog() *ActionLog {
	return &ActionLog{}
}