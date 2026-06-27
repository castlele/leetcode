require("busted.runner")()

local solution = require "solution"

describe("RGB To Hex", function()
  local tests = {
    {
      output = "ADFF2F",
      input = {
        r = 173,
        g = 255,
        b = 47,
      },
    },
    {
      output = "FFFFFF",
      input = {
        r = 300,
        g = 255,
        b = 255,
      },
    },
    {
      output = "090D05",
      input = {
        r = 9,
        g = 13,
        b = 5,
      },
    },
    {
      output = "000000",
      input = {
        r = 0,
        g = 0,
        b = -20,
      },
    },
    {
      output = "000000",
      input = {
        r = 0,
        g = 0,
        b = 0,
      },
    },
  }

  for index, test in ipairs(tests) do
    it(string.format("Test: %i", index), function()
      local input = test.input
      assert.are.same(test.output, solution.rgb(input.r, input.g, input.b))
    end)
  end
end)
