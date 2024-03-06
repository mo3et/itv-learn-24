package codeq

import (
	"fmt"
	"log"
)

// 定义接口 Reader
type Reader interface {
	Read() string
}

// 定义接口 Writer
type Writer interface {
	Write(data string)
}

// 定义结构体 DataProcessor，并实现接口 Reader 和 Writer
type DataProcessor struct {
	Data string
}

func (dp DataProcessor) Read() string {
	return dp.Data
}

func (dp *DataProcessor) Write(data string) {
	dp.Data = data
}

func Q1() {
	// 创建 DataProcessor 实例
	dp := &DataProcessor{Data: "I use the interface."}

	// 使用类型断言调用 Read 方法
	if reader, ok := interface{}(dp).(Reader); ok {
		fmt.Println("Reading", reader.Read())
	}
	reader := dp.Read()
	log.Print(reader)

	// 使用类型断言调用 Write 方法
	if writer, ok := interface{}(dp).(Writer); ok {
		writer.Write("New Data")
		fmt.Println("Writing:", dp.Data)
	}

	dp.Write("Pass Interface.")
	log.Print(dp.Data)
}
