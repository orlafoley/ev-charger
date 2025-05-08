package main

import (
	"log"
	"net/http"

	"personweb/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// gin.Default
	// is used as the default router
	r := gin.Default()
	models.ConnectDatabase()

	//r.LoadHTMLGlob("templates/*")
	//API v1
	v1 := r.Group("/api/vi")
	{
		v1.GET("getMember", getMember)
		v1.GET("getMember/:id", getMemberId)
		v1.POST("addMember", addMember)
		v1.PUT("updateMember/:id", updateMember)
		v1.DELETE("DeleteMember/:id", deleteMember)
		v1.OPTIONS("member", options)
	}

	//Router Config
	r.Run("localhost:5173")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getMember(c *gin.Context) {
	person, err := models.GetAllMember()
	checkErr(err)

	if err != nil { // Check for errors first
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve members"})
		return
	}

	if len(person) == 0 { // Check if the slice is empty (no data found)
		c.JSON(http.StatusOK, gin.H{"message": "No members found"}) // Or a 404 Not Found if that's more appropriate for your API design
	} else {
		c.JSON(http.StatusOK, gin.H{"data": person})
	}
}

func getMemberId(c *gin.Context) {
	// id comes from the URI of which we pass through from link
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getMemberID " + id})
}

func addMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addMembers"})

}

func updateMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateMembers"})
}

func deleteMember(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deleteMember " + id})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options"})
}
