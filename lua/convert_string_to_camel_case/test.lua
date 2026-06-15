require("busted.runner")()
local toCamelCase = require "solution"

describe("Tests", function () 
    it("test", function () 
        assert.equal('', toCamelCase(''), "An empty string was provided but not returned")    
        assert.equal("theStealthWarrior", toCamelCase("the_stealth_warrior"), "toCamelCase('the_stealth_warrior') did not return correct value")    
        assert.equal("TheStealthWarrior", toCamelCase("The-Stealth-Warrior"), "toCamelCase('The-Stealth-Warrior') did not return correct value")    
        assert.equal("ABC", toCamelCase("A-B-C"), "toCamelCase('A-B-C') did not return correct value")
        assert.equal("GMxgsuzTwlqpQybfbXttazjsmFbzsdrgwo", toCamelCase("gMxgsuzTwlqpQybfbXttazjsmFbzsdrgwo"), "toCamelCase('gMxgsuzTwlqpQybfbXttazjsmFbzsdrgwo') did not return correct value")
        assert.equal("PnObkoTcexlutzFyVzqmzxsJazlEfhgkFzrhgxPbzfnzdfoPnazb", toCamelCase("pnObkoTcexlutzFyVzqmzxsJazlEfhgkFzrhgxPbzfnzdfoPnazb"), "toCamelCase('pnObkoTcexlutzFyVzqmzxsJazlEfhgkFzrhgxPbzfnzdfoPnazb') did not return correct value")
        assert.equal("FowyOppme", toCamelCase("fowyOppme"), "toCamelCase('pnObkoTcexlutzFyVzqmzxsJazlEfhgkFzrhgxPbzfnzdfoPnazb') did not return correct value")
    end)
end)
