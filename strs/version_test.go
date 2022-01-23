package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareVersion(t *testing.T) {
	ret := CompareVersion("8.1", "8.1")
	assert.Equal(t, 0, ret)
	ret = CompareVersion("8.1", "8.01")
	assert.Equal(t, 0, ret)
	ret = CompareVersion("8.21", "8.01")
	assert.Equal(t, 1, ret)
	ret = CompareVersion("8.101", "8.201")
	assert.Equal(t, -1, ret)
	ret = CompareVersion("8.1.2", "8.01")
	assert.Equal(t, 1, ret)
	ret = CompareVersion("7.0.09.000", "7.0.09")
	assert.Equal(t, 0, ret)
	ret = CompareVersion("7.0.08.9999", "7.0.09.9999")
	assert.Equal(t, -1, ret)
	ret = CompareVersion("7.0.9.9999", "7.0.09.999")
	assert.Equal(t, 1, ret)
	ret = CompareVersion("9.01", "9.0")
	assert.Equal(t, 1, ret)
	ret = CompareVersion("7.0.9.9999", "6.6.0.0000")
	assert.Equal(t, 1, ret)
	ret = CompareVersion("", "9")
	assert.Equal(t, -1, ret)
}

func TestCompareVersionWithCache(t *testing.T) {
	ret := CompareVersionWithCache("8.1", "8.1")
	assert.Equal(t, 0, ret)
	ret = CompareVersionWithCache("8.1", "8.01")
	assert.Equal(t, 0, ret)
	ret = CompareVersionWithCache("8.21", "8.01")
	assert.Equal(t, 1, ret)
	ret = CompareVersionWithCache("8.101", "8.201")
	assert.Equal(t, -1, ret)
	ret = CompareVersionWithCache("8.1.2", "8.01")
	assert.Equal(t, 1, ret)
	ret = CompareVersionWithCache("7.0.09.000", "7.0.09")
	assert.Equal(t, 0, ret)
	ret = CompareVersionWithCache("7.0.08.9999", "7.0.09.9999")
	assert.Equal(t, -1, ret)
	ret = CompareVersionWithCache("7.0.9.9999", "7.0.09.999")
	assert.Equal(t, 1, ret)
	ret = CompareVersionWithCache("9.01", "9.0")
	assert.Equal(t, 1, ret)
	ret = CompareVersionWithCache("7.0.9.9999", "6.6.0.0000")
	assert.Equal(t, 1, ret)
	ret = CompareVersionWithCache("", "9")
	assert.Equal(t, -1, ret)
}

func BenchmarkCompareVersion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CompareVersion("7.0.09.000", "7.0.09")
		CompareVersion("7.0.08.9999", "7.0.09.9999")
		CompareVersion("9.01", "9.0")
	}
}

func BenchmarkCompareVersionWithCache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CompareVersionWithCache("7.0.09.000", "7.0.09")
		CompareVersionWithCache("7.0.08.9999", "7.0.09.9999")
		CompareVersionWithCache("9.01", "9.0")
	}
}
