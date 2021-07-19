from heapq import heappush, heappop
MAX_WEIGHT = 2001

# ALDS1_12_A: 最小全域木
def main():
    n = int(input())
    adj = []
    for i in range(n):
        edge = [MAX_WEIGHT if int(v) == -1 else int(v) for v in input().split(" ")[1:]]
        adj.append(edge)
    t = set({0})
    total = 0
    next = 0
    hq = []
    while(len(t) != n):
        for i, v in enumerate(adj[next]):
            if i in t:
                continue
            heappush(hq, (v, i))
        shortest = MAX_WEIGHT
        for _ in range(len(hq)):
            v, i = heappop(hq)
            if not i in t:
                shortest = v
                next = i
                break
        total += shortest
        t.add(next)
    print(total)

if __name__ == '__main__':
    main()
