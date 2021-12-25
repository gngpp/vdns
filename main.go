package main

import (
	"fmt"
)

func Demo(i int) {
	//定义10个元素的数组
	var arr [10]int
	//错误拦截要在产⽣错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产⽣panic异常 打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10
}

func main() {
	Demo(10)
	//产⽣错误后 程序继续
	fmt.Println("程序继续执⾏...")
}
