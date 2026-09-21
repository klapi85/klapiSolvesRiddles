fun main(args: Array<String>) {
    val input = "The quick brown fox jumps over the lazy dog" // true
    println( isPangram(input) )

    val input2 = "The quick brown fox jumps over the false doll" // false
    println( isPangram(input2) )
}

fun isPangram(input: String) : Boolean {
    val regex = """.""".toRegex()
    val result = regex.findAll(input).map{it.value.toLowerCase()}.sorted().distinct()
    println(result.toList().toString()) // debug
    return ("[ , a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z]" == result.toList().toString())
}
