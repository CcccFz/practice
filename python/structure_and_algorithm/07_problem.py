# 合并两个有序列表，类似归并排序的方式
def merge_sorted_list(l1, l2):
    temp = []
    while l1 and l2:
        if l1[0] >= l2[0]:
            temp.append(l2.pop(0))
        else:
            temp.append(l1.pop(0))
    while l1:
        temp.append(l1.pop(0))
    while l2:
        temp.append(l2.pop(0))
    return temp


# 台阶问题、2*n的矩形覆盖问题: 斐波那契数列
fib = lambda n: n if n <= 2 else fib(n - 1) + fib(n - 2)


# 也可以
def fib2(n):
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return b

# 变态台阶问题 fib = lambda n: n if n <= 2 else 2 * fib(n - 1)
