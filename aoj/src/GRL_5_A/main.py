from sys import stdin

class resolver:
    def __init__(self, adj):
        self.adj = adj

    def farthest(self, i, d, visited):
        visited.add(i)
        t = i
        w = 0
        for t2, w2 in self.adj[i]:
            if not t2 in visited:
                t3, w3 = self.farthest(t2, d+w2, visited)
                if w < w3:
                    w = w3
                    t = t3
        if w == 0:
            return i, d
        else:
            return t, w

    def resolve(self):
        s, _ = self.farthest(0, 0, set())
        _, w = self.farthest(s, 0, set())
        print(w)

# GRL_5_A: 木の直径
def main():
    readline = stdin.readline
    n = int(input())
    adj = [[] for _ in range(n)]
    for _ in range(n-1):
        s, t, w = map(int, readline().split())
        adj[s].append((t, w))
        adj[t].append((s, w))
    resolver(adj).resolve()
    
if __name__ == "__main__":
    main()
