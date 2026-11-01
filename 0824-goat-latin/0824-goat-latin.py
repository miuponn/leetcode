class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        words = sentence.split()
        res = []
        vowels = set("aeiouAEIOU")

        for index, word in enumerate(words):
            if word[0] not in vowels:
                word = word[1:] + word[0]

            word = word + "maa" + ("a"*index)
            res.append(word)    
        
        return " ".join(res)