package experiment

import kotlin.test.assertEquals
import org.junit.Test

class TestExample {
	@Test
	fun testSayHi() {
		val expected = "Hi"
		assertEquals(expected, sayHi())
	}
}
