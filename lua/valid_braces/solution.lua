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
  local opposites = {
    ["("] = ")",
    ["["] = "]",
    ["{"] = "}",
  }

  local function check_inner(start, stop)
    if start > stop then
      return true
    end

    local brace = string.sub(braces, start, start)
    local opposite = opposites[brace]

    if opposite == nil then
      return false
    end

    local occurrences = 1

    for j = start + 1, stop do
      local current = string.sub(braces, j, j)

      if current == opposite and occurrences == 1 then
        return check_inner(start + 1, j - 1) and check_inner(j + 1, stop)
      elseif current == opposite and occurrences > 1 then
        occurrences = occurrences - 1
      elseif current == brace then
        occurrences = occurrences + 1
      end
    end

    return false
  end

  return check_inner(1, #braces)
end

return kata
