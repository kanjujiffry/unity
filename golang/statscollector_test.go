package statscollector_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	statscollector "github.com/Unity-Technologies/mz-recruitment-golang-statscollector-assignment"
)

const requestTimeout = 19000

func TestMedian(t *testing.T) {
}

func TestAverage(t *testing.T) {
}

/*
Run benchmarks with multiple different input sizes.
Also being careful to have the same amount of work on each test iteration.
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go "Traps for young players"
*/

var benchmarkSizes = []int{1e3, 1e4, 1e5}

func BenchmarkRecord(b *testing.B) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			t0 := time.Now()
			for i := 0; i < b.N; i++ {
				collector := statscollector.New(requestTimeout)
				for n := 0; n < size; n++ {
					collector.Record(r.Intn(requestTimeout))
				}
			}
			b.ReportMetric(float64(time.Since(t0))/float64(b.N)/float64(size), "ns/Record()")
		})
	}
}

func BenchmarkMedian(b *testing.B) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	for _, size := range benchmarkSizes {
		collector := statscollector.New(requestTimeout)
		for n := 0; n < size; n++ {
			collector.Record(r.Intn(requestTimeout))
		}
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				collector.Median()
			}
		})
	}
}

func BenchmarkAverage(b *testing.B) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	for _, size := range benchmarkSizes {
		collector := statscollector.New(requestTimeout)
		for n := 0; n < size; n++ {
			collector.Record(r.Intn(requestTimeout))
		}
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				collector.Average()
			}
		})
	}
}
