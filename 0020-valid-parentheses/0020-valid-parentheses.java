class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();                             // stack to track opening brackets (LIFO)
        Map<Character, Character> map = Map.of(')','(', '}','{', ']','[');  // map each closing bracket to its corresponding opening bracket

        for (int i = 0; i < s.length(); i++){                                 // iterate through each character in the string
            char c = s.charAt(i);               

            if (map.containsKey(c)) {                                       // if it's a closing bracket
                if (!stack.isEmpty() && stack.peek() == map.get(c)) {       // check if stack is not empty and the top matches expected opening
                    stack.pop();                                            // valid match found, remove opening bracket from stack
                } else {                                                    
                    return false;                                           // mismatch or unmatched closing bracket
                }
            } else {                                                        // if it's an opening bracket, push to stack
                stack.push(c);                                       
            }
        }
        return stack.isEmpty();                                             // if stack is empty, all brackets were balanced
    }
}