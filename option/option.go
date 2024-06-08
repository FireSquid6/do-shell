// helper package because null pointers were made by satan
package option

import "fmt"

type Option[T any] struct {
	value  T
	isFull bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, isFull: true}
}

func None[T any]() Option[T] {
	return Option[T]{isFull: false}
}

func (s *Option[T]) Get() T {
	if !s.isFull {
		panic("Option is empty")
	}
	return s.value
}

func (s *Option[T]) GetOrElse(defaultValue T) (T, bool) {
	if !s.isFull {
		return defaultValue, true
	}
	return s.value, false
}

func (s *Option[T]) IsFull() bool {
	return s.isFull
}

func (s *Option[_]) IsEmpty() bool {
	return !s.isFull
}

func (s *Option[T]) String() string {
	if !s.isFull {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", s.value)
}
