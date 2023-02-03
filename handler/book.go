package handler

import (
	"log"

	"github.com/Suryaakula888/database"
	"github.com/Suryaakula888/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	Books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOk, Books)
}
func (h *Handler) SaveBooks(in *gin.Context) {
	book := models.Book{}
	in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOk, Books)
}
