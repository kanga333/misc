class Graph:
    def __init__(self, n):
        self.n = n
        self.matrix = [[0 for _ in range(n)] for _ in range(n)]
    
    def __str__(self):
        lines = []
        for line in self.matrix:
            lines.append(" ".join([str(v) for v in line]))
        return "\n".join(lines)
    
    def enable(self, x, y):
        self.matrix[x-1][y-1] = 1

# ALDS1_11_A: グラフ
def main():
    n = int(input())
    g = Graph(n)
    for _ in range(n):
        adj = [ int(v) for v in input().split(" ")]
        x = adj[0]
        for y in adj[2:]:
            g.enable(x, y)
    print(g)

if __name__ == '__main__':
    main()
