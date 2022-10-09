package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Server interface {
	SetRoute()
	StartServer() error
}
type GinServer struct {
	url    string
	port   int
	router *gin.Engine
}

func NewGinServer(url string, port int) Server {
	return &GinServer{url: url, port: port, router: gin.Default()}
}

func (g *GinServer) StartServer() error {
	return g.router.Run(fmt.Sprintf("%s:%d", g.url, g.port))
}
func (g *GinServer) SetRoute() {
	g.router.GET("/user/:id", getting)
	g.router.POST("/user", posting)
	g.router.PUT("/user/:id", putting)
	g.router.DELETE("/user/:id", deleting)
}

var m sync.Map

type User struct {
	Name string
	ID   string
}

func deleting(c *gin.Context) {
	query := c.Query("id")
	m.Delete(query)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func putting(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	m.Store(user.ID, user)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func posting(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	m.Store(user.ID, user)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func getting(c *gin.Context) {
	query := c.Query("id")
	value, ok := m.Load(query)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, value)
}
