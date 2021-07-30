from sys import stdin

class TopologicalSort:
    def __init__(self, adj):
        self.adj = adj
        self.visited = {}
        self.order = []
 
    def resolve(self):
        for s in range(len(self.adj)):
            self.dfs(s)
        self.order.reverse()
        return self.order

    def dfs(self, s):
        if self.visited.get(s):
            return
        self.visited[s] = True
        for d in self.adj[s]:
            self.dfs(d)
        self.order.append(s)
        

def main():
    readline = stdin.readline
    v, e = map(int, readline().split())
    adj = [[] for _ in range(v)]
    for _ in range(e):
        s, d = map(int, readline().split())
        adj[s].append(d)
    result = TopologicalSort(adj).resolve()
    print(*result, sep="\n")

if __name__ == "__main__":
    main()
    
