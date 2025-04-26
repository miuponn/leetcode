class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # initialize dict
        seen = {}                                   

        # iterate over index and num in array nums
        for i, num in enumerate(nums):              
            diff = target - num
            # if complement of current num exists in dict, return indices of complement and current num
            if diff in seen:                    
                return [seen[diff], i]
            # append current num and its array index to dict          
            seen[num] = i                      