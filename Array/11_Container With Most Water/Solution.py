from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        max_area, first, last = 0, 0, len(height)-1
        while first < last:
            if height[first] < height[last]:
                max_area = max(max_area, height[first] * (last - first))
                first += 1
            else:
                max_area = max(max_area, height[last] * (last - first))
                last -= 1
        return max_area

sol = Solution()
height = [1,8,6,2,5,4,8,3,7]
res = sol.maxArea(height)
print(res)