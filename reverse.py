class Solution:
    
    def expand_from_middle(self, s: str, left: int, right: int):
        if s == "" or left > right:
            return 0
        
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
       
        return right - left -1
    
    
    
    def longestPalindrome(self, s: str) -> str:
        start = 0
        end = 0
        
        for i in range(len(s)):
            len1 = self.expand_from_middle(s, i, i)
            len2 = self.expand_from_middle(s, i, i+1)
            length = max(len1, len2)
            
            if length > end - start:
                start = i - ((length -1)//2)
                end = i + length//2
        return s[start: end+1]