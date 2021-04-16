package deferMany

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeferred_AddNoReturn(t *testing.T) {
	actual := false
	func() {
		d := New()
		defer d.Defer()
		d.Add(func() {
			actual = true
		})
	}()
	assert.True(t, actual)
}

func TestDeferred_AddWithReturn(t *testing.T) {
	actual := false
	func() {
		d := New()
		defer d.Defer()
		d.Add(func() {
			actual = true
		})
		d.Return()
	}()
	assert.False(t, actual)
}

func TestDeferred_ReturnDoesNotExecuteOnReturn(t *testing.T) {
	actual := false
	cancel := func() func() {
		d := New()
		defer d.Defer()
		d.Add(func() {
			actual = true
		})
		return d.Return()
	}()
	assert.False(t, actual)
	cancel()
	assert.True(t, actual)
}
