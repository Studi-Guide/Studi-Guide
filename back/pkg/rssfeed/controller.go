package rssfeed

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"studi-guide/pkg/utils"
)

type Controller struct {
	router      *gin.RouterGroup
	rssProvider Provider
	httpClient  utils.HttpClient
}

func NewRssFeedController(router *gin.RouterGroup, provider Provider, client utils.HttpClient) error {
	b := Controller{
		router:      router,
		rssProvider: provider,
		httpClient:  client,
	}

	b.router.GET("/:rssFeedId", b.GetRssFeed)
	return nil
}

// GetRssFeedByName godoc
// @Summary Get RssFeed by a certain name
// @Description Get one RssFeed by name
// @ID get-rssfeed
// @Accept  json
// @Produce  plain
// @Tags RssFeedController
// @Param rssFeedId path string true "rssFeed of the campus"
// @Success 200
// @Router /rssfeed/{rssFeedId} [get]
func (c Controller) GetRssFeed(context *gin.Context) {
	rssFeedID := context.Param("rssFeedId")
	feed, err := c.rssProvider.GetRssFeed(rssFeedID)
	if err != nil {
		returnErrorCode(err, http.StatusBadRequest, context)
		return
	}

	uri, _ := url.Parse(feed.URL)
	resp, err := c.httpClient.Do(&http.Request{
		Method: "GET",
		URL:    uri,
	})

	if err != nil {
		statusCode := http.StatusBadRequest
		if resp != nil {
			statusCode = resp.StatusCode
		}

		returnErrorCode(err, statusCode, context)
		return
	}

	contentType := "application/json"
	if resp.Header != nil {
		cntTpye := resp.Header.Get("Content-Type")
		if len(cntTpye) > 0 {
			contentType = cntTpye
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		returnErrorCode(err, http.StatusInternalServerError, context)
		return
	}

	context.Data(http.StatusOK, contentType, body)
}

func returnErrorCode(err error, statusCode int, context *gin.Context) {
	fmt.Println("GetRssFeed failed with error", err)
	context.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": err.Error(),
	})
}
