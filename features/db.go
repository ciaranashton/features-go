package features

import (
	"log"
	"os"

	"github.com/CiaranAshton/features/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB interface allows us to switch out mongo for another db painlessly
type DB interface {
	GetAllFeatures(debug *log.Logger, fs *[]models.Feature) error
	GetFeature(debug *log.Logger, id string, f *models.Feature) error
	CreateFeature(debug *log.Logger, f *models.Feature) error
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
func (db Database) GetAllFeatures(debug *log.Logger, fs *[]models.Feature) error {
	debug.Println("[MongoDB] Fetching all features from Mongo")
	err := session.DB("cjla").C("features").Find(nil).All(fs)

	return err
}

// GetFeature is a query for getting a feature by id from the database
func (db Database) GetFeature(debug *log.Logger, id string, f *models.Feature) error {
	debug.Printf("[MongoDB] Fetching feature %v from Mongo", id)
	oid := bson.ObjectIdHex(id)
	err := session.DB("cjla").C("features").FindId(oid).One(f)

	return err
}

// CreateFeature persists a given feature in the databasae
func (db Database) CreateFeature(debug *log.Logger, f *models.Feature) error {
	debug.Printf("[MongoDB] Persisting feature %s to database\n", f.Name)
	err := session.DB("cjla").C("features").Insert(&f)

	return err
}

// DeleteFeature perminantly removes a feature (oid) from the database
func (db Database) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	err := session.DB("cjla").C("features").RemoveId(oid)

	return err
}
