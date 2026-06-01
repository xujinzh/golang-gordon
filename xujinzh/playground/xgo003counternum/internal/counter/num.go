package counter

import "fmt"

func CounterNum() {
	var num int64

	fmt.Print("请输入一个整数:")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("输入错误，请输入合法的整数:", err)
		return
	}
	numInit := num
	count := 0

	if num == 0 {
		count++
	}

	if num < 0 {
		num = -num
	}

	for {
		if num == 0 {
			break
		}
		num /= 10
		count++
	}

	fmt.Printf("整数%v的位数是%v\n", numInit, count)

}
