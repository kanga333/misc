import sys

class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def norm(self):
        return self.x * self.x + self.y * self.y
    
    def asu(self, r):
        self.x *= r
        self.y *= r
    
    def minus(self, p):
        return Vector(self.x - p.x, self.y - p.y)
    
    def plus(self, p):
        return Vector(self.x + p.x, self.y + p.y)

def dot(a, b):
    return a.x * b.x + a.y * b.y

def cross(a, b):
    return a.x * b.y - a.y * b.x

def project(p1, p2, p):
    base = p2.minus(p1)
    r = 1.0 * dot(p.minus(p1), base) / base.norm()
    base.asu(r)
    return base.plus(p1)

def main():
    readline = sys.stdin.readline
    x1, y1, x2, y2 = map(int, readline().split())
    p1 = Vector(x1, y1)
    p2 = Vector(x2, y2)

    n = int(input())
    for _ in range(n):
        x, y = map(int, readline().split())
        p = Vector(x, y)
        s = project(p1, p2, p)
        print('{:.10f}'.format(s.x), '{:.10f}'.format(s.y))

if __name__ == "__main__":
    main()
