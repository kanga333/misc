MAX = 2147483647

# ALDS1_12_B: 単一始点最短経路
def main():
    n = int(input())
    adj = [[ MAX for _ in range(n) ] for _ in range(n)]
    for i in range(n):
        line = [int(v) for v in input().split(" ")]
        num = line[1]
        for j in range(1, num+1):
            adj[i][line[2*j]] = line[2*j+1]
    d = [MAX for _ in range(n)]
    d[0] = 0
    s = set({0})
    while (len(s) != n ):
        min = MAX
        u = 0
        for i, v in enumerate(d):
            if min > v and not i in s:
                min = v
                u = i
        s.add(u)
        print(u)
        for i, v in enumerate(adj[u]):
            if d[i] > d[u] + v:
                d[i] = d[u] + v
    for i, v in enumerate(d):
        print(i, v)



if __name__ == "__main__":
    main()
