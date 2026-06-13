require("busted.runner")()
local high_and_low = require "solution".high_and_low

describe("Example Tests", function()
  it("Test 1", function()
    assert.are.same("42 -9", high_and_low "8 3 -5 42 -1 0 0 -9 4 7 4 -4")
  end)
  it("Test 2", function()
    assert.are.same("3 1", high_and_low "1 2 3")
  end)
end)
