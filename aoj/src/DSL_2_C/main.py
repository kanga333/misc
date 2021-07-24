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
    
    def find(self, sx, ex, sy, ey, depth):
        if self.loc == None:
            return []
        x = self.loc.x
        y = self.loc.y

        ans = []
        if sx <= x and x <= ex and sy <= y and y <= ey:
            ans.append(self.loc.id)
        
        if depth%2 == 0:
            if sx <= x:
                ans.extend(self.left.find(sx, ex, sy, ey, depth+1))
            if x <= ex:
                ans.extend(self.right.find(sx, ex, sy, ey, depth+1))
        else:
            if sy <= y:
                ans.extend(self.left.find(sx, ex, sy, ey, depth+1))
            if y <= ey:
                ans.extend(self.right.find(sx, ex, sy, ey, depth+1))
        return ans
class KDTree:
    def __init__(self, points) -> None:
        self.points = points
    
    def makeX(self, l, r):
        if l >= r:
            return Node(None, None, None)
        
        mid = math.floor((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x.x)
        return Node(self.points[mid], self.makeY(l, mid), self.makeY(mid+1, r))

    def makeY(self, l, r):
        if l >= r:
            return Node(None, None, None)
        
        mid = math.floor((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x.y)
        return Node(self.points[mid], self.makeX(l, mid), self.makeX(mid+1, r))

# DSL_2_C: 領域探索
def main():
    n = int(input())
    points = []
    for i in range(n):
        xy = [int(v) for v in input().split(" ")]
        points.append(Point(i, xy[0], xy[1]))
    tree = KDTree(points)
    node = tree.makeX(0, n)
    q = int(input())
    for _ in range(q):
        line = [int(v) for v in input().split(" ")]
        ans = node.find(line[0],line[1],line[2],line[3],0)
        ans.sort()
        for v in ans:
            print(v)
        print()
    
if __name__ == "__main__":
    main()
