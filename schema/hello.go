
// schema/hello.go
package schema

import (

	"github.com/graphql-go/graphql"
	"gql/model"
)

// 定义查询对象的字段，支持嵌套
var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Description:"唯一键",
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Description:`""之类的零值会被忽略`,

		},
		"test_col": &graphql.Field{
			Type:graphql.String,
			Description:"无意义的测试字段",
		},
	},
})



// 处理查询请求
var queryHello = graphql.Field{
	Name: "QueryHello",
	Description: "Query Hello",
	Type: graphql.NewList(helloType),
	// Args是定义在GraphQL查询中支持的查询字段，
	// 可自行随意定义，如加上limit,start这类
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
			Description:"唯一键",
			DefaultValue:0,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
			Description:"姓名",
		},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		// Args里面定义的字段在p.Args里面，对应的取出来
		// 因为是interface{}的值，需要类型转换，可参考类型断言(type assertion): https://zengweigang.gitbooks.io/core-go/content/eBook/11.3.html
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)


		// 调用Hello这个model里面的Query方法查询数据
		return (&model.Hello{}).Query(id, name),nil
	},
}

