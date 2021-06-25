package _52_打开转盘锁

import "math"

/*
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为'0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func openLock(deadends []string, target string) int {
	deadMap := make(map[string]struct{})
	for _, dead := range deadends {
		deadMap[dead] = struct{}{}
	}
	return afs(deadMap, target, []byte{'0', '0', '0', '0'}, 0)
}

func afs(deadends map[string]struct{}, target string, lock []byte, time int) int {
	now := string(lock)
	if target == now {
		return time
	}
	nextTime := time + 1
	minTime := math.MaxInt64
	for i := 0; i < 4; i++ {
		if lock[i] < '9' {
			lock[i]++
			minTime = min(minTime, afs(deadends, target, lock, nextTime))
			lock[i]--
		}
	}
	if minTime == math.MaxInt64 {
		return -1
	}
	return minTime
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
