package com.mayooot.lc75._1071;

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

public class Solution {
    public String gcdOfStrings(String str1, String str2) {
        if (!(str1 + str2).equals(str2 + str1)) {
            return "";
        }
        int max = gcd(str1.length(), str2.length());
        return str1.substring(0, max);
    }

    public int gcd(int x, int y) {
        if (y == 0) {
            return x;
        }

        return gcd(y, x%y);
    }
}