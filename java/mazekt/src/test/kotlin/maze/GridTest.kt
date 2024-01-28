package maze

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test


class GridTest {

    @Test
    fun `new grid`() {
        val grid = Grid(10, 9)

        assertEquals(10, grid.length)
        assertEquals(9, grid.width)
    }

}


