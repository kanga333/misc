import math

class Heap:
    def __init__(self, a) -> None:
        self.a = a
        self.a.insert(0, None)

    def parent(self, i):
        pi = math.floor(i / 2)
        return self.a[pi]
    
    def left(self, i):
        li = 2 * i
        if len(self.a) <= li:
            return None
        return self.a[li]


    def right(self, i):
        ri = (2 * i) + 1
        if len(self.a) <= ri:
            return None
        return self.a[ri]

    def print(self, i):
        val = "node {}: key = {}, ".format(i, self.a[i])
        p = self.parent(i)
        if p:
            val += "parent key = {}, ".format(p)
        l = self.left(i)
        if l:
            val += "left key = {}, ".format(l)
        r = self.right(i)
        if r:
            val += "right key = {}, ".format(r)
        print(val)
        

# ALDS1_9_A: 完全二分木
def main():
    n = int(input())
    a = [int(v) for v in input().split(" ")]
    h = Heap(a)
    for i in range(1, n+1):
        h.print(i)

if __name__ == '__main__':
    main()
