import sys

class Resolver:
    def __init__(self, adj):
        self.adj = adj
        self.t = 1
        self.visited = [False] * len(adj)
        self.parent = [0] * len(adj)
        self.pre = [0] * len(adj)
        self.lowest = [0] * len(adj)

    def resolve(self):
        self.dfs(0, 0)
        ret = []
        np = False
        for i in range(1, len(self.adj)):
            p = self.parent[i]
            if p == 0:
                np = True
            else:
                if self.pre[p] <= self.lowest[i]:
                    ret.append(p)
        if np:
            ret.append(0)
        ret = list(set(ret))
        ret.sort()
        return ret


    def dfs(self, i, previous_i):
        self.visited[i] = True
        self.pre[i] = self.t
        self.lowest[i] = self.t
        self.t += 1

        for v in self.adj[i]:
            if not self.visited[v]:
                self.parent[v] = i
                self.dfs(v, i)
                if self.lowest[i] > self.lowest[v]:
                    self.lowest[i] = self.lowest[v]
            elif v != previous_i:
                if self.lowest[i] > self.pre[v]:
                    self.lowest[i] = self.pre[v]

def main():
    readline = sys.stdin.readline
    v, e = map(int, readline().split())
    adj = [[] for _ in range(v)]
    for _ in range(e):
        s, d = map(int, readline().split())
        adj[s].append(d)
        adj[d].append(s)
    result = Resolver(adj).resolve()
    print(*result, sep="\n")

if __name__ == "__main__":
    main()
