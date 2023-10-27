package main

import (
	"errors"
	"fmt"
)

func main() {

	// 一个原始错误
	underlying := errors.New("underlying error")

	// 使用 %w 报告原始错误
	wrapped := fmt.Errorf("%w", underlying)

	// wrapped 会包含原始错误的信息
	fmt.Println("wrapped:", wrapped) // underlying error

	// 检查 wrapped 是否包含原始错误
	if errors.Is(wrapped, underlying) {
		fmt.Println("%w 报告原始错误")

	}

	// 使用 %v 简单打印错误
	other := fmt.Errorf("%v", underlying)

	// other 不包含原始错误的信息
	fmt.Println("other:", other) // errors.New("underlying error")

	// other 不包含原始错误
	fmt.Println("is other,", other)
	fmt.Println("is underlying,", underlying)
	if errors.Is(other, underlying) {
		fmt.Println("%v 也报告原始错误")
	} else {
		fmt.Println("%v 不报告原始错误")
	}

}
