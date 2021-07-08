def fib(n):
    dp = [1, 1]
    for i in range(2, n + 1):
        dp.append(dp[i-1] + dp[i-2])
    return dp[n]

# ALDS1_10_A: フィボナッチ数列
def main():
    n = int(input())
    print(fib(n))

if __name__ =='__main__':
    main()
