package general

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ProxyController struct {
	router *gin.RouterGroup
}

func NewProxyController(router *gin.RouterGroup) error {
	b := ProxyController{
		router: router,
	}

	b.router.GET("/:url", b.GetProxyRequest)
	return nil
}

func (c ProxyController) GetProxyRequest(context *gin.Context) {
	// url ist base64 url encodiert
	url := context.Param("url")
	encodedUrl, err := base64.URLEncoding.DecodeString(url)
	if err != nil {
		fmt.Println("GetProxyRequest failed with error", err)
		statusCode := http.StatusBadRequest
		context.JSON(statusCode, gin.H{
			"code":    statusCode,
			"message": err.Error(),
		})

		return
	}

	urlString := string(encodedUrl)
	fmt.Printf("GetProxyRequest with url: %v\n", urlString)
	resp, err := http.Get(urlString)
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
