package features

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// FeatureAPI structure
type FeatureAPI struct {
	session *mgo.Session
	info    *log.Logger
	warn    *log.Logger
	err     *log.Logger
}

// New function for creating an instance of FeatureAPI
func New(s *mgo.Session, i, w, e *log.Logger) *FeatureAPI {
	return &FeatureAPI{s, i, w, e}
}
