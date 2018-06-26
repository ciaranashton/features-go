package features

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// FeatureAPI structure
type FeatureAPI struct {
	session *mgo.Session
	info    *log.Logger
}

// New method for creating a new instance of a FeatureAPI
func New(s *mgo.Session, i *log.Logger) *FeatureAPI {
	return &FeatureAPI{s, i}
}
