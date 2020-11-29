package wfd_util

import "fmt"

func Hello() {
     fmt.Println("Hello from wfd_util.")
}

func gen_partitions(n int) <-chan []int {
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
