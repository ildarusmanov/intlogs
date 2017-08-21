package models

type ActionLogCollection []ActionLog

func CreateNewActionLogCollection() *ActionLogCollection {
	return &ActionLogCollection{}
}

func MakeNewActionLogCollection() ActionLogCollection {
	return make(ActionLogCollection, 0)
}