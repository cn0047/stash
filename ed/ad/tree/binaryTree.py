class Node:
    def __init__(self, info):
        self.info = info
        self.left = None
        self.right = None
        self.level = None

    def __str__(self):
        return str(self.info)


class BinarySearchTree:
    def __init__(self):
        self.root = None

    def addNode(self, val):
        if self.root == None:
            self.root = Node(val)
        else:
            current = self.root
            while True:
                if val < current.info:
                    if current.left:
                        current = current.left
                    else:
                        current.left = Node(val)
                        break
                elif val > current.info:
                    if current.right:
                        current = current.right
                    else:
                        current.right = Node(val)
                        break
                else:
                    break


def main():
    t = 6
    # input = "1 2 5 3 6 4" # topView: 1 2 5 6
    input = "4 2 3 1 7 6" # topView: 1 2 4 7
    tree = BinarySearchTree()
    arr = list(map(int, input.split()))
    for i in range(t):
        tree.addNode(arr[i])
    topView(tree.root)


"""
topView prints "Top View" for BST.
"""
def topView(root):
    queue = [root]
    heightValues = {root.info: 0}
    result = {0: [root.info]}
    while len(queue):
        node = queue.pop(0)
        height = heightValues[node.info]
        if result.get(height):
            result[height].append(node.info)
        else:
            result[height] = [node.info]
        if node.left:
            queue.append(node.left)
            h = height-1
            heightValues[node.left.info] = h
        if node.right:
            queue.append(node.right)
            h = height+1
            heightValues[node.right.info] = h
    r = []
    for k in sorted(result): r.append(result[k][0])
    print(' '.join(str(x) for x in r))


main()
