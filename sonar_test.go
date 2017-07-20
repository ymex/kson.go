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
	sonar := NewSonar(b).Find("code","last:data->mileage","message","result:data->passenger->students[0][1]")

	fmt.Println(sonar.GotFirst().ToInt()) //>>200
	fmt.Println(sonar.GotPosition(1).ToFloat())//>>253.56
	fmt.Println(sonar.Got("last").ToFloat())//>>253.56
	fmt.Println(sonar.Got("message").ToString())//>>
	fmt.Println(sonar.GotLast().Interface())//map[name:Celina age:17]
	fmt.Println(sonar.Got("result").Interface())//map[name:Celina age:17]


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
	fmt.Println(k.GotFirst().ToFloat())//>>93.2
	fmt.Println(k.GotLast().ToString())//>>red
}


