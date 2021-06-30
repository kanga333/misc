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
            return True
        if value < self.value:
            if self.left:
                return self.left.find(value)
        else:
            if self.right:
                return self.right.find(value)
        return False

# ALDS1_8_B: 二分探索木の探索
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
        else:
            if root is None:
                root = Node(value, None)
            else:
                root.insert(value)

if __name__ == '__main__':
    main()
