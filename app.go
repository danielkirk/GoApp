package main

type Rectangle struct {
	height int
	width  int
}

func calculateArea(rect Rectangle) int {
	return rect.height * rect.width
}

func calculateParameter(rect Rectangle) int {
	return (rect.height * 2) + (rect.width * 2)
}

func printInfo(rect Rectangle) {
	println("Area is ", calculateArea(rect))
	println("Parameter is ", calculateParameter(rect))
}

func main() {
	rect := Rectangle{2, 4}

	printInfo(rect)

	rect.height *= 2
	rect.height *= 2

	printInfo(rect)
}
