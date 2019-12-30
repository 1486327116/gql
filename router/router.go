package router

import (
	"github.com/gin-gonic/gin"
	"gql/controller/graphql"
)
func init() {
	Router= gin.Default()
}

var Router *gin.Engine

func SetupRouter()  {
	Router.POST("/gql",graphql.ApiHandler())
	Router.GET("/gql",graphql.ApiHandler())
}