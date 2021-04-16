package deferMany

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeferer_AddNoReturn(t *testing.T) {
	actual := false
	func() {
		deferable := New()
		defer deferable.Defer()
		deferable.Add(func() {
			actual = true
		})
	}()
	assert.True(t, actual)
}

func TestDeferer_AddWithReturn(t *testing.T) {
	actual := false
	func() {
		deferable := New()
		defer deferable.Defer()
		deferable.Add(func() {
			actual = true
		})
		deferable.Return()
	}()
	assert.False(t, actual)
}

func TestDeferer_ReturnDoesNotExecuteOnReturn(t *testing.T) {
	actual := false
	cancel := func() func() {
		deferable := New()
		defer deferable.Defer()
		deferable.Add(func() {
			actual = true
		})
		return deferable.Return()
	}()
	assert.False(t, actual)
	cancel()
	assert.True(t, actual)
}
