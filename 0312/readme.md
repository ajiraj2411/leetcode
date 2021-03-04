You are given n balloons, indexed from 0 to n - 1. Each balloon is painted with a number on it represented by an array nums. You are asked to burst all the balloons.

If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins. If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.

Return the maximum coins you can collect by bursting the balloons wisely.

 

Example 1:

Input: nums = [3,1,5,8]
Output: 167
Explanation:
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
Example 2:

Input: nums = [1,5]
Output: 10
 

Constraints:

n == nums.length
1 <= n <= 500
0 <= nums[i] <= 100


------------------
mysolution
------------------

func maxCoins(arr []int) int {
	coins := 0
	for len(arr) > 3 {
		idx := smallestNumberPostion(arr)
		coins += (arr[idx] * arr[idx+1] * arr[idx+2])
		arr = append(arr[:idx+1], arr[idx+2:]...)
	}
	if len(arr) == 3 {
		coins += (arr[0] * arr[1] * arr[2])
		arr = append(arr[:1], arr[2])
	}
	if len(arr) == 2 {
		if arr[0] < arr[1] {
			coins += ((arr[0] * arr[1]) + arr[1])
			return coins
		}
		coins += ((arr[0] * arr[1]) + arr[0])
		return coins
	}
	coins += arr[0]
	return coins
}

func smallestNumberPostion(array []int) int {
	needArray := array[1 : len(array)-1]
	smallestNumber := needArray[0]
	position := 0
	for i, element := range needArray {
		if element < smallestNumber {
			smallestNumber = element
			position = i
		}
	}
	return position
}



----------------
video - working 
----------------

func maxCoins(arr []int) int {
	arr = append([]int{1}, arr...)
	arr = append(arr, 1)
	return getCoins(arr, 0, len(arr)-1)
}

var mFor = make(map[string]int)

func getCoins(nums []int, l int, r int) int {
	key := strconv.Itoa(l) + "-" + strconv.Itoa(r)
	if _, ok := mFor[key]; !ok {
		coins := 0
		for i := l + 1; i < r; i++ {
			coin := nums[l] * nums[i] * nums[r]
			data := coin + getCoins(nums, l, i) + getCoins(nums, i, r)
			if data > coins {
				coins = data
			}
			mFor[key] = coins
		}
	}
	return mFor[key]
}
