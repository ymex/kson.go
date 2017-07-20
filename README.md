# sonar
a Go package to search for JSON value and  interact with arbitrary JSON 。 用于搜索JSON值并与任意JSON交互的go 语言json 库。

## 语法

采用链式函数结构保持代码简洁，

### 别名 

查找的key可以使用:来定义别名。如，`result:data` ，那么在结果 集中，result 就是data 的别名。

### 数组 

查找数组时使用`[] `

### 多级查找 

多级查找用`->`表示。 如，"result->books[2]->title" 表示查找 result 对象下数组books 的第二个元素对象的title.

###条件查找[暂未支持]
条件查找仅支持 `==`,`!=`,`>`,`<`,`>=`,`<=` ,查找内容放在`{}`中间。如`students->{age>24}`

```
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
	fmt.Println(sonar.Got("message").ToString())//>>success
	fmt.Println(sonar.GotLast().Interface())//map[name:Celina age:17]
	fmt.Println(sonar.Got("result").Interface())//map[name:Celina age:17]
```
