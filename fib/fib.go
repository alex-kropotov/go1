package fib

var cache = map[int]int{}

func fromCache(n int) int {
	val, ok := cache[n]
	if !ok {
		val = GetFibNoCache(n)
	}
	return val
}

func GetFibNoCache(number int) int {
	f := fib()
	res := 0
	for i := 0; i <= number; i++ {
		res = f()
	}
	return res
}

func GetFibWithCache(number int) int {
	res := 0
	for i := 0; i <= number; i++ {
		res = fromCache(i)
	}
	return res
}

func fib() func() int {
	first, second := 0, 1
	return func() int {
		defer func() {
			first, second = second, first + second
		}()
		return first
	}
}