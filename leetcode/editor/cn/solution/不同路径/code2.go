package 不同路径

type Node struct {
	i, j int
	M, N int
}

func (n *Node) Right() *Node {
	ne := &Node{
		i: n.i,
		j: n.j + 1,
	}
	if ne.j >= n.M {
		return nil
	}

	ne.M = n.M
	ne.N = n.N
	return ne
}

func (n *Node) Down() *Node {
	ne := &Node{
		i: n.i + 1,
		j: n.j,
	}
	if ne.i >= n.N {
		return nil
	}

	ne.M = n.M
	ne.N = n.N
	return ne
}

func (n *Node) IsEnd() bool {
	if n.i == n.N-1 && n.j == n.M-1 {
		return true
	}
	return false
}

// DFS 深度优先搜索来实现
func uniquePathsDFS(m int, n int) int {
	root := &Node{
		i: 0, j: 0, M: m, N: n,
	}
	return dfs(root)
}

func dfs(p *Node) int {
	count := 0
	stack := make([]*Node, 0)
	stack = append(stack, p)
	visited := make(map[*Node]int)

	for len(stack) > 0 {
		stack, p = pop(stack)
		if p == nil {
			continue
		}
		_, ok := visited[p]
		if ok {
			continue
		}

		if p.IsEnd() {
			count++
		} else {
			visited[p] = 1
		}

		r := p.Right()
		_, ok = visited[r]
		if r != nil && !ok {
			stack = append(stack, r)
		}
		d := p.Down()
		_, ok = visited[d]
		if d != nil && !ok {
			stack = append(stack, d)
		}
	}
	return count
}

func pop(stack []*Node) ([]*Node, *Node) {
	n := len(stack)
	if n == 0 {
		return stack, nil
	}
	v := stack[n-1]
	stack = stack[:n-1]
	return stack, v
}
