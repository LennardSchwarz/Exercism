object ScrabbleScore {

    fun scoreLetter(c: Char): Int {
        val onePoint = setOf('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T')
        val twoPoints = setOf('D', 'G')
        val threePoints = setOf('B', 'C', 'M', 'P')
        val fourPoints = setOf('F', 'H', 'V', 'W', 'Y')
        val fivePoints = setOf('K')
        val eightPoints = setOf('J', 'X')
        val tenPoints = setOf('Q', 'Z')

        return when (c.uppercaseChar()) {
            in onePoint -> 1
            in twoPoints -> 2
            in threePoints -> 3
            in fourPoints -> 4
            in fivePoints -> 5
            in eightPoints -> 8
            in tenPoints -> 10
            else -> 0
        }
    }

    fun scoreWord(word: String) = word.map { scoreLetter(it) }.sum()

}
