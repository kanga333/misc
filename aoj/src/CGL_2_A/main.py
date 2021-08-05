import sys

class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y

def dot(a, b):
    return a.x * b.x + a.y * b.y

def cross(a, b):
    return a.x * b.y - a.y * b.x

def main():
    readline = sys.stdin.readline

    n = int(input())
    for _ in range(n):
        x0, y0, x1, y1, x2, y2, x3, y3 = map(int, readline().split())
        a = Vector(x1-x0, y1-y0)
        b = Vector(x3-x2, y3-y2)
        if dot(a, b) == 0:
            print(1)
        elif cross(a, b) == 0:
            print(2)
        else:
            print(0)

if __name__ == "__main__":
    main()
