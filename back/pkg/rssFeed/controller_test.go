package rssFeed

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/utils"
	"testing"
)

func TestRssFeedController_GetRssFeedByName(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rssfeed/testfeed", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := NewMockProvider(ctrl)
	mockhttpclient := utils.NewMockHttpClient(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/rssfeed")
	_ = NewRssFeedController(mapRouter, provider, mockhttpclient)

	feed := ent.RssFeed{
		ID:   1,
		URL:  "http://www.testfeed/rss.xml",
		Name: "testfeed",
	}

	teststring := "hello world"
	r := ioutil.NopCloser(strings.NewReader(teststring)) // r type is io.ReadCloser
	provider.EXPECT().GetRssFeed("testfeed").Return(&feed, nil)
	mockhttpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{Body: r}, nil)
	router.ServeHTTP(rec, req)

	actual := rec.Body.String()
	if teststring != actual {
		t.Errorf("expected = %v; actual = %v", string(teststring), rec.Body.String())
	}
}

func TestRssFeedController_GetRssFeedByName_Negative(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rssfeed/testfeed", nil)

	ctrl := gomock.NewController(t)
	mockhttpclient := utils.NewMockHttpClient(ctrl)
	defer ctrl.Finish()

	provider := NewMockProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/rssfeed")
	_ = NewRssFeedController(mapRouter, provider, mockhttpclient)

	provider.EXPECT().GetRssFeed(gomock.Any()).Return(&ent.RssFeed{}, errors.New("not found"))

	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusBadRequest, "got", rec.Code)
	}
}

func TestRssFeedController_GetRssFeedByName_HttpError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rssfeed/testfeed", nil)

	ctrl := gomock.NewController(t)
	mockhttpclient := utils.NewMockHttpClient(ctrl)
	defer ctrl.Finish()

	provider := NewMockProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/rssfeed")
	_ = NewRssFeedController(mapRouter, provider, mockhttpclient)

	feed := ent.RssFeed{
		ID:   1,
		URL:  "http://www.testfeed/rss.xml",
		Name: "testfeed",
	}

	provider.EXPECT().GetRssFeed("testfeed").Return(&feed, nil)
	mockhttpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusBadGateway,
	}, errors.New("dummy error"))
	router.ServeHTTP(rec, req)

	if http.StatusBadGateway != rec.Code {
		t.Error("expected ", http.StatusBadGateway, "got", rec.Code)
	}
}
