class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int l = prices[0];                          // initialize to the first day's price (minimum so far)
        int res = 0;                                // tracks the maximum profit found so far

        for (int i = 1; i < prices.size(); i++){    
            if (prices[i] < l){                     // update the minimum price if a lower one is found
                l = prices[i];
            }
            int profit = prices[i] - l;             // calculate profit by selling at current price after buying at min so far
            if (profit > res){                      // update max profit if this profit is higher
                res = profit;
            }
        }
        return res;
    }
};