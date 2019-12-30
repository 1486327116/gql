package main

import (
	"gql/router"
)

//this is a sever for gql,just a try for me.
//use gin and go-graphql
func main() {

	engine:=router.Router
	router.SetupRouter()
	_ = engine.Run()
}
