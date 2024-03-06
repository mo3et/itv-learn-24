package codeq

// 定义接口 Fetcher
type Fetcher interface {
	Fetch() string
}

// 定义接口 Parser
type Parser interface {
	Parse(data string) string
}

// 定义接口 Processor，组合嵌套 Fetcher 和 Parser
type Processor interface {
	Fetcher
	Parser
}

// 定义结构体 WebProcessor，并实现 Processor 接口的所有方法
type WebProcessor struct {
	Data string
}

func (wp *WebProcessor) Fetch() string {
	wp.Data = "json send."
	return wp.Data
}

func (wp WebProcessor) Parse(data string) string {
	return "Parse web data:" + data
}

func Q4() {
}
