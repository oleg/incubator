package maze

class Row {
    val list: MutableList<List<Int>> = ArrayList()
    fun row(vararg cells: Int): Unit {
        list.add(cells.toList())
    }
}

class grid(
    row: Row.() -> Unit
) {
    row()
}


class Grid(
    val length: Int,
    val width: Int,
)
