from sys import stdin

class Node:
    def __init__(self, loc, left, right) -> None:
        self.loc = loc
        self.left = left
        self.right = right
    
    def find(self, sx, ex, sy, ey, depth):
        if self.loc == None:
            return []
        x = self.loc[1]
        y = self.loc[2]

        ans = []
        if sx <= x and x <= ex and sy <= y and y <= ey:
            ans.append(self.loc[0])
        
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
        
        mid = int((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x[1])
        return Node(self.points[mid], self.makeY(l, mid), self.makeY(mid+1, r))

    def makeY(self, l, r):
        if l >= r:
            return Node(None, None, None)
        
        mid = int((l + r) / 2)
        self.points[l:r] = sorted(self.points[l:r],key=lambda x:x[2])
        return Node(self.points[mid], self.makeX(l, mid), self.makeX(mid+1, r))

# DSL_2_C: 領域探索
def main():
    readline = stdin.readline
    n = int(input())
    points = []
    for i in range(n):
        x, y = map(int, readline().split())
        points.append((i, x, y))
    tree = KDTree(points)
    node = tree.makeX(0, n)
    q = int(input())
    for _ in range(q):
        sx, ex, sy, ey = map(int, readline().split())
        ans = node.find(sx, ex, sy, ey, 0)
        ans.sort()
        if len(ans) != 0:
            print(*ans, sep='\n')
        print()
    
if __name__ == "__main__":
    main()
