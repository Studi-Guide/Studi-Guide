package location

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
	"testing"
)

func TestLocationController_GetLocations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectLocations := []entityservice.Location{entityservice.Location{
		Id:          1,
		Name:        "bla",
		Description: "descr",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}}

	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllLocations().Return(expectLocations, nil)

	router := gin.Default()
	locationRouter := router.Group("/location")
	_ = NewLocationController(locationRouter, mock)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/location/", nil)

	router.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expected ", http.StatusOK)
	}

	expectJson, _ := json.Marshal(expectLocations)
	expectJson = append(expectJson, '\n')
	if string(expectJson) != rec.Body.String() {
		t.Error("expect", expectJson, ", got ", rec.Body.Bytes())
	}

}

func TestLocationController_GetLocations2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectLocations := []entityservice.Location{entityservice.Location{
		Id:          1,
		Name:        "bla",
		Description: "descr",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}}

	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().FilterLocations("abc", "taaaag", "1", "KA", "KA").Return(expectLocations, nil)

	router := gin.Default()
	locationRouter := router.Group("/location")
	_ = NewLocationController(locationRouter, mock)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/location/?name=abc&tag=taaaag&floor=1&building=KA&campus=KA&", nil)

	router.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expected ", http.StatusOK)
	}

	expectJson, _ := json.Marshal(expectLocations)
	expectJson = append(expectJson, '\n')
	if string(expectJson) != rec.Body.String() {
		t.Error("expect", expectJson, ", got ", rec.Body.Bytes())
	}
}

func TestLocationController_GetLocations3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().FilterLocations("abc", "taaaag", "1", "KA", "KA").Return(nil, errors.New("error text"))

	router := gin.Default()
	locationRouter := router.Group("/location")
	_ = NewLocationController(locationRouter, mock)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/location/?name=abc&tag=taaaag&floor=1&building=KA&campus=KA&", nil)

	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusInternalServerError)
	}
}

func TestLocationController_GetLocationByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectLocation := entityservice.Location{
		Id:          1,
		Name:        "bla",
		Description: "descr",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}

	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetLocation("bla").Return(expectLocation, nil)

	router := gin.Default()
	locationRouter := router.Group("/location")
	_ = NewLocationController(locationRouter, mock)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/location/name/bla", nil)

	router.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expected ", http.StatusOK)
	}

	expectJson, _ := json.Marshal(expectLocation)
	expectJson = append(expectJson, '\n')
	if string(expectJson) != rec.Body.String() {
		t.Error("expect", expectJson, ", got ", rec.Body.Bytes())
	}
}

func TestLocationController_GetLocationByName2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetLocation("bla").Return(entityservice.Location{}, errors.New("error text"))

	router := gin.Default()
	locationRouter := router.Group("/location")
	_ = NewLocationController(locationRouter, mock)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/location/name/bla", nil)

	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}


}