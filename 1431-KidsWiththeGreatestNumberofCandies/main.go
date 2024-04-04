package main

import "fmt"

/*
给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。



示例 1：

输入：candies = [2,3,5,1,3], extraCandies = 3
输出：[true,true,true,false,true]
解释：
孩子 1 有 2 个糖果，如果他得到所有额外的糖果（3个），那么他总共有 5 个糖果，他将成为拥有最多糖果的孩子。
孩子 2 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
孩子 3 有 5 个糖果，他已经是拥有最多糖果的孩子。
孩子 4 有 1 个糖果，即使他得到所有额外的糖果，他也只有 4 个糖果，无法成为拥有糖果最多的孩子。
孩子 5 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
示例 2：

输入：candies = [4,2,1,1,2], extraCandies = 1
输出：[true,false,false,false,false]
解释：只有 1 个额外糖果，所以不管额外糖果给谁，只有孩子 1 可以成为拥有糖果最多的孩子。
示例 3：

输入：candies = [12,1,12], extraCandies = 10
输出：[true,false,true]


提示：

2 <= candies.length <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, 0, len(candies))
	for i := 0; i < len(candies); i++ {
		tmp, isMax := candies[i]+extraCandies, true
		for j := 0; j < len(candies); j++ {
			if tmp < candies[j] {
				isMax = false
				break
			}
		}
		res = append(res, isMax)
	}

	return res
}

func kidsWithCandies2(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, v := range candies {
		maxCandies = max(maxCandies, v)
	}

	res := make([]bool, 0, len(candies))
	for _, v := range candies {
		if v+extraCandies >= maxCandies {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}

	return res
}

func main() {
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies2([]int{4, 2, 1, 1, 2}, 1))
}
