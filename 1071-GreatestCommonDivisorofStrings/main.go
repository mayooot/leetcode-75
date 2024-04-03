package main

import (
	"fmt"
)

/*
对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。

给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。



示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""


提示：

1 <= str1.length, str2.length <= 1000
str1 和 str2 由大写英文字母组成
*/

// 闹麻了，纯数学题
// 如果 str1 和 str2 都能被 x 除尽，也就是都完全由 x 组成，也就是 str1 由 n 个 x 组成，str2 由 m 个 x 组成。
// 所以 str1 + str2 == str2 + str1
//

// 然后就用 gcd 算法找到它们最大公约数
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcd(len(str1), len(str2))]
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))

	fmt.Println(gcd(8, 4))
	fmt.Println(gcd(4, 8))
}
