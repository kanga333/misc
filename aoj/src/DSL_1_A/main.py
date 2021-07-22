class Node:
    def __init__(self, id) -> None:
        self.parent = self
        self.id = id
        self.rank = 0
    
    def root(self):
        if self.parent == self:
            if self.rank != 0:
                self.rank = 1
            return self
        parent = self.parent.root()
        self.parent = parent
        return parent
    
    def merge(self, parent):
        self.parent = parent
        if parent.rank < self.rank + 1:
            parent.rank = self.rank + 1

# DSL_1_A: 互いに素な集合
def main():
    a = [int(v) for v in input().split(" ")]
    n = a[0]
    q = a[1]

    nodes = [Node(i) for i in range(n)]
    for _ in range(q):
        a = [int(v) for v in input().split(" ")]
        com = a[0]
        x = nodes[a[1]]
        y = nodes[a[2]]
        if com == 0:
            xr = x.root()
            yr = y.root()
            if xr.rank >= yr.rank:
                yr.merge(xr)
            else:
                xr.merge(yr)
        else:
            if x.root().id == y.root().id:
                print(1)
            else:
                print(0)

if __name__ == '__main__':
    main()
