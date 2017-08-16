package models

type ActionLogCollection []ActionLog

func CreateNewActionLogCollection() *ActionLogCollection {
	return &ActionLogCollection{}
}