 """
 Source: https://leetcode.com/problems/reverse-integer/
 Author: Lin Tanghui
 Date  : 2015/11/1
 """

 """
 Reverse digits of an integer.
 Example1: x = 123, return 321
 Example2: x = -123, return -321 
 """
class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
       
        flag=0
        if x<0:
            flag=1
            x=0-x
        x = str(x)
        temp =x
        lenth = len(temp)-1
        for i in range (len(temp)):       
                temp=temp[:lenth]+x[i]+temp[lenth+1:]
                lenth=lenth-1
               
        print sys.maxint
        if flag==1:
             return  0 if -int(temp)<-2147483648 else 0-int(temp)
        else:
            return  0 if int(temp)>2147483647  else int(temp)