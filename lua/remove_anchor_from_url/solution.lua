return {
  --[[
  Source: https://www.codewars.com/kata/51f2b4448cadf20ed0000386/train/lua

  Complete the function/method so that it returns the url with anything after the anchor (#) removed.
  Examples:
  ```
  "www.codewars.com#about" --> "www.codewars.com"
  "www.codewars.com?page=1" -->"www.codewars.com?page=1"
  ```
  ]]
  removeUrlAnchor = function(s)
    for i = 1, #s do
      local char = string.sub(s, i, i)

      if char == "#" then
        return string.sub(s, 1, i - 1)
      end
    end

    return s
  end
}
