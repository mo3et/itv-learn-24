# Reflect

通过反射机制，在执行期能获取 Type、Value，也能进行更改

## TypeOf

```go
v:= reflect.TypeOf(x interface{})
v.Name() // Get type Name
v.Kind()  // Get Kind

type person struct{}
reflect.TypeOf().Name() // person
reflect.TypeOf().Kind() // struct

```

## ValueOf
- Get Value
```go
v := reflect.ValueOf(x Any)
k := v.kind()

c:= reflect.ValueOf(10) // 将int的原始值转为Reflect.Value 类型
fmt.Printf("type c :%T\n", c) // type c :reflect.Value

```

- Set Value
```go
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

var a int64 = 100
reflectSetValue(&a) //必须是传地址值才能修改变量值，在反射中是使用Elem() 获取指针对应值

```

- `IsNil` & `IsValid()`
`IsNil()` 必须是**通道、函数、接口、Map、指针、Slice**，否则panic

`IsValid()` 除了**IsValid、String、Kind**以外的方法都panic
```go
    // *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())

	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

```

## struct Reflect

任意值通过`reflect.TypeOf()`获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（`reflect.Type`）的`NumField()`和`Field()`方法获得结构体成员的详细信息。

- StructField类型
```go
type StructField struct {
    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name    string
    PkgPath string
    Type      Type      // 字段的类型
    Tag       StructTag // 字段的标签
    Offset    uintptr   // 字段在结构体中的字节偏移量
    Index     []int     // 用于Type.FieldByIndex时的索引切片
    Anonymous bool      // 是否匿名字段
}

```

通过反射得到结构体后通过索引依次获取器字段信息，也可以通过字段名去获取指定字段信息。
```go
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

```

编写函数 `printMethod(s interface{})` 来遍历打印s包含的方法
```go
// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
```