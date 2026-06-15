--[[
  Source: https://www.codewars.com/kata/517abf86da9663f1d2000003/train/lua

  Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case). The next words should be always capitalized.
Examples

  "the-stealth-warrior" gets converted to "theStealthWarrior"

  "The_Stealth_Warrior" gets converted to "TheStealthWarrior"

  "The_Stealth-Warrior" gets converted to "TheStealthWarrior"
]]
return function(str)
  if #str == 0 then
    return str
  end

  local words = {}
  local currentWord = ""

  for i = 1, #str do
    local char = str:sub(i, i)

    if char == "-" or char == "_" then
      table.insert(words, currentWord)
      currentWord = ""
    else
      if currentWord == "" and #words > 0 then
        currentWord = string.upper(char)
      else
        currentWord = currentWord .. char
      end
    end
  end

  table.insert(words, currentWord)

  if #words == 1 then
    local word = words[1]
    local firstChar = word:sub(1, 1):upper()
    local remain = word:sub(2, #word)

    return firstChar .. remain
  end

  return table.concat(words, "")
end
