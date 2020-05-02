#include <stdio.h>
int sum(int *nums, int i, int j)
{
    int sum = 0;
    for (; i < j; j++)
    {
        sum += nums[i];
    }
    return sum;
}

int main()
{
    // scanf("%d", &n);
    // int nums[n];
    // for (i = 0; i < n; i++)
    // {
    //     scanf("%d", &nums[i]);
    // }
    // printf("----- %d,%d",i,n);
    //  for (i = 0; i < n; i++)
    // {
    //     printf("%d", nums[i]);
    // }
    int n = 5;
    int i;
    int nums[5] = {3,-2,1,0,-2};
    int max = 0;
    // int sum[n];
    printf("----------");
    for (i = 0; i < n - 1; i++)
    {
        printf("----2------");

        for (int j = i + 1; i < n; j++)
        {
            int s = sum(nums, i, j);
            printf("%d", s);
            if (s == 0)
            {
                int l = j - i;
                if (l > max)
                {
                    max = l;
                }
            }
        }
    }
    printf("%d\n", max);
}