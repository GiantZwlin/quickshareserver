package ent

import "gorm.io/gorm"

type Gist struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}
