package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"studi-guide/pkg/shoppinglist/models"
	"studi-guide/pkg/shoppinglist/utils"
)

type ShoppingListController struct {
	router          *gin.RouterGroup
	subRouterPrefix string
	packagePrefix   string
	shoppingList    []models.ShoppingItem
}

func NewShoppingListController(router *gin.RouterGroup, subRouterPrefix, packagePrefix string) *ShoppingListController {

	l := ShoppingListController{
		router,
		subRouterPrefix,
		packagePrefix,
		[]models.ShoppingItem{models.ShoppingItem{Name: "Item 0", Description: ""}}}

	log.Print("Mapping static files..")
	utils.PrintMainDirectory()

	//l.router.Handle("/", http.StripPrefix(subRouterPrefix, http.FileServer(http.Dir(packagePrefix + "/views/"))))
	//gin.Default().LoadHTMLFiles(packagePrefix + "/views/*")

	//PrintCurrentDir()
	if _, err := os.Stat("./ionicframework/index.html"); err == nil {
		log.Print("Ionic folder found. Mapping files...")
		//l.router.StaticFile("/start/index.html", "./ionic/index.html")
	} else {
		l.router.StaticFS("/index/", http.Dir(packagePrefix+"/views/"))
	}

	l.router.POST("/list/", l.addItem)
	l.router.GET("/list/", l.getShoppingList)
	l.router.DELETE("/list/:name/", l.removeItem)
	l.router.GET("/list/:name/", l.getItem)

	return &l
}

func testMethod(c *gin.Context) {
	//w.Write("Hello testMethod!")
	fmt.Println("writer:", c.Writer)
	fmt.Println("request:", c.Request)
	fmt.Fprint(c.Writer, "Hello testMethod!")
}

// GetShoppingList godoc
// @Summary Get Shopping List
// @Description Gets all shopping items
// @ID get-shopping-list
// @Accept  json
// @Produce  json
// @Tags ShoopingListController
// @Success 200 {array} models.ShoppingItem
// @Router /shoppinglist/list/ [get]
func (l *ShoppingListController) getShoppingList(c *gin.Context) {
	c.JSON(http.StatusOK, l.shoppingList)
}

func (l *ShoppingListController) getItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(c.Request)["name"]

	for _, item := range l.shoppingList {
		if item.Name == name {
			c.JSON(http.StatusOK, item)
		}
	}
}

func (l *ShoppingListController) addItem(c *gin.Context) {
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	var item models.ShoppingItem
	json.Unmarshal(reqBody, &item)

	l.shoppingList = append(l.shoppingList, item)
}

func (l *ShoppingListController) removeItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(c.Request)["name"]
	for index, item := range l.shoppingList {
		if item.Name == name {
			l.shoppingList = append(l.shoppingList[:index], l.shoppingList[index+1:]...)
		}
	}
}
