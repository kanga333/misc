idx = 0

class Node:
    def __init__(self, id):
        self.id = id
        self.left = None
        self.right = None
    
    def preorder(self):
        print(f" {self.id}", end="")
        if self.left is not None:
            self.left.preorder()
        if self.right is not None:
            self.right.preorder()

    def inorder(self):
        if self.left is not None:
            self.left.inorder()
        print(f" {self.id}", end="")
        if self.right is not None:
            self.right.inorder()
    
    def postorder(self):
        if self.left is not None:
            self.left.postorder()
        if self.right is not None:
            self.right.postorder()
        print(f" {self.id}", end="")

def restore_tree(pre_order, in_order):
    global idx
    if len(in_order) == 0:
        return None
    id = pre_order[idx]
    t = Node(id)
    idx += 1
    for i, v in enumerate(in_order):
        if v == id:
            left = restore_tree(pre_order, in_order[:i])
            right= restore_tree(pre_order, in_order[i+1:])
            t.left = left
            t.right = right
            return t
    return None

# ALDS1_7_D: 木の復元
def main():
    _ = int(input())
    pre_order = [int(v) for v in input().split(" ")]
    in_order = [int(v) for v in input().split(" ")]
    t = restore_tree(pre_order, in_order)
    t.postorder()
        

if __name__ == '__main__':
    main()
