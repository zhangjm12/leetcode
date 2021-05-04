from typing import List

def rob(nums:List[int])->int:
    n = len(nums)
    r = [0] * n

    if n == 0:
        return 0
    if n == 1:
        return nums[0]

    r[0] = nums[0]
    r[1] = max(nums[0], nums[1])

    for i in range(2, n):
        r[i] = max(nums[i] + r[i-2], r[i-1])
    
    return r[n-1]


def rob1(nums: List[int]) -> int:
    n = len(nums)

    if n == 1:
        return nums[0]
    if n == 2:
        return max(nums[0], nums[1])

    return max(rob_range(nums, 0, n - 2), rob_range(nums, 1, n - 1))


def rob_range(nums: List[int], start: int, end: int) -> int:
    r = [0] * (end - start + 1)

    r[0] = nums[start]
    r[1] = nums[start+1]
    for i in range(2, end-start+1):
        r[i] = max(nums[start+i] + r[i-2], r[i-1])

    return r[end-start]


if __name__ == '__main__':
    test = [3, 1, 2, 5, 3, 1, 4, 11]
    print(rob(test))
    print(rob1(test))
