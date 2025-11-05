package com.castlelecs.algo

import kotlin.math.sqrt
import kotlin.math.pow

// source: https://coderun.yandex.ru/problem/triangles/description
fun findTriangles(dots: List<Pair<Int, Int>>): Int {
    val path = mutableListOf<Pair<Int, Int>>()
    var triangles = 0

    fun backtrack(start: Int) {
        if (path.size == 3) {
            val a = path[0]
            val b = path[1]
            val c = path[2]

            if (isIsoscelesTriangle(a, b, c)) triangles++

            return
        }

        for (i in start..<dots.size) {
            path.add(dots[i])

            backtrack(i + 1)

            path.removeLast()
        }
    }

    backtrack(0)

    return triangles
}

private fun isIsoscelesTriangle(a: Pair<Int, Int>, b: Pair<Int, Int>, c: Pair<Int, Int>): Boolean {
    val d1 = a squaredDistanceTo b
    val d2 = a squaredDistanceTo c
    val d3 = b squaredDistanceTo c

    val s = area(a, b, c)
    if (s == 0L) return false

    return d1 == d2 || d1 == d3 || d2 == d3
}

private infix fun Pair<Int, Int>.squaredDistanceTo(other: Pair<Int, Int>): Long {
    val dx = (other.first - first).toLong()
    val dy = (other.second - second).toLong()
    return dx * dx + dy * dy
}

private fun area(a: Pair<Int, Int>, b: Pair<Int, Int>, c: Pair<Int, Int>): Long {
    return kotlin.math.abs(
        (a.first * (b.second - c.second)
        + b.first * (c.second - a.second)
        + c.first * (a.second - b.second))
        .toLong()
    )
}
