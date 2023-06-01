package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

// task print current time and simple string.
func task() {
	currentTime := time.Now()
	fmt.Println(
		fmt.Sprintf(
			"[%s]",
			currentTime.Format("2006-01-02 15:04:05"),
		),
		"I am running a task.",
	)
}

// main starts a scheduler and a web service.
func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(3).Seconds().Do(task)
	// s.StartBlocking()
	s.StartAsync()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run(":8080")
}
