package dsa

import "testing"

/*s
worst

cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkDuplicates/100-8         	  298772	      3391 ns/op	       0 B/op	       0 allocs/op
BenchmarkDuplicates/1000-8        	    4032	    289381 ns/op	       0 B/op	       0 allocs/op
BenchmarkDuplicates/100000-8      	       1	2849398968 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkDuplicates(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates(s1000)
		}

	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates(s100000)
		}

	})
}

/*
sorting and removing duplicates
BenchmarkDuplicates2/100-8         	  539743	      2168 ns/op	      24 B/op	       1 allocs/op
BenchmarkDuplicates2/1000-8        	   31281	     36546 ns/op	      24 B/op	       1 allocs/op
BenchmarkDuplicates2/100000-8      	     193	   5915067 ns/op	      24 B/op	       1 allocs/op
*/
func BenchmarkDuplicates3(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates3(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates3(s1000)
		}
	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates3(s100000)
		}
	})
}

/*
hashmap

BenchmarkDuplicates4/100-8         	  131220	      8648 ns/op	    5389 B/op	      25 allocs/op
BenchmarkDuplicates4/1000-8        	   12423	     95940 ns/op	   69693 B/op	      84 allocs/op
BenchmarkDuplicates4/100000-8      	     129	   9260453 ns/op	 8249819 B/op	    3765 allocs/op
*/
func BenchmarkDuplicates4(b *testing.B) {
	s100 := randNumSlice(100)
	s1000 := randNumSlice(1000)
	s100000 := randNumSlice(100000)

	b.ResetTimer()
	b.ReportAllocs()
	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates4(s100)
		}
	})
	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates4(s1000)
		}
	})
	b.Run("100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			duplicates4(s100000)
		}
	})
}
