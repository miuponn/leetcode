class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}                                   # initialize dict

        for i, num in enumerate(nums):              # iterate over index and element in array nums
            diff = target - num
            if diff in seen:                        # if difference exists in dict
                return [seen[diff], i]              # return indices of difference, element
            seen[num] = i                           # else add element, index to dict
