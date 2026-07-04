local latinLatters = 26

--[[
Source: https://www.codewars.com/kata/545cedaa9943f7fe7b000048/train/lua

A pangram is a sentence that contains every single letter of the alphabet at least once. For example, the sentence "The quick brown fox jumps over the lazy dog" is a pangram, because it uses the letters A-Z at least once (case is irrelevant).

Given a string, detect whether or not it is a pangram. Return True if it is, False if not. Ignore numbers and punctuation.
]]
---@param s string
---@return boolean
local function is_pangram(s)
  local counter = 0
  local alphabet = {}

  for char in s:gmatch("%a") do
    local byte = char:lower():byte()

    if not alphabet[byte] then
      alphabet[byte] = true
      counter = counter + 1
    end

    if counter == latinLatters then
      return true
    end
  end

  return #alphabet == latinLatters
end

return is_pangram
