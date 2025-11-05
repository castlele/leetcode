package com.castlelecs.algo

import kotlin.test.Test
import kotlin.test.Ignore

class TrianglesTest {
    @Test
    fun `Test case 1`() {
        val input = listOf(
            0 to 0,
            2 to 2,
            -2 to 2,
        )
        val exp = 1
        val sut = ::findTriangles

        val res = sut(input)

        assert(res == exp, { "Exp: $exp; got: $res" })
    }

    @Test
    fun `Test case 2`() {
        val input = listOf(
            0 to 0,
            1 to 1,
            1 to 0,
            0 to 1,
        )
        val exp = 4
        val sut = ::findTriangles

        val res = sut(input)

        assert(res == exp, { "Exp: $exp; got: $res" })
    }
}
