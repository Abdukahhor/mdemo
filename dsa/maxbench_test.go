package dsa

import "testing"

/*
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkFindMax/100-8         	18238957	        55.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMax/1000-8        	 2007378	       590.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMax/100000-8      	   21228	     58118 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkFindMax2(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax2(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax2(s1000)
		}

	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax2(s100000)
		}

	})
}

/*
BenchmarkFindMax4/100-8         	15055790	        69.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMax4/1000-8        	 1716471	       730.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMax4/100000-8      	   17169	     71735 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkFindMax4(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax4(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax4(s1000)
		}

	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax4(s100000)
		}

	})
}

/*
BenchmarkFindMax6/100-8         	  568369	      2144 ns/op	      24 B/op	       1 allocs/op
BenchmarkFindMax6/1000-8        	   31206	     39028 ns/op	      24 B/op	       1 allocs/op
BenchmarkFindMax6/100000-8      	     195	   5946409 ns/op	      24 B/op	       1 allocs/op
*/
func BenchmarkFindMax6(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax6(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax6(s1000)
		}

	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMax6(s100000)
		}

	})
}
