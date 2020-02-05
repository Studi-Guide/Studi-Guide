package models

import "image"

type Room struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Coordinates image.Rectangle `json:"coordinates"`
}