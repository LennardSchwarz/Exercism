object Hamming {

    fun compute(leftStrand: String, rightStrand: String) =
            if (leftStrand.length != rightStrand.length)
                throw IllegalArgumentException("left and right strands must be of equal length")
            else
                (leftStrand zip rightStrand).count { it.first != it.second }
}