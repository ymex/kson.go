package kson

import (
	"testing"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func TestN(t *testing.T) {
	b :=[]byte(`{
		 "code":200,
		 "message":"success",
		 "data":{
		 	"busId":24,
		 	"mileage":253.56,
		 	"passenger":{
		 		"students":[
					[{"name":"Bili","age":16},{"name":"Celina","age":17},{"name":"Serafina","age":18}],
					[{"name":"Abby","age":19},{"name":"Amaris","age":20},{"name":"Fiona","age":21}],
					[{"name":"Snow","age":24},{"name":"Muse","age":23},{"name":"Gina","age":22}]
		 		],
		 		"teachers":[
		 		 {
		 		 	"name":"Tom",
		 		 	"age":37,
		 		 	"teach":"math"
		 		 },
		 		  {
		 		 	"name":"Li",
		 		 	"age":37,
		 		 	"teach":"math"
		 		 }
		 		]
		 	}
		 }
		}`)
	kson := Unmarshal(b).Find("code","last:data->mileage","message","result:data->passenger->students[0][1]")

	fmt.Println(kson.GotFirst().ToInt())        //>>200
	fmt.Println(kson.GotPosition(1).ToFloat())  //>>253.56
	fmt.Println(kson.Got("last").ToFloat())     //>>253.56
	fmt.Println(kson.Got("message").ToString()) //>>
	fmt.Println(kson.GotLast().Interface())     //map[name:Celina age:17]
	fmt.Println(kson.Got("result").Interface()) //map[name:Celina age:17]


	arrjson := []byte(`[{
		  "width":24.0,
		  "height":50.82,
		  "color":"red"
		 },{
		  "width":93.2,
		  "height":234.19,
		  "color":"yellow"
		 }]`)

	k := Unmarshal(arrjson).Find("[1]->width","[0]->color")
	fmt.Println(k.GotFirst().ToFloat())//>>93.2
	fmt.Println(k.GotLast().ToString())//>>red
}


