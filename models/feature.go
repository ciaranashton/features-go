package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Feature model
type Feature struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name"`
	Enabled bool          `json:"enabled"`
}
