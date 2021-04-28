from sys import stdin

def link_node(earlier, later):
    earlier.next = later
    later.prev = earlier

class Node:
    def __init__(self, key):
        self.key = key
        self.prev = self
        self.next = self
class LinkedList:
    def __init__(self):
        self.first = Node("-1")
        self.last = Node("-1")
        link_node(self.first, self.last)
    
    def delete_first(self):
        link_node(self.first, self.first.next.next)
    
    def delete_last(self):
        link_node(self.last.prev.prev, self.last)
    
    def insert(self, x):
        n = Node(x)
        link_node(n ,self.first.next)
        link_node(self.first ,n)
    
    def delete(self, x):
        node = self.first.next
        while node.key != "-1":
            key = node.key
            if key == x:
                link_node(node.prev, node.next)
                break
            node = node.next
    
    def result(self):
        node = self.first.next
        result = []
        while node.key != "-1":
            result.append(node.key)
            node = node.next
        return " ".join(result)

# ALDS1_3_C: 連結リスト
def main():
    l = LinkedList()
    n = int(input())
    for _ in range(0, n):
        cmd = stdin.readline().strip()
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
    
    print(l.result())

if __name__ == '__main__':
    main()
