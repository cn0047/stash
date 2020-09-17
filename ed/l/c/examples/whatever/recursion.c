#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int find_nth_term(int n, int a, int b, int c)
{
    int s = a + b + c;
    n--;
    if (n >= 0) {
        return find_nth_term(n, b, c, s);
    } else {
        return c;
    }
}

void term() {
    int n = 5, a = 1, b = 2, c = 3;
    n -= 3;
    int ans = find_nth_term(n, a, b, c);

    printf("%d", ans);
}

int main()
{
    term();
    return 0;
}
