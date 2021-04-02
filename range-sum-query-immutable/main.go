package main

type NumArray struct {
	sums []int
	nums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums), len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[i] = nums[i]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
	}
	return NumArray{sums: sums, nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right] - this.sums[left]
}
