package models

type ActionLog struct {
	Id           string				`bson:"_id,omitempty"`
	ActionName   string				`bson:"action_name" validate:"nonzero,min=1,max=255"`
	ActionTarget string				`bson:"action_target" validate:"min=1,max=255"`
	ActionCost   int64				`bson:"action_cost"`
	UserId       string				`bson:"user_id" validate:"min=1,max=100"`
	GuestUserId  string				`bson:"guest_user_id" validate:"nonzero,min=1,max=100"`
	Url          string				`bson:"url" validate:"min=1,max=255"`
	CreatedAt    int64				`bson:"created_at" validate:"nonzero,min=1"`
	Params		 map[string]string  `bson:"params" validate:"max=100"`
	Tags		 []string 			`bson:"tags" validate:"max=100"`
}

func CreateNewActionLog() *ActionLog {
	return &ActionLog{}
}