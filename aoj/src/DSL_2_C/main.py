from sys import stdin
class KDTree:
    def __init__(self, points, nodes):
        self.points = points
        self.nodes = nodes
        self.np = 0
        self.ans = []
    
    def makeX(self, l, r):
        if l >= r:
            return -1
        
        mid = int((l + r) / 2)
        t = self.np
        self.np += 1
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x[1])
        self.nodes[t][0] = mid
        self.nodes[t][1] = self.makeY(l, mid)
        self.nodes[t][2] = self.makeY(mid+1, r)
        return t

    def makeY(self, l, r):
        if l >= r:
            return -1
        
        mid = int((l + r) / 2)
        t = self.np
        self.np += 1
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x[2])
        self.nodes[t][0] = mid
        self.nodes[t][1] = self.makeX(l, mid)
        self.nodes[t][2] = self.makeX(mid+1, r)
        return t
    
    def find(self, t, sx, ex, sy, ey, depth):
        if t == -1:
            return
        x = self.points[self.nodes[t][0]][1]
        y = self.points[self.nodes[t][0]][2]

        if sx <= x and x <= ex and sy <= y and y <= ey:
            self.ans.append(self.points[self.nodes[t][0]][0])
        
        if depth%2 == 0:
            if sx <= x:
                self.find(self.nodes[t][1],sx, ex, sy, ey, depth+1)
            if x <= ex:
                self.find(self.nodes[t][2],sx, ex, sy, ey, depth+1)
        else:
            if sy <= y:
                self.find(self.nodes[t][1],sx, ex, sy, ey, depth+1)
            if y <= ey:
                self.find(self.nodes[t][2],sx, ex, sy, ey, depth+1)
        return

# DSL_2_C: 領域探索
def main():
    readline = stdin.readline
    n = int(input())
    points = []
    nodes = [[0] * 3 for _ in range(n)]
    for i in range(n):
        x, y = map(int, readline().split())
        points.append((i, x, y))
    tree = KDTree(points, nodes)
    _ = tree.makeX(0, n)
    q = int(input())
    for _ in range(q):
        sx, ex, sy, ey = map(int, readline().split())
        tree.find(0, sx, ex, sy, ey, 0)
        tree.ans.sort()
        if len(tree.ans) != 0:
            print(*tree.ans, sep='\n')
        print()
        tree.ans = []

if __name__ == "__main__":
    main()
