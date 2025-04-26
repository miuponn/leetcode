class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<>();            // initialize new hashmap

        for (int i = 0; i < nums.length; i++) {                     // iterate over elements in nums
            if (map.containsKey(target - nums[i])) {                // if map contains difference between target and element
                return new int[] {map.get(target - nums[i]), i};    // return indices of element and difference
            }
            map.put(nums[i], i);                                    // else update hashmap with element, index
        }
        return new int[] {};                                      
    }
}