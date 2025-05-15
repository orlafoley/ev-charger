package main

import (
	"log"
	"net/http"
	"src/backend/src/backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// gin.Default
	// is used as the default router
	r := gin.Default()

	//Enable CORS for Cross-Origin Resource Sharing
	r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Connection to DB
	models.ConnectDatabase()

	// Set handler for /api
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API SET"})
	})

	//r.LoadHTMLGlob("templates/*")
	//API v1
	v1 := r.Group("/api")
	{
		v1.GET("getMember", getMember)
		v1.GET("getMember/:id", getMemberId)
		v1.POST("addMember", addMember)
		v1.PUT("updateMember/:id", updateMember)
		v1.DELETE("DeleteMember/:id", deleteMember)
		v1.OPTIONS("member", options)
		v1.POST("bookings", quickBooking)
		v1.GET(("getAllBookings"), getAllBooking)
	}

	//Router Config
	r.Run("localhost:8080")
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

func quickBooking(c *gin.Context) {
	var request models.BookingRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	booking := models.Bookings{
		Name:     request.Name,
		Email:    request.Email,
		Date:     request.Date,
		Time:     request.Time,
		Duration: request.Duration,
	}

	err := models.InsertNewBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert booking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking successful"})

}

func getAllBooking(c *gin.Context) {
	booking, err := models.GetAllBooking()
	checkErr(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve members"})
		return
	}

	if len(booking) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Bookings Yet"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": booking})
	}

}
