class Node:
    def __init__(self, id):
        self.id = id
        self.parent = -1
        self.depth = 0
        self.children = []

    def type(self):
        if self.parent == -1:
            return "root"
        if self.children == []:
            return "leaf"
        return "internal node"

    def __str__(self):
        return "node {}: parent = {}, depth = {}, {}, {}".format(
            self.id, self.parent, self.depth, self.type(), self.children
        )


def set_depth(node, nodes, depth):
    node.depth = depth
    for c in node.children:
        set_depth(nodes[c], nodes, depth+1)

# ALDS1_7_A: 根付き木
def main():
    n = int(input())
    nodes = [ Node(i) for i in range(0, n)]
    for _ in range(0, n):
        items = [int(v) for v in input().split(" ")]
        id = items[0]
        children = items[2:]
        node = nodes[id]
        node.children = children
        for c in children:
            nodes[c].parent = id
    root = None
    for n in nodes:
        if n.parent == -1:
            root = n
            break
    set_depth(root, nodes, 0)
    for n in nodes:
        print(str(n))
        

if __name__ == '__main__':
    main()
