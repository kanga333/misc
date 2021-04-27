class Node:
    def __init__(self, key):
        self.key = key
        self.prev = None
        self.next = None
    
    def set_next(self, x):
        self.next = x
    
    def set_prev(self, x):
        self.prev = x


class NodeQueue:
    def __init__(self):
        self.first = None
        self.last = None
    
    def delete_first(self):
        if self.first:
            next = self.first.next
            if next:
                next.set_prev(None)
            self.first = next
        
        if self.first == None:
            self.last = None
    
    def delete_last(self):
        if self.last:
            prev = self.last.prev
            if prev:
                prev.set_next(None)
            self.last = prev
        
        if self.last == None:
            self.first = None
    
    def insert(self, x):
        n = Node(x)
        n.set_next(self.first)
        if self.first:
            self.first.set_prev(n)
        self.first = n

        if self.last == None:
            self.last = n
    
    def delete(self, x):
        node = self.first
        while node != None:
            key = node.key
            if key == x:
                prev = node.prev
                next = node.next
                if prev:
                    prev.set_next(next)
                else:
                    self.first = next
                if next:
                    next.set_prev(prev)
                else:
                    self.last = prev
                break
            node = node.next
    
    def print(self):
        node = self.first
        result = []
        while node != None:
            result.append(node.key)
            node = node.next
        print(" ".join(result))

# ALDS1_3_C: 連結リスト
def main():
    l = NodeQueue()
    n = int(input())
    for i in range(0, n):
        cmd = input()
        if cmd == "deleteFirst":
            l.delete_first()
            continue
        if cmd == "deleteLast":
            l.delete_last()
            continue
        splited = cmd.split(" ")
        target = splited[1]
        if splited[0] == "delete":
            l.delete(target)
            continue
        # insert
        l.insert(target)
    
    l.print()

if __name__ == '__main__':
    main()
