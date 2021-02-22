import sys

MAX = 10**9

def main():
    n = int(sys.stdin.readline())
    profit = MAX * -1
    bottom = int(sys.stdin.readline())

    for i in range(n-1):
        price = int(sys.stdin.readline())
        if profit < price - bottom:
            profit = price - bottom
        if bottom > price:
            bottom = price

    print(profit)

if __name__ == '__main__':
    main()
