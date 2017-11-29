package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gopher struct {
	Name string `json: "name"`
	Job  string `json: "job"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/list", listGopher)
	r.POST('list', newGopher)

	r.Run(":9999")
}

func listGopher(c *gin.Context) {
	var list []Gopher
	list = append(list, Gopher{"Song", "SDE"})
	list = append(list, Gopher{"Wang", "SDE"})
	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}

func newGopher(c *gin.Context) {
	var body struct {
		Data map[string]interface{}
	}	

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}	

	name, _ := body.Data["name"].(string)
	job, _ := body.Data["job"].(string)

	if name == "" || job == "" {
		c.JSON(http.StatusConflict, gin.H{"error" : "Invalid Data"})
		return
	}

	gopher := Gopher{Name: name, Job: job}
	c.JSON(http.StatusCreated, gin.H{"data": gopher.Map()})
}

func (gopher Gopher) Map() map[stiring]interface{} {
	return map[string]interface{} {
		"name": gopher.Name,
		"job": gopher.Job,
	}
}
