# メイン関数
# 比較関数
# バイナリサーチ
from sys import stdin
import math

MAX_WEIGHT = 10_000

def carriable(k, max, wi):
    shipped = 0
    picked = 0
    for w in wi:
        if picked + w > max:
            picked = 0
            shipped +=1
        picked += w
        if shipped > k or w > max:
            return False
    shipped += 1
    return shipped <= k


def binarySearch(k, max, wi):
    right = max
    left = 0
    while left < right:
        mid = int((left + right) /2 )
        if carriable(k, mid, wi):
            right = mid
        else:
            left = mid + 1
    return right

def main():
    args = stdin.readline().strip().split(" ")
    n = int(args[0])
    k = int(args[1])
    wi = [int(input()) for _ in range(0, n)]
    max = math.ceil(n * MAX_WEIGHT / k)
    result = binarySearch(k, max, wi)
    print(result)

if __name__ == '__main__':
    main()
