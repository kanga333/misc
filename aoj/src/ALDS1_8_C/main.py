class Node:
    def __init__(self, value, parent):
        self.value = value
        self.parent = parent
        self.left = None
        self.right = None

    def insert(self, value):
        next = None
        if value < self.value:
            if self.left:
                self.left.insert(value)
            else:
                self.left = Node(value, self)
        else:
            if self.right:
                self.right.insert(value)
            else:
                self.right = Node(value, self)
    
    def preorder(self):
        print(f" {self.value}", end="")
        if self.left:
            self.left.preorder()
        if self.right:
            self.right.preorder()

    def inorder(self):
        if self.left:
            self.left.inorder()
        print(f" {self.value}", end="")
        if self.right:
            self.right.inorder()
    
    def find(self, value):
        if self.value == value:
            return self
        if value < self.value:
            if self.left:
                return self.left.find(value)
        else:
            if self.right:
                return self.right.find(value)
        return None
    
    def next(self):
        if self.left:
            return self.left.next()
        return self
    
    def delete(self):
        if self.right and self.left:
            next = self.right.next()
            self.value = next.value
            next.delete()
            return      
        tmp = None
        if self.right:
            tmp = self.right
        if self.left:
            tmp = self.left
        p = self.parent
        if tmp:
            tmp.parent = p
        if p.right and p.right.value == self.value:
            p.right = tmp
        else:
            p.left = tmp
        return

# ALDS1_8_C: 二分探索木の削除
def main():
    root = None
    n = int(input())
    for _ in range(0, n):
        cmd = input()
        if cmd == "print":
            root.inorder()
            print()
            root.preorder()
            print()
            continue
        cmds = cmd.split(" ")
        value = int(cmds[1])
        if cmds[0] == "find":
            if root.find(value):
                print("yes")
            else:
                print("no")
        elif cmds[0] == "delete":
            target = root.find(value)
            if target:
                target.delete()
        else:
            if root is None:
                root = Node(value, None)
            else:
                root.insert(value)

if __name__ == '__main__':
    main()
