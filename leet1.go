/*
1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

这道题最优的做法时间复杂度是 O(n)
顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。

*/

package leetcode

func twoSum(nums []int, target int) []int {
	a:=len(nums)
	mapNumber := make(map[int]int)
	for i:=0;i<a;i++{
		for k,v:= range mapNumber {
			if k==target -nums[i]{
				return []int{i,v}
			}
		}
		mapNumber[nums[i]] =i
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	mapNumber := make(map[int]int)
	for i:=0;i<len(nums);i++{
		another:=target-nums[i]
		if _,ok:=mapNumber[another];ok{
			return []int{i,mapNumber[another]}
		}
		mapNumber[nums[i]] =i
	}
	return nil
}
