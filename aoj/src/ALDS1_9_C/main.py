import math

MAX_COMMAND = 2_000_000

class Heap:
    def __init__(self) -> None:
        self.a = [0 for _ in range(MAX_COMMAND+1)]
        self.length = 1

    def parent(self, i):
        pi = math.floor(i / 2)
        return pi
    
    def left(self, i):
        li = 2 * i
        if self.length <= li:
            return None
        return li


    def right(self, i):
        ri = (2 * i) + 1
        if self.length <= ri:
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

    def insert(self, val):
        self.a[self.length] = val
        i = self.length
        self.length += 1
        while(True):
            p = self.parent(i) 
            if p == 0:
                break
            if self.a[i] < self.a[p]:
                break
            tmp = self.a[i]
            self.a[i] = self.a[p]
            self.a[p] = tmp
            i = p

    def extract_max(self):
        max = self.a[1]
        self.length -= 1
        self.a[1] = self.a[self.length]
        self.maxHeapify(1)
        return max

# ALDS1_9_C: 優先度付きキュー
def main():
    h = Heap()
    while(True):
        cmd = input()
        if cmd == "end":
            break
        if cmd == "extract":
            print(h.extract_max())
            continue
        val = int(cmd.split(" ")[1])
        h.insert(val)

if __name__ == '__main__':
    main()
