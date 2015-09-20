# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.

# Algorithm

I used dynamic programming to solve finding the longest compound word. It basically finds the first subsequence in the word which is a actual word in the list and then tries to find all other words in the list. Also it uses an boolean array/slice to store which indices of the word des it form a compund word or a word itself. I make sure that if the word is not a compund word then the last element is false. Also if the last element in the boolean array is true then I return true. The running time of this algorithm is O(NM) where N is the number of words in the dictionary and M is the length of the words. It also takes O(M) memory to store the booleans for each word. Also it takes O(N) memory to create the dictionary. So memory space requirement is O(N+M)