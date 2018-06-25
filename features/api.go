package features

import (
	mgo "gopkg.in/mgo.v2"
)

// FeatureAPI structure
type FeatureAPI struct {
	session *mgo.Session
}

// New method for creating a new instance of a FeatureAPI
func New(s *mgo.Session) *FeatureAPI {
	return &FeatureAPI{s}
}
