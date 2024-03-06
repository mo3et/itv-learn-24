# interface

## 简单用法

```go
 type 接口名 interface{
     方法名1(参数列表1) 返回值列表1
     方法名2(参数列表2) 返回值列表2
     …
 }
```

如何要一个 struct 可以使用接口？
只要让 struct 实现接口里面的全部方法即可，并且接收者为 struct.

-------

## 面试题

> 当涉及到编程面试题时，通常涉及到一些实际场景中接口的应用。以下是一些可能的 Golang 接口的面试题，希望能够帮助你巩固对接口的理解：
> 
> ### 1. 接口的嵌套与类型断言
> 
> **题目：**
> 定义两个接口，`Reader` 和 `Writer`，每个接口包含一个方法分别为 `Read` 和 `Write`。然后定义一个结构体 `DataProcessor`，实现这两个接口。在 `main` 函数中，创建一个 `DataProcessor` 实例，然后通过类型断言分别调用 `Read` 和 `Write` 方法。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义接口 Reader
> // ...
> 
> // 定义接口 Writer
> // ...
> 
> // 定义结构体 DataProcessor，并实现接口 Reader 和 Writer
> // ...
> 
> func main() {
>     // 创建 DataProcessor 实例
>     // ...
> 
>     // 使用类型断言调用 Read 方法
>     // ...
> 
>     // 使用类型断言调用 Write 方法
>     // ...
> }
> ```
> 
> ### 2. 空接口的应用
> 
> **题目：**
> 定义一个函数 `PrintInfo`，该函数接受一个空接口参数，并打印出传入参数的类型和值。在 `main` 函数中测试该函数，传入不同类型的参数。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义函数 PrintInfo
> // ...
> 
> func main() {
>     // 测试 PrintInfo 函数，传入不同类型的参数
>     // ...
> }
> ```
> 
> ### 3. 多态性与接口的应用
> 
> **题目：**
> 定义一个接口 `Shape` 包含 `Area` 方法，表示计算形状的面积。然后实现两个结构体 `Circle` 和 `Rectangle` 分别实现这个接口。在 `main` 函数中，创建一个包含多个不同形状的切片，并计算它们的总面积。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义接口 Shape
> // ...
> 
> // 定义结构体 Circle 和 Rectangle，实现接口 Shape
> // ...
> 
> func main() {
>     // 创建包含多个不同形状的切片
>     // ...
> 
>     // 计算切片中形状的总面积
>     // ...
> }
> ```
> 
> 
> 
> ### 4. 接口的组合与嵌套
> 
> **题目：**
> 定义两个接口 `Fetcher` 和 `Parser`，分别包含 `Fetch` 和 `Parse` 方法。然后创建一个新的接口 `Processor`，嵌套 `Fetcher` 和 `Parser`。实现一个结构体 `WebProcessor`，实现 `Processor` 接口的所有方法。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义接口 Fetcher
> // ...
> 
> // 定义接口 Parser
> // ...
> 
> // 定义接口 Processor，组合嵌套 Fetcher 和 Parser
> // ...
> 
> // 定义结构体 WebProcessor，并实现 Processor 接口的所有方法
> // ...
> 
> func main() {
>     // 创建 WebProcessor 实例
>     // ...
> }
> ```
> 
> ### 5. 自定义错误类型与接口
> 
> **题目：**
> 定义一个自定义错误类型 `ValidationError`，实现 `error` 接口。然后编写一个函数 `ValidateData`，该函数接受一个参数并返回 `ValidationError` 类型的错误。在 `main` 函数中测试该函数。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义 ValidationError 错误类型
> // ...
> 
> // 定义函数 ValidateData，返回 ValidationError 类型的错误
> // ...
> 
> func main() {
>     // 测试 ValidateData 函数
>     // ...
> }
> ```
> 
> ### 6. 接口的空方法与可选实现
> 
> **题目：**
> 定义一个接口 `OptionalMethod` 包含一个可选的方法 `Optional() bool`。然后创建一个结构体 `MyStruct`，实现了该接口的空方法。在 `main` 函数中，创建一个变量并检查它是否实现了 `OptionalMethod` 接口。
> 
> **代码框架：**
> 
> ```go
> package main
> 
> import "fmt"
> 
> // 定义接口 OptionalMethod 包含一个可选的方法 Optional() bool
> // ...
> 
> // 定义结构体 MyStruct，实现 OptionalMethod 接口的空方法
> // ...
> 
> func main() {
>     // 创建变量并检查是否实现 OptionalMethod 接口
>     // ...
> }
> ```
> 
> 



-----
# References
- [理解 Go interface 的 5 个关键点- 三月沙](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)