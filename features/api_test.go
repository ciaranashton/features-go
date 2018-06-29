package features

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CiaranAshton/features/models"
	"github.com/gavv/httpexpect"
	"gopkg.in/mgo.v2/bson"
)

type TestDatabase struct{}

func NewTestDatabase() DB {
	return &TestDatabase{}
}

func (db TestDatabase) GetAllFeatures(fa FeatureAPI, w http.ResponseWriter) ([]models.Feature, error) {
	f1 := models.Feature{
		Id:      "001",
		Name:    "Test 01",
		Enabled: true,
	}

	f2 := models.Feature{
		Id:      "002",
		Name:    "Test 02",
		Enabled: false,
	}

	fs := []models.Feature{f1, f2}

	return fs, nil
}

func (db TestDatabase) GetFeature(fa FeatureAPI, id string, w http.ResponseWriter) (models.Feature, error) {
	f1 := models.Feature{
		Id:      "001",
		Name:    "Test 01",
		Enabled: true,
	}

	return f1, nil
}

func (db TestDatabase) CreateFeature(fa FeatureAPI, f models.Feature) (models.Feature, error) {
	return f, nil
}

func (db TestDatabase) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	return nil
}

func TestGetFeatures(t *testing.T) {
	l := log.New(ioutil.Discard, "", 0)

	db := NewTestDatabase()
	api := New(db, l, l, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array().
		Contains(models.Feature{
			Id:      "001",
			Name:    "Test 01",
			Enabled: true,
		}, models.Feature{
			Id:      "002",
			Name:    "Test 02",
			Enabled: false,
		})
}

func TestGetFeature(t *testing.T) {
	l := log.New(ioutil.Discard, "", 0)

	db := NewTestDatabase()
	api := New(db, l, l, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features/001").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(models.Feature{
			Id:      "001",
			Name:    "Test 01",
			Enabled: true,
		})
}

func TestCreateFeature(t *testing.T) {
	l := log.New(ioutil.Discard, "", 0)

	db := NewTestDatabase()
	api := New(db, l, l, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.POST("/features").
		WithJSON(struct {
			Name    string
			Enabled bool
		}{
			Name:    "Test 01",
			Enabled: true,
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object().
		ValueEqual("name", "Test 01").
		ValueEqual("enabled", true)
}

func TestDeleteFeature(t *testing.T) {
	l := log.New(ioutil.Discard, "", 0)

	db := NewTestDatabase()
	api := New(db, l, l, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.DELETE("/features/5b315dc2379785611a23e4be").
		Expect().
		Status(http.StatusOK)
}
