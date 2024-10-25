package routes

import (
	"docflow-backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getDocsForUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Println("Binding error:", err) // Log the error for debugging
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user ID."})
		return
	}

	docs, err := models.GetAllDocsForUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get docs for user."})
		return
	}

	context.JSON(http.StatusOK, docs)
}

func getDocByID(context *gin.Context) {
	docId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse doc ID."})
		return
	}

	doc, err := models.GetDocByID(docId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get docs by ID."})
		return
	}

	context.JSON(http.StatusOK, doc)
}

func generateDocForUser(context *gin.Context) {
	var doc models.Doc
	err := context.ShouldBindJSON(&doc)
	if err != nil {
		log.Println("Binding error:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")
	doc.UserID = userId

	err = doc.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create a doc."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Doc created", "doc": doc})
}