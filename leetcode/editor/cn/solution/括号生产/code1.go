package 括号生产

// 自己写的破代码

type S struct {
	A byte
	L int
}

// 基于一颗虚拟的树，进行dfs
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	stack := make([]S, 0)
	r := make([]byte, 0, n*2)

	stack = append(stack, S{'(', 1})
	var p S
	var set bool

	for len(stack) > 0 {
		stack, p = popByte(stack)
		r = append(r, p.A)
		if len(r) >= n*2 {
			if !set {
				result = append(result, string(r))
				set = true
			} else {
				r = append(r[:p.L-1], r[len(r)-1])
				set = false
			}
		}

		cc := countC(r)
		cr := countReverseC(r)

		if cc < n {
			stack = append(stack, S{'(', p.L + 1})
		}

		if cr < cc {
			stack = append(stack, S{')', p.L + 1})
		}
	}
	return result
}

func countReverseC(r []byte) (count int) {
	for i := 0; i < len(r); i++ {
		if r[i] == ')' {
			count++
		}
	}
	return count
}

func countC(r []byte) (count int) {
	for i := 0; i < len(r); i++ {
		if r[i] == '(' {
			count++
		}
	}
	return count
}

func popByte(stack []S) ([]S, S) {
	n := len(stack)
	p := stack[n-1]
	stack = stack[:n-1]
	return stack, p
}
