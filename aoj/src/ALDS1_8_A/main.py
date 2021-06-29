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

# ALDS1_8_A: 二分探索木の挿入
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
        else:
            value = int(cmd.split(" ")[1])
            if root is None:
                root = Node(value, None)
            else:
                root.insert(value)

if __name__ == '__main__':
    main()
