package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/cperdiansyah/golang-h8-dts/final-project/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	router.UserRouter(r)
	router.PhotoRouter(r)
	router.CommentRouter(r)
	router.SocialMediaRouter(r)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
