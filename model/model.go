package model

import (
	"github.com/gin-gonic/gin"
)

type Hello struct {

}

func (p *Hello) Query(id int , name string) interface{} {
	data:=[]gin.H{
		{"id":11, "name":"anni","test_col":"this test col11"},
		{"id":12, "name":"jeff","test_col":"this test col12"},
		{"id":13, "name":"kk","test_col":"this test col13"},
		{"id":14, "name":"loft","test_col":"this test col14"},
		{"id":15, "name":"ken","test_col":"this test col15"},
		{"id":16, "name":"kimi","test_col":"this test col16"},
		{"id":17, "name":"jack","test_col":"this test col17"},
		{"id":18,"test_col":"this test col18"},
	}

	var result []gin.H
	if id!=0 {
		if  name ==""{
			for _, v := range data {
				if v["id"] == id {
					result=append(result, v)
				}
			}
		}else {
			for _, v := range data {
				if v["id"] == id && v["name"]==name{
					result=append(result, v)
				}
			}
		}


	}else {
		result =data
	}

	return result
}