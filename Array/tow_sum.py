# encoding: utf-8
"""
@author: Shaohe Guo 
@contact: 842125706@qq.com
@site: www.guoshaohe.com

@file: test.py
@time: 2019/10/29 10:44

"""

"""
    1、Two Sum
        在数组中找到相加等于target的下标索引, 不能重复使用同一元素值
            Given nums = [2, 7, 11, 15], target = 9,
    
            Because nums[0] + nums[1] = 2 + 7 = 9,
            return [0, 1]
"""
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        if len(nums) <= 1:
            return False
        buff_dict = {}
        for index, value in enumerate(nums):
            tmp_value = target - value
            if value in buff_dict:
                return (buff_dict[value], index)
            else:
                buff_dict[tmp_value] = index




if __name__ == '__main__':
    sol = Solution()
    nums = [2, 7, 11, 15]
    target = 22
    res = sol.twoSum(nums, target)
    print(res)
