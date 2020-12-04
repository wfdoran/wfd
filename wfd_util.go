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

func ConvexMin(f func(float64) float64, left float64, right float64, eps float64) (float64, float64) {
	a := left
	b := (left + right) / 2.0
	c := right

	fa := f(a)
	fb := f(b)
	fc := f(c)

	for c - a > eps {
		if fb >= fa {
			c = b
			fc = fb
			b = (a + c) / 2.0
			fb = f(b)
			continue
		}

		if fb >= fc {
			a = b
			fa = fb
			b = (a + c) / 2.0
			fb = f(b)
			continue
		}

		if b - a > c - b {
			x := (a + b) / 2.0
			fx := f(x)

			if (fx > fb) {
				a = x
				fa = fx
			} else {
				c = b
				fc = fb
				b = x
				fb = fx
			}
		} else {
			x := (b + c) / 2.0
			fx := f(x)

			if (fx > fb) {
				c = x
				fc = fx
			} else {
				a = b
				fa = fb
				b = x
				fb = fx
			}
		}
	}

	return b, fb
}
