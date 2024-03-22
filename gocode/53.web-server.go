package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method:", r.Method)
	fmt.Println("Enter it!")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tmpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 执行 登录数据进行登录逻辑判断
		_ = r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if pwd := r.Form.Get("username"); pwd == "123456" { // 验证密码
			fmt.Fprintf(w, "welcome to login, %s!", r.Form.Get("username")) // 输出到客户端信息
		} else {
			fmt.Fprintf(w, "input the pass!!") // 输出到客户端信息
		}
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm() // 3. 解析参数，默认中不会解析
	// 解析URL传递的参数，对于post则解析对应包的主体 (Request body)
	fmt.Println(r.Form)             // 4.输出到服务器的打印信息
	fmt.Println("Path", r.URL.Path) // 输出路径
	fmt.Println("Host", r.Host)     // 输出端口
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello Web,%s!", r.Form.Get("name")) // 5. 写到w的是输出到客户端的内容
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil) // 设置监听接口
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
