require("busted.runner")()

local solution = require "solution"
describe("solution", function()
  it("test1", function()
    assert.are.same("Ths wbst s fr lsrs LL!",
      solution.disemvowel("This website is for losers LOL!"))
  end)
end)
