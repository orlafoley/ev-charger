package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default is used as the default router
	r := gin.Default()

	//API v1
	v1 := r.Group("/api/vi")
	{
		v1.GET("member", getMember)
		v1.GET("member/:id", getMemberId)
		v1.POST("member", addMember)
		v1.PUT("member/:id", updateMember)
		v1.DELETE("member/:id", deleteMember)
		v1.OPTIONS("member", options)
	}

	//Router Config
	r.Run("localhost:5173")
}

func getMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getMembers Called"})
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
