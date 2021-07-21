from heapq import heappush, heappop

MAX = 2147483647

# ALDS1_12_C: 単一始点最短経路の高速化
def main():
    n = int(input())
    adj = []
    for _ in range(n):
        line = [int(v) for v in input().split(" ")]
        num = line[1]
        adj.append([ (line[2*j], line[2*j+1]) for j in range(1, num+1)])
    d = [MAX for _ in range(n)]
    d[0] = 0
    s = set()
    hq = []
    heappush(hq, (0, 0))
    while (len(s) != n):
        u = 0
        for _ in range(len(hq)):
            _, i = heappop(hq)
            if not i in s:
                u = i
                break
        s.add(u)
        for i, v in adj[u]:
            if d[i] > d[u] + v:
                d[i] = d[u] + v
                heappush(hq, (d[i], i))
    for i, v in enumerate(d):
        print(i, v)

if __name__ == '__main__':
    main()
