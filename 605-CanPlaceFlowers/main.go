package main

import "fmt"

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。



示例 1：

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true
示例 2：

输入：flowerbed = [1,0,0,0,1], n = 2
输出：false


提示：

1 <= flowerbed.length <= 2 * 104
flowerbed[i] 为 0 或 1
flowerbed 中不存在相邻的两朵花
0 <= n <= flowerbed.length
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	var count int
	if flowerbed[0] == 0 {
		if len(flowerbed) == 1 {
			// 第一个位置是 0，并且数组长度为 1
			count++
			flowerbed[0] = 1
		} else {
			// 第一个位置是0，并且数组长度不是 0，就要判断第二个位置是不是 0
			if flowerbed[1] == 0 {
				count++
				flowerbed[0] = 1
			}
		}
	}

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			count++
			flowerbed[i] = 1
		}
	}

	if flowerbed[len(flowerbed)-1] == 0 {
		// 这里不需要判断 len(flowerbed-2) 是否出现数组越界异常
		// 因为如果长度为 1，值为 0 的情况下，这个位置已经被插了，进不到这里
		// 如果长度为 1，值为 1 的情况下，更进不到这里
		if flowerbed[len(flowerbed)-2] == 0 {
			count++
			flowerbed[len(flowerbed)-1] = 1
		}
	}

	return count >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
