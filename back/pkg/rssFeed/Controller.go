package rssFeed

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	router      *gin.RouterGroup
	rssProvider Provider
}

func NewRssFeedController(router *gin.RouterGroup, provider Provider) error {
	b := Controller{
		router:      router,
		rssProvider: provider,
	}

	b.router.GET("/:rssFeedId", b.GetRssFeed)
	return nil
}

func (c Controller) GetRssFeed(context *gin.Context) {
	// url ist base64 url encodiert
	rssFeedID := context.Param("rssFeedId")
	feed, err := c.rssProvider.GetRssFeed(rssFeedID)

	resp, err := http.Get(feed.URL)
	if err != nil {
		fmt.Println("GetProxyRequest failed with error", err)
		statusCode := http.StatusBadRequest
		if resp != nil {
			statusCode = resp.StatusCode
		}
		context.JSON(statusCode, gin.H{
			"code":    statusCode,
			"message": err.Error(),
		})

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

	context.Data(http.StatusOK, contentType, body)
}
