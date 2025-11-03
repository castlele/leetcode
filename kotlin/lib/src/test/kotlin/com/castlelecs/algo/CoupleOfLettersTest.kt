package com.castlelecs.algo

import kotlin.test.Test

class CoupleOfLettersTest {
    @Test
    fun `Test case 1`() {
        val text = "ABCABC A"
        val exp = "BC"
        val sut = ::coupleOfLetters

        val res = sut(text)

        assert(res == exp, { "Got: $res, Exp: $exp" })
    }
}
