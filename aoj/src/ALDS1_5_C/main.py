import math

th = math.pi * 60.0 / 180.0

class point:
    def __init__(self, x, y) -> None:
        self.x = x
        self.y = y

def calcTriangle(p1, p2):
    return point(
        (p2.x-p1.x)*math.cos(th) - (p2.y-p1.y)*math.sin(th) + p1.x,
        (p2.x-p1.x)*math.sin(th) + (p2.y-p1.y)*math.cos(th) + p1.y,
    )

def divVector(ratio, p1, p2):
    return point(
		(1.0-ratio)*p1.x + ratio*p2.x,
		(1.0-ratio)*p1.y + ratio*p2.y,
	)

def koch(p1, p2, n):
    if n == 0:
        return
    
    s = divVector(1.0/3.0, p1, p2)
    t = divVector(2.0/3.0, p1, p2)
    u = calcTriangle(s, t)

    koch(p1, s, n-1)
    print('{:.8f}'.format(s.x), '{:.8f}'.format(s.y))
    koch(s, u, n-1)
    print('{:.8f}'.format(u.x), '{:.8f}'.format(u.y))
    koch(u, t, n-1)
    print('{:.8f}'.format(t.x), '{:.8f}'.format(t.y))
    koch(t, p2, n-1)
    return

# ALDS1_5_C: コッホ曲線
def main():
    n = int(input())
    p1 = point(0.0, 0.0)
    p2 = point(100.0, 0.0)
    print('{:.8f}'.format(p1.x), '{:.8f}'.format(p1.y))
    koch(p1, p2, n)
    print('{:.8f}'.format(p2.x), '{:.8f}'.format(p2.y))

if __name__ == "__main__":
    main()
