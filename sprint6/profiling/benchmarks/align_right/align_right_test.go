package align_right

import (
	"strconv"
	"testing"
)

func BenchmarkLoadBalancer(b *testing.B) {
	const stringCount = 100
	const triesCount = 1000

	strings := make([]string, stringCount)
	for i := 0; i < stringCount; i++ {
		strings = append(strings, strconv.Itoa(i))
	}

	b.ResetTimer()

	b.Run("naive", func(b *testing.B) {
		for _, s := range strings {
			AlignRight(s, triesCount, rune(12))
		}
	})
	b.Run("buffer", func(b *testing.B) {
		for _, s := range strings {
			AlignRightBuffer(s, triesCount, rune(12))
		}
	})
	b.Run("repeat", func(b *testing.B) {
		for _, s := range strings {
			AlignRightRepeat(s, triesCount, rune(12))
		}
	})
}
