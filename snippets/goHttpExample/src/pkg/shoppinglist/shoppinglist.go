package shoppinglist

import (
	"github.com/gorilla/mux"
	"httpExample/pkg/shoppinglist/controllers"
)

type ShoppingListApp struct {
	subRouterPrefix string
	packagePrefix string
	shoppingList *controllers.ShoppingListController
}


func (a *ShoppingListApp) Initialize(router *mux.Router, subRouterPrefix , packagePrefix string) {

	a.subRouterPrefix = subRouterPrefix
	a.packagePrefix = packagePrefix
	a.shoppingList = controllers.NewShoppingListController(router, subRouterPrefix, packagePrefix)

}

