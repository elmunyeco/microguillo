package main

import (
	"strconv"

	gin "github.com/gin-gonic/gin"
)

type Culo struct {
	ID    int    `json:"id"`
	Roto  string `json:"roto"`
	FooId int    `json:"foo_id"`
}

var culos = []Culo{
	{ID: 1, Roto: "baz"},
	{ID: 2, Roto: "qux", FooId: 1},
	{ID: 3, Roto: "quux"},
}

func listCulo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, culos)
	}
}

func getCulo(c *gin.Context) {
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		for _, a := range culos {
			if a.ID == id {
				c.JSON(200, a)
				return
			}
		}
		c.JSON(404, gin.H{"error": "culo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid culo id"})
	}
}

func postCulo(c *gin.Context) {
	var json Culo
	if c.Bind(&json) == nil {
		culos = append(culos, json)
		c.JSON(200, json)
	}
}

func putCulo(c *gin.Context) {
	var json Culo
	if c.Bind(&json) == nil {
		for i, a := range culos {
			if a.ID == json.ID {
				culos[i] = json
				c.JSON(200, gin.H{"success": "culo updated"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "culo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid culo id"})
	}
}

func deleteCulo(c *gin.Context) {
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		for i, a := range culos {
			if a.ID == id {
				culos = append(culos[:i], culos[i+1:]...)
				c.JSON(200, gin.H{"success": "culo deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "culo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid culo id"})
	}
}

func main() {
	r := gin.Default()
	r.GET("/culo", listCulo())
	r.GET("/culo/:id", getCulo)
	r.POST("/culo", postCulo)
	r.PUT("/culo/:id", putCulo)
	r.DELETE("/culo/:id", deleteCulo)
	r.Run("localhost:8080")
}
