local kata = {}

--[[
Source: https://www.codewars.com/kata/51b62bf6a9c58071c600001b/train/lua

Create a function taking a positive integer between 1 and 3999 (both included) as its parameter and returning a string containing the Roman Numeral representation of that integer.

Modern Roman numerals are written by expressing each digit separately starting with the leftmost digit and skipping any digit with a value of zero. There cannot be more than 3 identical symbols in a row.

In Roman numerals:

    1990 is rendered: 1000=M + 900=CM + 90=XC; resulting in MCMXC.
    2008 is written as 2000=MM, 8=VIII; or MMVIII.
    1666 uses each Roman symbol in descending order: MDCLXVI.

Example:

   1 -->       "I"
1000 -->       "M"
1666 --> "MDCLXVI"

Help:

Symbol    Value
I          1
V          5
X          10
L          50
C          100
D          500
M          1,000

[More about roman numerals](https://en.wikipedia.org/wiki/Roman_numerals)
]]
kata.romanEncoder = function(number)
  local thousands = math.floor(number / 1000)
  local hundreds = math.floor((number - 1000 * thousands) / 100)
  local tenths = math.floor(((number - 1000 * thousands) - 100 * hundreds) / 10)
  local units = math.floor(((number - 1000 * thousands) - 100 * hundreds) -
    10 * tenths)

  local res = {}

  while thousands > 0 do
    table.insert(res, "M")
    thousands = thousands - 1
  end

  table.insert(res, kata.getRoman(hundreds, {
    "C",
    "CC",
    "CCC",
    "CD",
    "D",
    "DC",
    "DCC",
    "DCCC",
    "CM",
  }))
  table.insert(res, kata.getRoman(tenths, {
    "X",
    "XX",
    "XXX",
    "XL",
    "L",
    "LX",
    "LXX",
    "LXXX",
    "XC",
  }))
  table.insert(res, kata.getRoman(units, {
    "I",
    "II",
    "III",
    "IV",
    "V",
    "VI",
    "VII",
    "VIII",
    "IX",
  }))

  return table.concat(res, "")
end

---@param count integer
---@param alphabet table
---@return string
function kata.getRoman(count, alphabet)
  if count == 0 then
    return ""
  end

  return alphabet[count]
end

return kata
