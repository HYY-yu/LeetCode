package 买卖股票的最佳时机

import "math"

func maxProfit(prices []int) int {
	// MP[i,j] // i代表天数，j代表当天是否持有股票 MP代表i天的最大利润
	// MP[i,0] = MAX( MP[i-1,0], MP[i-1,1]-a[i])
	// MP[i,1] = MAX( MP[i-1,1], MP[i-1,0]+a[i])
	// MP[0,0] = 0
	// MP[0,1] = a[0]
	if len(prices) == 1 || len(prices) == 0 {
		return 0
	}
	MP := make([][]int, 2)
	MP[0] = make([]int, 2)
	MP[0][0] = 0
	MP[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		x, y := (i-1)%2, i%2
		if MP[y] == nil {
			MP[y] = make([]int, 2)
		}
		MP[y][0] = int(math.Max(float64(MP[x][0]), float64(MP[x][1]+prices[i])))
		MP[y][1] = int(math.Max(float64(MP[x][1]), float64(-prices[i])))
	}
	return MP[(len(prices)-1)%2][0]
}
