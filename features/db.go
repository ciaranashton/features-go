package features

import (
	"log"
	"net/http"
	"os"

	"github.com/CiaranAshton/features/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB interface allows us to switch out mongo for another db painlessly
type DB interface {
	GetAllFeatures(fa FeatureAPI, w http.ResponseWriter) ([]models.Feature, error)
	GetFeature(fa FeatureAPI, oid bson.ObjectId, f *models.Feature) error
	CreateFeature(fa FeatureAPI, f *models.Feature) error
	DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error
}

// Database type contains all methods for accessing the db
type Database struct{}

// Store the session for db methods
var session *mgo.Session

// NewDatabase initiates a new database
func NewDatabase() DB {
	session = getSession()
	return &Database{}
}

// GetSession Currently, hitting a local instance of mongo.
func getSession() *mgo.Session {
	log.Println("[MongoDB] Connecting to MongoDB...")
	s, err := mgo.Dial(os.Getenv("MONGO_CONNECT"))

	if err != nil {
		log.Fatalln(err)
	}
	return s
}

// GetAllFeatures does something
func (db Database) GetAllFeatures(fa FeatureAPI, w http.ResponseWriter) ([]models.Feature, error) {
	fs := []models.Feature{}

	if err := session.DB("cjla").C("features").Find(nil).All(&fs); err != nil {
		w.WriteHeader(404)
		return fs, err
	}

	return fs, nil
}

// GetFeature is a query for getting a feature by id from the database
func (db Database) GetFeature(fa FeatureAPI, oid bson.ObjectId, f *models.Feature) error {
	err := session.DB("cjla").C("features").FindId(oid).One(&f)

	return err
}

// CreateFeature persists a given feature in the databasae
func (db Database) CreateFeature(fa FeatureAPI, f *models.Feature) error {
	err := session.DB("cjla").C("features").Insert(&f)

	return err
}

// DeleteFeature perminantly removes a feature (oid) from the database
func (db Database) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	err := session.DB("cjla").C("features").RemoveId(oid)

	return err
}
