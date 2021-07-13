class Graph:
    def __init__(self, n):
        self.n = n
        self.matrix = [[0 for _ in range(n)] for _ in range(n)]
        self.linked_list = [[] for _ in range(n)]
    
    def __str__(self):
        lines = []
        for line in self.matrix:
            lines.append(" ".join([str(v) for v in line]))
        return "\n".join(lines)
    
    def enable(self, x, y):
        self.matrix[x-1][y-1] = 1
    
    def get_linked_list(self, id):
        if self.linked_list[id - 1] != []:
            return self.linked_list[id - 1]
        val = [i + 1 for i, v in enumerate(self.matrix[id - 1]) if v == 1]
        self.linked_list[id - 1] = val
        return val


class DFS:
    def __init__(self, g):
        self.g = g
        self.stuck = []
        self.visited = [ False for _ in range(self.g.n + 1)]
        self.t = 0
        self.start = {}
        self.end = {}
    
    def do(self):
        for i in range(1, self.g.n + 1):
            if self.visited[i]:
                continue
            self.t += 1
            self.visited[i] = True
            self.start[i] = self.t
            self.visit(i)
        for i in range(1, self.g.n + 1):
            print(i, self.start[i], self.end[i]) 
    
    def visit(self, id):
        list = self.g.get_linked_list(id)
        for v in list:
            if self.visited[v]:
                continue
            self.t += 1
            self.visited[v] = True
            self.start[v] = self.t
            self.visit(v)
        self.t += 1
        self.end[id] = self.t

# ALDS1_11_B: 深さ優先探索
def main():
    n = int(input())
    g = Graph(n)
    for _ in range(n):
        adj = [ int(v) for v in input().split(" ")]
        x = adj[0]
        for y in adj[2:]:
            g.enable(x, y)
    d = DFS(g)
    d.do()

if __name__ == '__main__':
    main()
