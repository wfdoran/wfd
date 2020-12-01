package wfd_util

import "fmt"

func Hello() {
     fmt.Println("Hello from wfd_util.")
}

func GenPartitions(n int) <-chan []int {
	ch := make(chan []int)

	go func(ch chan []int) {
		state := []int{n}
	
		for {
			x := make([]int, len(state))
			copy(x, state)
			ch <- x
		
			if state[0] == 1 {
				close(ch)
				return
			}
		
			idx := len(state) - 1
			for state[idx] == 1 {
				idx--
			}
		
			need := state[idx] + len(state) - idx - 1
			val := state[idx] - 1
		
			state = state[:idx]
		
			for need > 0 {
				if need >= val {
					state = append(state, val)
					need -= val
				} else {
					state = append(state, need)
					need -= need
				}
			}
		}
	}(ch)
	return ch
}

func ExtendedGcd(a, b int64) (int64, int64, int64) {
	if a < b {
		x, y, g := ExtendedGcd(b, a)
		return y, x, g
	}

	if b == 0 {
		return 1, 0, a
	}

	t := a / b
	r := a % b

	x, y, g := ExtendedGcd(b, r)

	return y, x - t *y, g
}

func ModInv(a, m int64) int64 {
	_, x, g := ExtendedGcd(m, a)
	if g != 1 {
		return 0
	}
	return (x + m) % m
}

