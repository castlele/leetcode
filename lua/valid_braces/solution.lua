local kata = {}

--[[
https://www.codewars.com/kata/5277c8a221e209d3f6000b56/train/lua

Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.
What is considered Valid?

A string of braces is considered valid if all braces are matched with the correct brace.
Examples:
```
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
```
--]]
kata.validBraces = function(braces)
  local stack = {}

  for brace in braces:gmatch(".") do
    if stack[#stack] == brace then
      table.remove(stack)
    elseif brace == "(" then
      table.insert(stack, ")")
    elseif brace == "[" then
      table.insert(stack, "]")
    elseif brace == "{" then
      table.insert(stack, "}")
    else
      return false
    end
  end

  return #stack == 0
end

return kata
