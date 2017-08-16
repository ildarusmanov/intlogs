package models

type ActionLog struct {
	Id string
	ActionName string
	ActionTarget string
	UserId string
	GuestUserId string
	Url string
	CreatedAt int
}

func CreateNewActionLog() *ActionLog {
	return &ActionLog{}
}