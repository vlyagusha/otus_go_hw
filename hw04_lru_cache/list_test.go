package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("only 1 item", func(t *testing.T) {
		l := NewList()

		l.PushFront("Go, baby, go!")
		require.Equal(t, 1, l.Len())
		require.Equal(t, "Go, baby, go!", l.Front().Value)
		require.Equal(t, "Go, baby, go!", l.Back().Value)

		l = NewList()

		l.PushBack("Go, baby, go!")
		require.Equal(t, 1, l.Len())
		require.Equal(t, "Go, baby, go!", l.Front().Value)
		require.Equal(t, "Go, baby, go!", l.Back().Value)
	})

	t.Run("moving to front single item", func(t *testing.T) {
		l := NewList()

		l.PushFront(1.23)
		l.MoveToFront(l.Front())
		require.Equal(t, 1, l.Len())
		require.Equal(t, 1.23, l.Front().Value)
		require.Equal(t, 1.23, l.Back().Value)

		l.MoveToFront(l.Back())
		require.Equal(t, 1, l.Len())
		require.Equal(t, 1.23, l.Front().Value)
		require.Equal(t, 1.23, l.Back().Value)
	})

	t.Run("distinct type items", func(t *testing.T) {
		l := NewList()

		l.PushFront(1.23)
		l.PushFront(4)
		l.PushFront("Hello, Otus!")
		l.PushFront('q')

		require.Equal(t, 4, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []interface{}{1.23, 4, "Hello, Otus!", 'q'}, elems)
	})

	t.Run("items order", func(t *testing.T) {
		const maxItems = 10

		l := NewList()
		for i := 0; i < maxItems; i++ {
			l.PushFront(i)
		}
		for i := 0; i < maxItems; i++ {
			require.Equal(t, maxItems-i, l.Len())
			require.Equal(t, i, l.Front().Value)

			l.Remove(l.Front())
		}
	})

	t.Run("items reverse order", func(t *testing.T) {
		const maxItems = 10

		l := NewList()
		for i := 0; i < maxItems; i++ {
			l.PushBack(i)
		}
		for i := maxItems - 1; i >= 0; i-- {
			require.Equal(t, i, l.Len())
			require.Equal(t, maxItems-i, l.Front().Value)

			l.Remove(l.Front())
		}
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
