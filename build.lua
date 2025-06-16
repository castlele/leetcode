---@diagnostic disable

---@return string
local function getCurrentFileName()
   local filePath = vim.fn.expand("%")
   local pathComponents = vim.fn.split(filePath, "/")

   return pathComponents[#pathComponents-1]
end

conf = {
   testcur = string.format('go test ./go/%s -coverpkg=./...', getCurrentFileName()),
   test = "go test ./go/... -coverpkg=./...",
   run = "go run main.go",
   gofmt = "go fmt ./go/...",
}
