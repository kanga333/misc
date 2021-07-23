import math

class Point:
    def __init__(self, id, x, y) -> None:
        self.id = id
        self.x = x
        self.y = y

class Node:
    def __init__(self, loc, left, right) -> None:
        self.loc = loc
        self.left = left
        self.right = right

class KDTree:
    def __init__(self, points) -> None:
        self.points = points
        self.nodes = [Node(0, 0, 0) for _ in points]
        self.np = 0
        self.ans = []
    
    def makeX(self, l, r):
        if l >= r:
            return -1
        
        t = self.np
        self.np += 1
        mid = math.floor((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x.x)
        self.nodes[t].loc = mid
        self.nodes[t].left = self.makeY(l, mid)
        self.nodes[t].right = self.makeY(mid+1, r)

        return t

    def makeY(self, l, r):
        if l >= r:
            return -1
        
        t = self.np
        self.np += 1
        mid = math.floor((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x.y)
        self.nodes[t].loc = mid
        self.nodes[t].left = self.makeX(l, mid)
        self.nodes[t].right = self.makeX(mid+1, r)

        return t
    
    def find(self, t, sx, ex, sy, ey, depth):
        if t == -1:
            return
        x = self.points[self.nodes[t].loc].x
        y = self.points[self.nodes[t].loc].y

        if sx <= x and x <= ex and sy <= y and y <= ey:
            self.ans.append(self.points[self.nodes[t].loc].id)
        
        if depth%2 == 0:
            if sx <= x:
                self.find(self.nodes[t].left, sx, ex, sy, ey, depth+1)
            if x <= ex:
                self.find(self.nodes[t].right, sx, ex, sy, ey, depth+1)
        else:
            if sy <= y:
                self.find(self.nodes[t].left, sx, ex, sy, ey, depth+1)
            if y <= ey:
                self.find(self.nodes[t].right, sx, ex, sy, ey, depth+1)

# DSL_2_C: 領域探索
def main():
    n = int(input())
    points = []
    for i in range(n):
        xy = [int(v) for v in input().split(" ")]
        points.append(Point(i, xy[0], xy[1]))
    tree = KDTree(points)
    tree.makeX(0, n)
    q = int(input())
    for _ in range(q):
        line = [int(v) for v in input().split(" ")]
        tree.find(0, line[0],line[1],line[2],line[3],0)
        ans = tree.ans
        tree.ans = []
        ans.sort()
        for v in ans:
            print(v)
        print()
    
if __name__ == "__main__":
    main()
