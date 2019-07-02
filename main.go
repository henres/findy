package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Ressource struct {
	ID          int    `json:"id" binding:"required"`
	Kind        string `json:"kind" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location"`
	Tags        []Tag  `json:"tags"`
}

type Tag struct {
	ID    int    `json:"id" binding:"required"`
	Value string `json:"string" binding:"required"`
}

type AddTag struct {
	Tags string `json:"tags" binding:"required"`
}

/** we'll create a list of jokes */
var ressources = []Ressource{
	Ressource{0, "person", "Henri d'Auvigny", "Bordeaux", []Tag{tags[0], tags[1]}},
	Ressource{0, "person", "Jean Dupont", "Bordeaux", []Tag{tags[2]}},
	Ressource{0, "doc", "Kubernetes", "Bordeaux", []Tag{tags[0], tags[1]}},
}

/** we'll create a list of jokes */
var tags = []Tag{
	Tag{0, "kubernetes"},
	Tag{1, "php"},
	Tag{2, "datawok"},
	Tag{3, "accounting"},
	Tag{4, "golang"},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file to load")
	} else {
		log.Print(".env loaded")
	}

	clientPath := os.Getenv("CLIENT_PATH")

	if _, err := os.Stat(clientPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "error: %v\nm ", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(clientPath, true)))

	api := router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/ressources", getRessourceHandler)
		api.GET("/ressources/:ressourceId", getRessourcesByIDHandler)
		api.POST("/ressources/:ressourceId/tags", addTagToRessourceHandler)
		api.GET("/tags", getTagHandler)
	}

	router.Run(":3000")
}

func getRessourceHandler(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	context.JSON(http.StatusOK, ressources)
}

func getRessourcesByIDHandler(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	var ressource *Ressource
	var ressourceID int
	var err error

	if ressourceID, err = strconv.Atoi(context.Param("ressourceId")); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ressource, err = getRessourceByID(ressourceID); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, ressource)
}

func getRessourceByID(ressourceID int) (*Ressource, error) {
	for _, ressource := range ressources {
		if ressource.ID == ressourceID {
			return &ressource, nil
		}
	}
	return nil, errors.New("Joke not found")
}

func getTagHandler(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	context.JSON(http.StatusOK, tags)
}

func addTagToRessourceHandler(context *gin.Context) {
	var ressource *Ressource
	var ressourceID int
	var err error

	if ressourceID, err = strconv.Atoi(context.Param("ressourceId")); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ressource, err = getRessourceByID(ressourceID); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	var jsonData AddTag
	context.BindJSON(&jsonData)
	tagsParam := strings.Split(jsonData.Tags, ",")

	for _, tagParam := range tagsParam {
		alreadyExistInTagList := false
		for key, tag := range tags {
			if tag.Value == tagParam {
				alreadyExistInObject := false
				for _, ressourceTag := range ressource.Tags {
					if ressourceTag == tags[key] {
						alreadyExistInObject = true
					}
				}
				if alreadyExistInObject == false {
					ressource.Tags = append(ressource.Tags, tags[key])
				}
				alreadyExistInTagList = true
			}
		}
		if alreadyExistInTagList == false {
			tags = append(tags, Tag{ID: len(tags), Value: tagParam})
			ressource.Tags = append(ressource.Tags, tags[len(tags)-1])
		}
	}

	fmt.Println(tags)
	fmt.Println(ressource)
	ressources[ressourceID] = *ressource

	context.JSON(http.StatusOK, ressource)
}
