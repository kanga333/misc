from sys import stdin

MAX = 9223372036854775807

def warsdhall_floyd(adj, v):
    for k in range(v):
        for i in range(v):
            if adj[i][k] == MAX:
                continue
            for j in range(v):
                if adj[k][j] == MAX:
                    continue
                r = adj[i][k] + adj[k][j]
                if r < adj[i][j]:
                    adj[i][j] = r

def negative(adj, v):
    for i in range(v):
        for j in range(v):
            if adj[i][i] < 0:
                return True
    return False

# GRL_1_C: 全点対間最短経路
def main():
    readline = stdin.readline
    v, e = map(int, readline().split())
    adj = [[MAX] * v for _ in range(v)]
    for i in range(v):
        adj[i][i] = 0
    for _ in range(e):
        i, j, w = map(int, readline().split())
        adj[i][j] = w
    warsdhall_floyd(adj, v)
    if negative(adj, v):
        print("NEGATIVE CYCLE")
    for i in range(v):
        ret = map(lambda x: 'INF' if x == MAX else str(x) ,adj[i])
        print(*ret)

if __name__ == "__main__":
    main()
