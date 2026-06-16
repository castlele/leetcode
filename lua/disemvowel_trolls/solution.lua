kata = {}

--[[
Source: https://www.codewars.com/kata/52fba66badcd10859f00097e/train/lua

Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
]]
function kata.disemvowel(s)
  local vowels = {
    a = true,
    e = true,
    i = true,
    o = true,
    u = true,
  }

  local censored = {}

  for i = 1, #s do
    local char = s:sub(i, i)

    if not vowels[char:lower()] then
      table.insert(censored, char)
    end
  end

  return table.concat(censored, "")
end

return kata
