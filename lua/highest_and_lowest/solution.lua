local solution = {}

--[[
Source: https://www.codewars.com/kata/554b4ac871d6813a03000035/train/lua

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.
Examples:
```txt
high_and_low "1 2 3 4 5" --> return "5 1"
high_and_low "1 2 -3 4 5" --> return "5 -3"
high_and_low "1 9 3 4 -5" --> return "9 -5"
```

Notes:
    All numbers are valid Int32, no need to validate them.
    There will always be at least one number in the input string.
    Output string must be two numbers separated by a single space, and highest number is first.
]]
function solution.high_and_low(numberstrings)
  local min, max

  for number in numberstrings:gmatch("[%-%d][%d]*") do
    number = tonumber(number)
    if min == nil and max == nil then
      min = number
      max = number
    end

    if number < min then
      min = number
    end

    if number > max then
      max = number
    end
  end

  return string.format("%d %d", max, min)
end

return solution
