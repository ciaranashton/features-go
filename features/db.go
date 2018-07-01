package features

import (
	"log"
	"os"

	"github.com/CiaranAshton/features-go/logger"
	"github.com/CiaranAshton/features-go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB interface allows us to switch out mongo for another db painlessly
type DB interface {
	GetAllFeatures(l *logger.Logger, fs *[]models.Feature) error
	GetFeature(l *logger.Logger, id string, f *models.Feature) error
	CreateFeature(l *logger.Logger, f *models.Feature) error
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
	l := logger.NewLogger()

	l.Info.Println("[MongoDB] Connecting to MongoDB...")
	s, err := mgo.Dial(os.Getenv("MONGO_CONNECT"))

	if err != nil {
		log.Fatalln(err)
	}
	return s
}

// GetAllFeatures does something
func (db Database) GetAllFeatures(l *logger.Logger, fs *[]models.Feature) error {
	l.Debug.Println("[MongoDB] Fetching all features from Mongo")
	err := session.DB("cjla").C("features").Find(nil).All(fs)

	return err
}

// GetFeature is a query for getting a feature by id from the database
func (db Database) GetFeature(l *logger.Logger, id string, f *models.Feature) error {
	l.Debug.Printf("[MongoDB] Fetching feature %v from Mongo", id)
	oid := bson.ObjectIdHex(id)
	err := session.DB("cjla").C("features").FindId(oid).One(f)

	return err
}

// CreateFeature persists a given feature in the databasae
func (db Database) CreateFeature(l *logger.Logger, f *models.Feature) error {
	l.Debug.Printf("[MongoDB] Persisting feature %s to database\n", f.Name)
	err := session.DB("cjla").C("features").Insert(&f)

	return err
}

// DeleteFeature perminantly removes a feature (oid) from the database
func (db Database) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	err := session.DB("cjla").C("features").RemoveId(oid)

	return err
}
