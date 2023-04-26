package main

func main() {
	data := [3]string{"a", "b", "c"}
	for i, s := range data {
		println(i, s)
	}

	for i := range data {
		println(i, data[i])
	}

	for s := range data {
		println(s)
	}

	for range data {
	}

	for i, s := range data {

		// 局部变量会重复使用
		println(&i, &s)
	}
}
