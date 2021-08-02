import sys

class resolver:
    def __init__(self, adj):
        self.adj = adj

    def farthest(self, i):
        next = [(i, 0)]
        visited = set([i])
        at = i
        aw = 0
        while(len(next) != 0):
            i, tw = next.pop(0)
            for t, pw in self.adj[i]:
                if not t in visited:
                    nw = tw +pw
                    visited.add(t)
                    next.append((t, nw))
                    if aw < nw:
                        aw = nw
                        at = t
        return at, aw

    def resolve(self):
        s, _ = self.farthest(0)
        _, w = self.farthest(s)
        print(w)

# GRL_5_A: 木の直径
def main():
    readline = sys.stdin.readline
    n = int(input())
    adj = [[] for _ in range(n)]
    for _ in range(n-1):
        s, t, w = map(int, readline().split())
        adj[s].append((t, w))
        adj[t].append((s, w))
    resolver(adj).resolve()
    
if __name__ == "__main__":
    main()
