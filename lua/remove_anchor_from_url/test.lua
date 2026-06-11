require("busted.runner")()

local solution = require "solution"
describe("solution", function()
  it("test for something", function()
    assert.are.same("www.codewars.com",
      solution.removeUrlAnchor("www.codewars.com#about"))
  end)
end)
