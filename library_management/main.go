package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	library := models.Library{Books: make(map[int64]*models.Book), Member: make(map[int64]*models.Member)}
	library_management := services.LibraryManager{Library: &library}
	controllers.Run(&library_management)
}
