package experiment

import spark.Spark.get;

fun sayHi(): String {
	return "Hi"
}

fun main(args: Array<String>) {

	get("/hello") { req, res -> "Hello World" }

	get("/hi") { req, res -> sayHi() }
}


