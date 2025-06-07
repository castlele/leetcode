---@diagnostic disable

---@return string
local function getCurrentFileName()
   local filePath = vim.fn.expand("%")
   local pathComponents = vim.fn.split(filePath, "/")

   return pathComponents[#pathComponents-2].."/"..pathComponents[#pathComponents-1].."/"..pathComponents[#pathComponents]
end

conf = {
   testcur = string.format('go test ./"%s" -coverpkg=./...', getCurrentFileName()),
}
