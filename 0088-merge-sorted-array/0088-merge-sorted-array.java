class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {

        int last = m + n - 1;                      // pointer to last index of nums1 (merged array)

        while (m > 0 && n > 0) {                   // compare elements from the end of nums1 and nums2
            if (nums1[m - 1] > nums2[n - 1]){      // place the larger element at the current last position
                nums1[last] = nums1[m - 1];        
                m--;                           
            } else {                               
                nums1[last] = nums2[n - 1];        
                n--;                           
            }  
            last --;                               // move last pointer to the left
        }
        
        while (n > 0) {                            // place any remaining nums2 elements into nums1
            nums1[last] = nums2[n - 1];            
            n--;
            last--;
        }
    }
}