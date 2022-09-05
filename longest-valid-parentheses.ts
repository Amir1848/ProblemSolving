// Given a string containing just the characters '(' and ')'
// find the length of the longest valid (well-formed) parentheses substring.


// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".


// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".



// not correct yet
function longestValidParentheses(s: string): number {

    let longestValidString = 0;

    let counter = 0;
    let startIndex = 0;
    for(let i = 0; i < s.length; i++){
        
        if(s[i] == '('){
            counter++;
        }else{
            counter--;
        }
        
        
        if(counter < 0){
            
            let validLengthString: number = i - startIndex + 1; 
            if(validLengthString > longestValidString){
                longestValidString = validLengthString;
            }
            
            counter = 0;
            startIndex = i + 1;
        }
        
    }
    
    return longestValidString;

}