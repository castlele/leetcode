--[[
Source: https://www.codewars.com/kata/546f922b54af40e1e90001da/train/lua

Welcome.

In this kata you are required to, given a string, replace every letter with its position in the alphabet.

If anything in the text isn't a letter, ignore it and don't return it.

"a" = 1, "b" = 2, etc.
Example
```
Input = "The sunset sets at twelve o' clock."
Output = "20 8 5 19 21 14 19 5 20 19 5 20 19 1 20 20 23 5 12 22 5 15 3 12 15 3 11"
```
]]
local function alphabet_position(text)
  local res = {}

  for char in text:gmatch("%a") do
    -- a is 97 bytes
    local byte = char:lower():byte()
    local place = byte - 96
    table.insert(res, place)
  end

  return table.concat(res, " ")
end

return alphabet_position
