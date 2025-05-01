class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int l = prices[0];
        int res = 0;

        for (int i = 1; i < prices.size(); i++){
            if (prices[i] < l){
                l = prices[i];
            }
            int profit = prices[i] - l;
            if (profit > res){
                res = profit;
            }
        }
        return res;
    }
};