package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"httpExample/pkg/shoppinglist/models"
	"httpExample/pkg/shoppinglist/utils"
	"io/ioutil"
	"log"
	"net/http"
)
type ShoppingListController struct {
	router       *mux.Router
	subRouterPrefix string
	packagePrefix string
	shoppingList []models.ShoppingItem
}

func NewShoppingListController(router *mux.Router, subRouterPrefix, packagePrefix string) (*ShoppingListController) {

	l := ShoppingListController{
		router,
		subRouterPrefix,
		packagePrefix,
		[]models.ShoppingItem{models.ShoppingItem{Name: "Item 0", Description: ""}}}


	log.Print("Mapping static files..")
	utils.PrintMainDirectory()
	//f exist, err := exists("./static"); !exist {
	//	log.Fatal(err)
	//	log.Fatal("static folder does not exist")
	// }

	l.router.Handle("/", http.StripPrefix(subRouterPrefix, http.FileServer(http.Dir(packagePrefix + "/views/"))))
	l.router.HandleFunc("/", testMethod)
	l.router.HandleFunc("/list/", l.addItem).Methods("POST")
	l.router.HandleFunc("/list/", l.getShoppingList)
	l.router.HandleFunc("/list/{name}/", l.removeItem).Methods("DELETE")
	l.router.HandleFunc("/list/{name}/", l.getItem)

	return &l
}

func testMethod(w http.ResponseWriter, r *http.Request) {
	//w.Write("Hello testMethod!")
	fmt.Println("writer:", w)
	fmt.Println("request:", r)
	fmt.Fprint(w, "Hello testMethod!")
}

func (l *ShoppingListController) getShoppingList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(l.shoppingList)
}

func (l *ShoppingListController) getItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	for _, item := range l.shoppingList {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func (l *ShoppingListController) addItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item models.ShoppingItem
	json.Unmarshal(reqBody, &item)

	l.shoppingList = append(l.shoppingList, item)
}

func (l *ShoppingListController) removeItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	for index, item := range l.shoppingList {
		if item.Name == name {
			l.shoppingList = append(l.shoppingList[:index], l.shoppingList[index+1:]...)
		}
	}
}