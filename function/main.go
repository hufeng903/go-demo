package main

type myInt int

func (i myInt) add(a int) myInt {
	i = i + myInt(a)

	return i
}

func (i *myInt) add1(a int) myInt {
	*i = *i + myInt(a)

	return *i
}

func main() {
	i1 := myInt(1)

	println(i1)

	i1.add(1)

	println(i1)

	i1.add1(1)

	println(i1)
}
