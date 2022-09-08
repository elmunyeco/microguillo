package main

import (
	"strconv"

	gin "github.com/gin-gonic/gin"
)

type Foo struct {
	ID  int    `json:"id"`
	Bar string `json:"bar"`
}

var foos = []Foo{
	{ID: 1, Bar: "baz"},
	{ID: 2, Bar: "qux"},
	{ID: 3, Bar: "quux"},
}

func listFoo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, foos)
	}
}

func getFoo(c *gin.Context) {
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		for _, a := range foos {
			if a.ID == id {
				c.JSON(200, a)
				return
			}
		}
		c.JSON(404, gin.H{"error": "foo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid foo id"})
	}
}

func postFoo(c *gin.Context) {
	var json Foo
	if c.Bind(&json) == nil {
		foos = append(foos, json)
		c.JSON(200, json)
	}
}

func putFoo(c *gin.Context) {
	var json Foo
	if c.Bind(&json) == nil {
		for i, a := range foos {
			if a.ID == json.ID {
				foos[i] = json
				c.JSON(200, gin.H{"success": "foo updated"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "foo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid foo id"})
	}
}

func deleteFoo(c *gin.Context) {
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		for i, a := range foos {
			if a.ID == id {
				foos = append(foos[:i], foos[i+1:]...)
				c.JSON(200, gin.H{"success": "foo deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "foo not found"})
	} else {
		c.JSON(400, gin.H{"error": "invalid foo id"})
	}
}

func main() {
	r := gin.Default()
	r.GET("/foo", listFoo())
	r.GET("/foo/:id", getFoo)
	r.POST("/foo", postFoo)
	r.PUT("/foo/:id", putFoo)
	r.DELETE("/foo/:id", deleteFoo)
	r.Run("localhost:8080")
}
