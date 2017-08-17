package stores

import (
	"intlogs/models"

	"gopkg.in/mgo.v2"
)

type ActionLogStore struct {
	collection *mgo.Collection
}

func CreateNewActionLogStore(collection *mgo.Collection) *ActionLogStore {
	return &ActionLogStore{}
}

func (store *ActionLogStore) Save(log *models.ActionLog) *models.ActionLog {
	store.collection.Insert(log)

	return log
}

func (store *ActionLogStore) All(logs *models.ActionLogCollection) *models.ActionLogCollection {
	store.collection.Find(nil).All(logs)

	return logs
}

func (store *ActionLogStore) Count() (n int, err error) {
	return store.collection.Count()
}