#include <vector>

using namespace std;

class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int buy(0), sell(0), pBuy(INT_MIN), pSell(0), ppSell(0);
        for (int p : prices)
        {
            buy = max(ppSell - p, pBuy);
            sell = max(pBuy + p, pSell);
            ppSell = pSell;
            pSell = sell;
            pBuy = buy;
        }
        return sell;
    }
};