# 二分法查找，必须是有序的序列
def binary_search(value):
    nums = [1, 3, 9, 22, 41, 55, 63, 75, 88, 91, 98, 102]
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] > value:
            right = mid - 1
        elif nums[mid] < value:
            left = mid + 1
        else:
            return True
    return False


# 杨氏矩阵查找，必须是从左往右，从上往下的增量矩阵
def matrix_search(l, x):
    row_max = len(l) - 1
    row, col = 0, len(l[0]) - 1

    while col >= 0 and row <= row_max:
        value = l[row][col]
        if value == x:
            return True
        elif value > x:
            col -= 1
        elif value < x:
            row += 1
    return False


if __name__ == '__main__':
    print(binary_search(88))
