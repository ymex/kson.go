# sonar
a Go package to search for JSON value and  interact with arbitrary JSON 。 用于搜索JSON值并与任意JSON交互的go 语言json 库。

## 语法

采用链式函数结构保持代码简洁，

### 别名 

查找的key可以使用:来定义别名。如，result:data ，那么在结果 集中，result 就是data 的别名。

### 数组 

查找数组时使用[] 

### 多级查找 

多级查找用`->`表示。 如，"result->books[2]->title" 表示查找 result 对象下数组books 的第二个元素对象的title.
