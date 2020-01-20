package shoppinglist

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ShoppingItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ShoppingList struct {
	router       *mux.Router
	shoppingList []ShoppingItem
}

func (l *ShoppingList) Initialize(router *mux.Router) {
	l.shoppingList = append(l.shoppingList, ShoppingItem{Name: "Item 0", Description: ""})
	l.router = router

	log.Print("Mapping static files..")
	printMainDirectory()
	//f exist, err := exists("./static"); !exist {
	//	log.Fatal(err)
	//	log.Fatal("static folder does not exist")
	// }
	l.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	l.router.Handle("/", http.RedirectHandler("/static/", 301))
	l.router.HandleFunc("/shoppinglist", l.AddItem).Methods("POST")
	l.router.HandleFunc("/shoppinglist", l.GetShoppingList)
	l.router.HandleFunc("/shoppinglist/{name}", l.RemoveItem).Methods("DELETE")
	l.router.HandleFunc("/shoppinglist/{name}", l.GetItem)
}

func (l *ShoppingList) Run(addr string) {
	http.ListenAndServe(addr, l.router)
}

func (l *ShoppingList) GetShoppingList(w http.ResponseWriter, r *http.Request) {
	log.Print("[ShoppingList] Request shoppingList received")
	json.NewEncoder(w).Encode(l.shoppingList)
}

func (l *ShoppingList) GetItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	for _, item := range l.shoppingList {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func (l *ShoppingList) AddItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item ShoppingItem
	json.Unmarshal(reqBody, &item)

	l.shoppingList = append(l.shoppingList, item)
}

func (l *ShoppingList) RemoveItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	for index, item := range l.shoppingList {
		if item.Name == name {
			l.shoppingList = append(l.shoppingList[:index], l.shoppingList[index+1:]...)
		}
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func printMainDirectory() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	log.Print(exPath)
}
