package shoppinglist

import (
	"github.com/gin-gonic/gin"
	"httpExample/pkg/shoppinglist/controllers"
)

type ShoppingListApp struct {
	subRouterPrefix string
	packagePrefix   string
	shoppingList    *controllers.ShoppingListController
}

func (a *ShoppingListApp) Initialize(router *gin.RouterGroup, subRouterPrefix, packagePrefix string) {

	a.subRouterPrefix = subRouterPrefix
	a.packagePrefix = packagePrefix
	a.shoppingList = controllers.NewShoppingListController(router, subRouterPrefix, packagePrefix)

}
