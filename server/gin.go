package server

import (
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	gin *gin.Engine
}

func InitGinServer() *ginServer {
	ginInstance := gin.Default()
	return &ginServer{gin: ginInstance}
}

func (g *ginServer) Run() {
	g.gin.Run(":8099")
}

func (g *ginServer) GetServerInstance() *gin.Engine {
	return g.gin
}
