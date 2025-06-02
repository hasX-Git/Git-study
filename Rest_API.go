package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type model struct {
	Name      string `json:"name"`
	Vertcount uint32 `json:"vertcount"`
	Readyanim bool   `json:"readyanim"`
}

var models = []model{
	{Name: "kettle", Vertcount: 2932, Readyanim: true},
	{Name: "human", Vertcount: 12010, Readyanim: true},
	{Name: "bird", Vertcount: 15532, Readyanim: false},
}

func getModels(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models)
}

func addModel(context *gin.Context) {
	var newModel model

	if err := context.BindJSON(&newModel); err != nil {
		return
	}

	models = append(models, newModel)

	context.IndentedJSON(http.StatusCreated, newModel)
}

func getModelByName(name string) (*model, error) {
	for i, m := range models {
		if m.Name == name {
			return &models[i], nil
		}
	}

	return nil, errors.New("model not found")
}

func getModel(context *gin.Context) {
	name := context.Param("name")
	model, err := getModelByName(name)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "model not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, model)
}

func toggleModelStatus(context *gin.Context) {
	name := context.Param("name")
	model, err := getModelByName(name)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "model not found"})
		return
	}

	model.Readyanim = !model.Readyanim

	context.IndentedJSON(http.StatusOK, model)
}

func main() {
	router := gin.Default()
	router.GET("/models", getModels)
	router.GET("/models/:name", getModel)
	router.PATCH("/models/:name", toggleModelStatus)
	router.POST("/models", addModel)
	router.Run("localhost:9090")
}
