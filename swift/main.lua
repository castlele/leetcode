local nums = {1, 2, 3, 4}
local prefix = {1, 1, 1, 1}
local suffix = {1, 1, 1, 1}
local result = {0, 0, 0, 0}

print("collecting prefix")
for i = 2, #nums do
  print(i, prefix[i - 1], nums[i - 1])
  prefix[i] = prefix[i - 1] * nums[i - 1]
end

print("collecting suffix")
for i = #nums - 1, 1, -1 do
  print(i, suffix[i + 1], nums[i + 1])
  suffix[i] = suffix[i + 1] * nums[i + 1]
end

print("collecting result")
for i = 1, #nums do
  result[i] = prefix[i] * suffix[i]
end

print("nums: ", vim.inspect(nums))
print("prefix: ", vim.inspect(prefix))
print("suffix: ", vim.inspect(suffix))
print("result: ", vim.inspect(result))
