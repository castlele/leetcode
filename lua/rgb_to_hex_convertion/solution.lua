local kata = {}

--[[
[Source](https://www.codewars.com/kata/513e08acc600c94f01000001/train/lua)

The rgb function is incomplete. Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned. Valid decimal values for RGB are 0 - 255. Any values that fall out of that range must be rounded to the closest valid value.

Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.

Examples (input --> output):

```txt
255, 255, 255 --> "FFFFFF"
255, 255, 300 --> "FFFFFF"
0, 0, 0       --> "000000"
148, 0, 211   --> "9400D3"
```
]]
function kata.rgb(r, g, b)
  local rVal = math.max(math.min(r, 255), 0)
  local gVal = math.max(math.min(g, 255), 0)
  local bVal = math.max(math.min(b, 255), 0)

  return kata.toHex(rVal) .. kata.toHex(gVal) .. kata.toHex(bVal)
end

---@param num integer 0...255 integer value
---@return string
function kata.toHex(num)
  return string.format("%02X", num)
end

return kata
