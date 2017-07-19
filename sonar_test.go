package sonar

import (
	"testing"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func TestN(t *testing.T) {
	bson :=[]byte(`{
		"name":"ymex",
		"person":{
			"sex":"man",
			"marry":"hellow"
		},
		"age":4.2,
		"car":"T",
		"boxes":[
			 [
			  {"hi":"hello","say":"bye"},
			  {"hi":"hello","say":"bye"}
			 ],[
			  {"hi":"hello","say":"bye"},
			  {"hi":"hello","say":"bye"}
			 ],[
			  {"hi":"hello","say":"bye"},
			  {"hi":"hello","say":"bye"}
			 ]
		]
	}`)
	kson := NewSonar(bson).Find("age","last:person->sex","boxes:boxes[0][1]")
	fmt.Println(kson.GotFirst().ToFloat())
	fmt.Println(kson.GotPosition(1).ToString())
	fmt.Println(kson.GotLast().Map())

	arrjson := []byte(`[{
		  "width":24.0,
		  "height":50.82,
		  "color":"red"
		 },{
		  "width":93.2,
		  "height":234.19,
		  "color":"yellow"
		 }]`)

	k := NewSonar(arrjson).Find("[1]->width","[0]->color")
	fmt.Println(k.GotFirst().ToFloat())
	fmt.Println(k.GotLast().ToString())
}


