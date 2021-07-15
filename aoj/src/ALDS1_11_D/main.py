class Graph:
    def __init__(self, n):
        self.n = n
        self.linked_list = [[] for _ in range(n)]
    
    def link(self, x, y):
        self.linked_list[x].append(y)
        self.linked_list[y].append(x)


class DFS:
    def __init__(self, g):
        self.g = g
        self.stuck = []
        self.visited = [ -1 for _ in range(self.g.n)]
    
    def do(self):
        for i in range(self.g.n):
            if self.visited[i] != -1:
                continue
            self.visited[i] = i
            self.visit(i, i)
    
    def visit(self, id, root):
        list = self.g.linked_list[id]
        for i in list:
            if self.visited[i] != -1:
                continue
            self.visited[i] = root
            self.visit(i, root)
    
    def connected(self, x, y):
        return self.visited[x] == self.visited[y]

# ALDS1_11_D: 連結成分
def main():
    nm = [int(v) for v in input().split(" ")]
    n, m = nm
    g = Graph(n)
    for _ in range(m):
        xy = [ int(v) for v in input().split(" ")]
        g.link(*xy)
    d = DFS(g)
    d.do()
    an = int(input())
    for _ in range(an):
        xy = [ int(v) for v in input().split(" ")]
        if d.connected(*xy):
            print("yes")
        else:
            print("no")

if __name__ == '__main__':
    main()
