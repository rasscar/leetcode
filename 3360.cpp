class Solution {
public:
    bool canAliceWin(int n) {
        int b = 10;
        int a = 0;

        while (n >= b) {
            n -= b;
            --b;
            ++a;
        }

        return a % 2 == 1;
    }
};
