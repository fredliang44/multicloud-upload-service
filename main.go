package main

import (
	"log"
	"net/http"

	"github.com/fredliang44/multicloud-upload-service/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}

		raw, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		defer raw.Close()

		err = handler.FileWriter(file.Filename, raw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"file": file.Filename, "size": file.Size})
	})
	router.Run(":8080")
}
