package main

func main() {
	a, b, c, x := 1, 2, 3, 2
	switch x {
	case a, b:
		println("a | b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println("z")
	}

	// switch 支持初始化语句，从上到下，从左到右匹配case执行
	// 全部失败的时候，执行default，建议放到最后
	switch x1 := 5; x1 {
	default:
		x1 += 100
		println(x1)
	case 5:
		x1 += 50
		println(x1)
	}

	y := 2
	switch y {
	case a:
	//单条件，内容为空，隐式：case a: break;
	case b:
		println("b")
	}

	// 不能重复出现case常量值
	switch z := 5; z {
	case 5:
		println("a")
	//case 6, 5:
	case 6, 4:
		println("b")
	}

}
