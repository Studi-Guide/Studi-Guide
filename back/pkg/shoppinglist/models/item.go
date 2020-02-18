package models

type ShoppingItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Database functions etc. here ...
// e.g. to save item, load item, ...
