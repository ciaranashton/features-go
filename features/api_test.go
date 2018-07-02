package features

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CiaranAshton/features-go/logger"
	"github.com/CiaranAshton/features-go/models"

	"github.com/gavv/httpexpect"
	"gopkg.in/mgo.v2/bson"
)

type TestDatabase struct{}

func NewTestDatabase() DB {
	return &TestDatabase{}
}

func (db TestDatabase) GetAllFeatures(l *logger.Logger, fs *[]models.Feature) error {
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

	*fs = []models.Feature{f1, f2}

	return nil
}

func (db TestDatabase) GetFeature(l *logger.Logger, id string, f *models.Feature) error {
	*f = models.Feature{
		Id:      "001",
		Name:    "Test 01",
		Enabled: true,
	}

	return nil
}

func (db TestDatabase) CreateFeature(l *logger.Logger, f *models.Feature) error {
	return nil
}

func (db TestDatabase) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	return nil
}

func (db TestDatabase) UpdateFeature(l *logger.Logger, oid bson.ObjectId, f *models.Feature) error {
	return nil
}

func TestGetFeatures(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

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
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features/5b315dc2379785611a23e4be").
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

func TestGetFeatureNonObjectId(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features/123").
		Expect().
		Status(http.StatusBadRequest)
}

func TestCreateFeature(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

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
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.DELETE("/features/5b315dc2379785611a23e4be").
		Expect().
		Status(http.StatusOK)
}

func TestDeleteFeatureNonObjectId(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.DELETE("/features/123").
		Expect().
		Status(http.StatusBadRequest)
}

func TestUpdateFeature(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.PUT("/features/5b315dc2379785611a23e4be").
		WithJSON(struct {
			Name string
		}{
			Name: "Test 02",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("name", "Test 02").
		ValueEqual("enabled", false)
}

func TestUpdateFeatureNonObjectId(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.PUT("/features/123").
		WithJSON(struct {
			Name string
		}{
			Name: "Test 02",
		}).
		Expect().
		Status(http.StatusBadRequest)
}

// TestErrorDatabase methods always return an error
type TestErrorDatabase struct{}

func NewTestErrorDatabase() DB {
	return &TestErrorDatabase{}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (db TestErrorDatabase) GetAllFeatures(l *logger.Logger, fs *[]models.Feature) error {
	return &errorString{"Error"}
}

func (db TestErrorDatabase) GetFeature(l *logger.Logger, id string, f *models.Feature) error {
	return &errorString{"Error"}
}

func (db TestErrorDatabase) CreateFeature(l *logger.Logger, f *models.Feature) error {
	return &errorString{"Error"}
}

func (db TestErrorDatabase) DeleteFeature(fa FeatureAPI, oid bson.ObjectId) error {
	return &errorString{"Error"}
}

func (db TestErrorDatabase) UpdateFeature(l *logger.Logger, oid bson.ObjectId, f *models.Feature) error {
	return &errorString{"Error"}
}

func TestGetFeaturesNotFound(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestErrorDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features").
		Expect().
		Status(http.StatusNotFound)
}

func TestGetFeatureNotFound(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestErrorDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/features/5b315dc2379785611a23e4be").
		Expect().
		Status(http.StatusNotFound)
}

func TestCreateFeatureDBError(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestErrorDatabase()
	api := New(db, l).API()

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
		Status(http.StatusBadRequest)
}

func TestDeleteFeatureDBError(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestErrorDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.DELETE("/features/5b315dc2379785611a23e4be").
		Expect().
		Status(http.StatusBadRequest)
}

func TestUpdateFeatureDBError(t *testing.T) {
	l := logger.NewLogger(true)

	db := NewTestErrorDatabase()
	api := New(db, l).API()

	server := httptest.NewServer(api)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.PUT("/features/5b315dc2379785611a23e4be").
		WithJSON(struct {
			Name string
		}{
			Name: "Test 02",
		}).
		Expect().
		Status(http.StatusNotFound)
}
