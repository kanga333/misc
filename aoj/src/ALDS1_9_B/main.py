import math

class Heap:
    def __init__(self, a) -> None:
        self.a = a
        self.a.insert(0, None)

    def parent(self, i):
        pi = math.floor(i / 2)
        return pi
    
    def left(self, i):
        li = 2 * i
        if len(self.a) <= li:
            return None
        return li


    def right(self, i):
        ri = (2 * i) + 1
        if len(self.a) <= ri:
            return None
        return ri
    
    def maxHeapify(self, i):
        l = self.left(i)
        r = self.right(i)
        largest = i
        if l and self.a[l] > self.a[largest]:
            largest = l
        if r and self.a[r] > self.a[largest]:
            largest = r
        if largest != i:
            tmp = self.a[i]
            self.a[i] = self.a[largest]
            self.a[largest] = tmp
            self.maxHeapify(largest)
    
    def print(self):
        s = " ".join([str(v) for v in self.a[1:]])
        print(" ", s)

# ALDS1_9_B: 最大ヒープ
def main():
    n = int(input())
    a = [int(v) for v in input().split(" ")]
    h = Heap(a)
    for i in range(math.floor(len(h.a)/2), 0, -1):
        h.maxHeapify(i)
    h.print()

if __name__ == '__main__':
    main()
