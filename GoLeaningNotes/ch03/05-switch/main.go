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

	// 无须显式执行break语句，case执行完毕后自动中断。如须贯通后续case（源码顺序），
	//须执行fallthrough，但不再匹配后续条件表达式

	switch z1 := 5; z1 {
	default:
		println(z1)
	case 5:
		z1 += 10
		println("z1: ", z1)
		if z1 >= 15 {
			break
		}

		fallthrough
	case 6:
		z1 += 20
		println("z1: ", z1)
	}
}
