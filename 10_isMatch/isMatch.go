package isMatch

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/

/*
动态规划法 本题的dp数组的含义就是：dp[i][j]就是s的前i个元素是否可以被p的前j个元素所匹配。
我们知道了dp数组的含义之后就知道了dp数组的几个细节：

1. dp[0][0]一定是true，因为s为空且p也为空的时候一定是匹配的；dp[1][0]一定是false，因为s有一个字符但是p为空的时候一定是不匹配的。
2. 这个boolean类型的dp数组的大小应该是dp[s.length+1][p.length+1]，因为我们不仅仅要分别取出s和p的所有元素，还要表示分别取s和p的0个元素时候(都为空)的情况。
3. 当写到dp[s.length][p.length]的时候，我们就得到了最终s和p的匹配情况。
4. dp[1][0]~dp[s.length][0]这一列都是false，因为s不为空但是p为空一定不能匹配。
*/

/*
如果 p.charAt(j) == s.charAt(i) : dp[i][j] = dp[i-1][j-1]；
如果 p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1]；
如果 p.charAt(j) == '*'：
如果 p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2] //in this case, a* only counts as empty
如果 p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.'：
dp[i][j] = dp[i-1][j] //in this case, a* counts as multiple a
or dp[i][j] = dp[i][j-1] // in this case, a* counts as single a
or dp[i][j] = dp[i][j-2] // in this case, a* counts as empty
*/
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	dp := make([][]Bool, sLen+1)
	for i := range dp {
		dp[i] = make([]Bool, pLen+1)
	}

	dp[0][0] = true
	for i := range s {
		for j := range p {
			if p[j] == s[i] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
			}
		}
	}

	return bool(dp[sLen][pLen])
}

type Bool bool

func (b Bool) String() string {
	if b {
		return "√\t"
	}
	return "×\t"
}
