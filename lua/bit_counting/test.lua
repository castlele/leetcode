require("busted.runner")()

local solution = require "solution"

describe("Fixed Tests", function()
  local input = { 0, 4, 7, 9, 10, 26, 77231418, 12525589, 3811, 392902058, 1044, 10030245, 183337941, 20478766, 103021, 287, 115370965, 31, 417862, 626031, 89, 674259 }
  local expected = { 0, 1, 3, 2, 2, 3, 14, 11, 8, 17, 3, 10, 16, 14, 9, 6, 15, 5, 7, 12, 4, 10 }
  for i = 1, #input do
    it("Basic test " .. i,
      function() assert.are.same(expected[i], solution.count_bits(input[i])) end)
  end
end)
