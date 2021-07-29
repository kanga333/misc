from sys import stdin

class TopologicalSort:
    def __init__(self, adj):
        self.adj = adj
        self.visited = {}
        self.order = []
 
    def resolve(self):
        for s in range(len(self.adj)):
            if self.visited.get(s):
                continue
            self.dfs(s)
        self.order.reverse()
        return self.order

    def dfs(self, s):
        for d in self.adj[s]:
            if self.visited.get(d):
                continue
            self.visited[d] = True
            self.dfs(d)
        self.order.append(s)
        

def main():
    readline = stdin.readline
    v, e = map(int, readline().split())
    adj = [[] for _ in range(v)]
    for _ in range(v):
        s, e = map(int, readline().split())
        adj[s].append(e)
    result = TopologicalSort(adj).resolve()
    print(*result, sep="\n")

if __name__ == "__main__":
    main()
    
