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


class BFS:
    def __init__(self, g):
        self.g = g
        self.stuck = []
        self.visited = [ False for _ in range(self.g.n + 1)]
        self.d = {}
    
    def do(self):
        self.visit(1, 0)
        while(self.stuck != []):
            self.next_visit()
            
        for i in range(1, self.g.n + 1):
            if not self.visited[i]:
                self.d[i] = -1
            print(i, self.d[i])
    
    def next_visit(self):
        id, d = self.stuck.pop(0)
        list = self.g.get_linked_list(id)
        for i in list:
            self.visit(i, d)
    
    def visit(self, id, d):
        if self.visited[id]:
            return
        self.visited[id] = True
        self.d[id] = d
        self.stuck.append((id, d+1))

# ALDS1_11_C: 幅優先探索
def main():
    n = int(input())
    g = Graph(n)
    for _ in range(n):
        adj = [ int(v) for v in input().split(" ")]
        x = adj[0]
        for y in adj[2:]:
            g.enable(x, y)
    d = BFS(g)
    d.do()

if __name__ == '__main__':
    main()
