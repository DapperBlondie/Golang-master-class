package Chapter7_2

import "testing"

var RESULT int

func BenchmarkFibo1(b *testing.B) {

	var r int
	for i := 0; i < b.N; i++ {

		r = fibo1(50)
	}

	RESULT = r
}

func BenchmarkFibo2(b *testing.B) {

	var r int
	for i := 0; i < b.N; i++ {

		r = fibo2(50)
	}

	RESULT = r
}

func BenchmarkFibo3(b *testing.B) {

	var r int
	for i := 0; i < b.N; i++ {

		r = fibo3(50)
	}

	RESULT = r
}
