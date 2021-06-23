class Node:
    def __init__(self, id):
        self.id = id
        self.parent = None
        self.left = None
        self.right = None
        self.height = None
        self.depth = None

    def type(self):
        if self.parent is None:
            return "root"
        if self.left is None and self.right is None:
            return "leaf"
        return "internal node"
    
    def get_parent_id(self):
        if self.parent is None:
            return -1
        return self.parent.id
    
    def get_sibling_id(self):
        if self.parent is None:
            return -1
        if self.parent.left is not None and self.parent.left.id != self.id:
            return self.parent.left.id
        if self.parent.right is not None and self.parent.right.id != self.id:
            return self.parent.right.id
        return -1
    
    def get_degree(self):
        degree = 0
        if self.left is not None:
            degree += 1
        if self.right is not None:
            degree += 1
        return degree
    
    def get_height(self):
        if self.height is not None:
            return self.height
        
        height = 0
        if self.left is not None:
            height = self.left.get_height() + 1
        if self.right is not None:
            right_height = self.right.get_height() + 1
            if height < right_height:
                height = right_height
        self.height = height
        return height

    def get_depth(self):
        if self.depth is not None:
            return self.depth
        depth = 0
        if self.parent is not None:
            depth =  self.parent.get_depth() + 1
        self.depth = depth
        return depth

    def __str__(self):
        return "node {}: parent = {}, sibling = {}, degree = {}, depth = {}, height = {}, {}".format(
            self.id, self.get_parent_id(), self.get_sibling_id(), self.get_degree(), self.get_depth(), self.get_height(), self.type()
        )


# ALDS1_7_A: 2分木
def main():
    n = int(input())
    nodes = [ Node(i) for i in range(0, n)]
    for _ in range(0, n):
        items = [int(v) for v in input().split(" ")]
        id = items[0]
        left = items[1]
        right = items[2]
        node = nodes[id]
        if left != -1:
            left_node = nodes[left]
            left_node.parent = node
            node.left = left_node
        if right != -1:
            right_node = nodes[right]
            right_node.parent = node
            node.right = right_node
    for n in nodes:
        print(str(n))
        

if __name__ == '__main__':
    main()
