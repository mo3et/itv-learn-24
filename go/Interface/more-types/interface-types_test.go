// package moretypes_test

// import "testing"

// // func Test(t *testing.T) {
// // 	expectedOutput := []interface{}{1, "abc"}

// // 	actualOutput := moretypes.MoreTypes()
// // 	fmt.Println(expectedOutput, actualOutput)
// // }

// func Test(t *testing.T) {
// 	hello()
// }

package moretypes

import (
	"testing"
)

func TestHello(t *testing.T) {
	// 只运行被测试的函数，不进行断言和验证输出
	hello()
}
