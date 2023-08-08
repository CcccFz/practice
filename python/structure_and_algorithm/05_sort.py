# 冒泡排序
def bubble_sort(items):
    length = len(items)
    for i in range(length-1):
        for j in range(length-1-i):
            if items[j] > items[j+1]:
                items[j], items[j+1] = items[j+1], items[j]


# 选择排序
def select_sort(items):
    length = len(items)
    for i in range(length-1):
        min_idx = i
        for j in range(i+1, length):
            if items[j] < items[min_idx]:
                min_idx = j
        items[min_idx], items[i] = items[i], items[min_idx]


# 快速排序
def quick_sort(items):
    if len(items) < 2:
        return items

    pivot = items[0]
    less_items = [x for x in items if x < pivot]
    greater_items = [x for x in items if x > pivot]
    center_items = [x for x in items if x == pivot]
    return quick_sort(less_items) + center_items + quick_sort(greater_items)


# 插入排序
def insert_sort(items):
    length = len(items)
    for i in range(1, length):
        for j in range(i):
            if items[i] < items[j]:
                items.insert(j, items[i])
                items.pop(i+1)
                break


# 希尔排序
def shell_sort(items):
    length = len(items)
    gap = length // 2
    while gap > 0:
        for i in range(gap, length):
            j = i
            while j >= gap and items[j] < items[j-gap]:
                items[j], items[j-gap] = items[j-gap], items[j]
                j -= gap
        gap = gap // 2


# 归并排序
def merge_sort(items):
    length = len(items)
    if length < 2:
        return items

    def merge(left_items, right_items):
        merged_items = []
        while left_items and right_items:
            merged_items.append(left_items.pop(0) if left_items[0] <= right_items[0] else right_items.pop(0))
        while left_items:
            merged_items.append(left_items.pop(0))
        while right_items:
            merged_items.append(right_items.pop(0))
        return merged_items

    mid_idx = length // 2
    return merge(merge_sort(items[:mid_idx]), merge_sort(items[mid_idx:]))


if __name__ == '__main__':
    nums = [2, 7, 3, 6, 1, 5, 8, 7, 4]
    bubble_sort(nums)
    # select_sort(nums)
    # nums = quick_sort(nums)
    # insert_sort(nums)
    # shell_sort(nums)
    # nums = merge_sort(nums)
    print(nums)
