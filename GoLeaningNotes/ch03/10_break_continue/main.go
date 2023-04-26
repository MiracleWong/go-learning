package main

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 5 {
			break
		}

		println(i)
	}

outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				println()
				continue outer
			}

			if x > 2 {
				break outer
			}
			print(x, ":", y, " ")
		}
	}
}
