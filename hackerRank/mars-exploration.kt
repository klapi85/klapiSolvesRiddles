fun main(args: Array<String>) {
    val input = "SOSSOSSORXOX"
    val regex = """...""".toRegex()
    println( regex.findAll(input).map { checkTriple(it.value) }.sum())
}


private fun checkTriple(triple: String): Int {
    var i: Int = 0
    if ("S" == triple.substring(0, 1)) {
        i++
    }
    if ("O" == triple.substring(1, 2)) {
        i++
    }
    if ("S" == triple.substring(2, 3)) {
        i++
    }
    return i
}
