class Node(object):
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right


tree = Node(1, Node(3, Node(7, Node(0)), Node(6)), Node(2, Node(5), Node(4)))
tree_bak = Node(1, Node(3, Node(7, Node(0)), Node(6)), Node(2, Node(5), Node(4)))


# 层次遍历
def lookup(root):
    row = [root]
    while row:
        print([x.data for x in row])
        row = [kid for item in row for kid in (item.left, item.right) if kid]


# 深度遍历
def deep(root):
    if not root:
        return
    print(root.data)
    deep(root.left)
    deep(root.right)


# 中序遍历
def mid(root):
    if not root:
        return False
    if root.left:
        mid(root.left)
    print(root.data)
    if root.right:
        mid(root.right)


# 前序遍历
def pre(root):
    if not root:
        return False
    print(root.data)
    if root.left:
        pre(root.left)
    if root.right:
        pre(root.right)


# 后续遍历
def last(root):
    if not root:
        return False
    if root.left:
        last(root.left)
    if root.right:
        last(root.right)
    print(root.data)


# 最大深度
def max_depth(root):
    if not root:
        return 0
    return max(max_depth(root.left), max_depth(root.right)) + 1


# 两树是否相同
def is_same(p, q):
    if not p and not q:
        return True
    elif p and q:
        return p.data == q.data and is_same(p.left, q.left) and is_same(p.right, q.right)
    else:
        return False


# 已知前序中序求后序
def rebuild(pre, center):
    root = Node(pre[0])
    idx = center.index(pre[0])  # 有了idx，即能得到左子树的中序遍历，同时得到左子树的节点个数，进而结合pre推出左子树的前序遍历

    left_tree_pre = pre[1:idx+1]
    left_tree_center = center[:idx]
    root.left = rebuild(left_tree_pre, left_tree_center)

    right_tree_pre = pre[idx+1:]
    right_tree_center = center[idx+1:]
    root.right = rebuild(right_tree_pre, right_tree_center)

    return root


if __name__ == '__main__':
    lookup(tree)
    # deep(tree)
    # mid(tree)
    # pre(tree)
    # last(tree)
    # print(max_depth(tree))
    # print('两树相等' if is_same(tree, tree_bak) else '两树不等')
