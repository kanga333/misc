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
    while(len(t) != n):
        shortest = MAX_WEIGHT
        next = 0
        # TODO: 2分木で高速化できる
        for d in t:
            for i, v in enumerate(adj[d]):
                if i in t:
                    continue
                if v < shortest:
                    shortest = v
                    next = i
        total += shortest
        t.add(next)
    print(total)

if __name__ == '__main__':
    main()
