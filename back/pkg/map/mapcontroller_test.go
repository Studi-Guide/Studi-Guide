package maps

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/roomcontroller/controllers"
	"studi-guide/pkg/roomcontroller/models"
	"testing"
)

func TestMapController_GetMapItems(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/", nil)

	provider :=  controllers.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(GetExpectedJson(provider.RoomList, provider.ConnectorList))
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

// Helper method
func GetExpectedJson(rooms []models.Room, connectors []models.ConnectorSpace) ([]models.MapItem)	 {
	var mapItems []models.MapItem
	for _, room := range rooms {
		mapItems = append(mapItems, room.MapItem)
	}

	for _, connector := range connectors {
		mapItems = append(mapItems, connector.MapItem)
	}

	return mapItems
}
