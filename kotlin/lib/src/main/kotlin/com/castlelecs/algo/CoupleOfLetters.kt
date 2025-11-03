package com.castlelecs.algo

// source: https://coderun.yandex.ru/problem/couple-of-letters/description
fun coupleOfLetters(text: String): String {
    val words = text.split(" ")
    val frequencies = mutableMapOf<String, Int>()

    for (word in words) {
        for (i in 1..<word.length) {
            val sub = "${word[i-1]}${word[i]}"
            frequencies[sub] = (frequencies[sub] ?: 0) + 1
        }
    }

    var maxSubStr = ""
    var max = 0

    for (frequency in frequencies) {
        if (frequency.value > max || (frequency.value == max && frequency.key.compareTo(maxSubStr) > 0)) {
            maxSubStr = frequency.key
            max = frequency.value
        }
    }

    return maxSubStr
}
