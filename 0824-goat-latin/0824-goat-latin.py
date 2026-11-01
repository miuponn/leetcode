class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        words = sentence.split()                    # declare result array, vowel set
        res = []                                        
        vowels = set("aeiouAEIOU")                  # split str into array of words

        for index, word in enumerate(words):        # iterate over each index, word
            if word[0] not in vowels:               # if first letter in word is not vowel
                word = word[1:] + word[0]               # move consonant to end of word

            word = word + "ma" + "a" * (index + 1)  # add "ma" and "a" * index (starting from 1)
            res.append(word)                        # add new word to result array
        
        return " ".join(res)                        # join all words from result array into string and return