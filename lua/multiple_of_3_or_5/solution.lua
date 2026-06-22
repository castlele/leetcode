local solution = {}

--[[
Source: https://www.codewars.com/kata/514b92a657cdc65150000006/train/lua

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.

Additionally, if the number is negative, return 0.

Note: If a number is a multiple of both 3 and 5, only count it once.

Courtesy of projecteuler.net ([Problem 1](https://projecteuler.net/problem=1))
]]
function solution.solution(value)
  --[[
  x <- value
  d <- step
  a_1 <- d

  S_n = n * ((a_1 + a_n) / 2)
  a_n = a_1 + (n-1) * d
  n = math.floor((x - 1) / d)

  k, 2k, 3k, ..., mk


  10, 3, 3
  n = math.floor(9 / 3) = 3
  a_n = 3 + 2 * 3 = 9
  S_n = 3 * ((3 + 9) / 2) = 3 * 6 = 18
  ]]
  ---@param x number
  ---@param step number difference between a_n and a_n+1
  ---@return number
  function arithmetic_sum(x, step)
    local first_in_series = step
    local n = math.floor((x - 1) / step)
    local last_in_serice = first_in_series + (n - 1) * step

    return n * ((first_in_series + last_in_serice) / 2)
  end

  return arithmetic_sum(value, 3)
      + arithmetic_sum(value, 5)
      - arithmetic_sum(value, 3 * 5)
end

return solution
