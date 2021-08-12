import sys

MAX = 10_001

def resolve(n, c):
    t = [MAX] * (n+1)
    t[0] = 0
    for v in c:
        for j in range(v, (n+1)):
            if j-v < 0:
                continue
            if t[j] > t[j-v]+1:
                t[j] = t[j-v]+1
    return t[n]


def main():
    readline = sys.stdin.readline
    n, _ = map(int, readline().split())
    c = [int(n) for n in readline().split()]
    print(resolve(n,c))

if __name__ == "__main__":
    main()
