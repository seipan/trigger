package trigger

import (
	"fmt"
)

type Step[T any] struct {
	Name      string
	Condition func(T) bool
	Action    func(T) error
}

type Flow[T any] struct {
	steps        []Step[T]
	abortOnError bool
}

func NewFlow[T any]() *Flow[T] {
	return &Flow[T]{abortOnError: true}
}

func (f *Flow[T]) AbortOnError(flag bool) *Flow[T] {
	f.abortOnError = flag
	return f
}

func (f *Flow[T]) Add(name string, cond func(T) bool, action func(T) error) *Flow[T] {
	f.steps = append(f.steps, Step[T]{Name: name, Condition: cond, Action: action})
	return f
}

func (f *Flow[T]) Run(ctx T) error {
	var errs []error
	for _, s := range f.steps {
		if s.Condition == nil || s.Condition(ctx) {
			if err := s.Action(ctx); err != nil {
				if f.abortOnError {
					return fmt.Errorf("%s: %w", s.Name, err)
				}
				errs = append(errs, fmt.Errorf("%s: %w", s.Name, err))
				continue
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("multiple errors: %v", errs)
	}
	return nil
}
