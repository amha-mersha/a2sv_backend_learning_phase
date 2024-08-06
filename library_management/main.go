package main

import (
	"library_management/controllers"
	"library_management/models"
)

func main() {
	library := models.Library{Books: make(map[int64]*models.Book), Member: make(map[int64]*models.Member)}
	controllers.Run(&library)
}
