package stores

import (
	"intlogs/models"

	"gopkg.in/mgo.v2"
)

type ActionLogStore struct {
	collection *mgo.Collection
}

func CreateNewActionLogStore(collection *mgo.Collection) *ActionLogStore {
	return &ActionLogStore{collection}
}

func (store *ActionLogStore) Save(log *models.ActionLog) (*models.ActionLog, error) {
	err := store.collection.Insert(log)

	return log, err
}

func (store *ActionLogStore) All(logs *models.ActionLogCollection, limit int, offset int) (*models.ActionLogCollection, error) {
	err := store.collection.Find(nil).Limit(limit).Skip(offset).All(logs)

	return logs, err
}

func (store *ActionLogStore) Count() (n int, err error) {
	return store.collection.Count()
}
