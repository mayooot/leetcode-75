package main

import "fmt"

/*
345. 反转字符串中的元音字母
简单
相关标签
相关企业
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。



示例 1：

输入：s = "hello"
输出："holle"
示例 2：

输入：s = "leetcode"
输出："leotcede"


提示：

1 <= s.length <= 3 * 105
s 由 可打印的 ASCII 字符组成
*/

var sets = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}

type kv struct {
	key   int
	value byte
}

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}

	list := make([]kv, 0, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := sets[s[i]]; ok {
			list = append(list, kv{
				key:   i,
				value: s[i],
			})
		}
	}

	for i := 0; i < len(list)/2; i++ {
		list[i].value, list[len(list)-i-1].value = list[len(list)-i-1].value, list[i].value
	}

	ss := []byte(s)
	for _, kv := range list {
		ss[kv.key] = kv.value
	}
	return string(ss)
}

// 双指针
var sets2 = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}

func reverseVowels2(s string) string {
	if len(s) == 1 {
		return s
	}

	ss := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		_, ok := sets2[ss[l]]
		_, ok2 := sets2[ss[r]]
		if !ok && ok2 {
			l++
		} else if ok && !ok2 {
			r--
		} else if ok && ok2 {
			ss[l], ss[r] = ss[r], ss[l]
			l++
			r--
		} else if !ok && !ok2 {
			l++
			r--
		}
	}

	return string(ss)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
	fmt.Println("-------------------------")

	fmt.Println(reverseVowels2("hello"))
	fmt.Println(reverseVowels2("leetcode"))
	fmt.Println(reverseVowels2("aA"))
}
