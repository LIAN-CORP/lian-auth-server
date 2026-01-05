package model

type ContentPage[T any] struct {
	TotalPage     int
	TotalElements int64
	PageNumber    int
	PageSize      int
	IsFirst       bool
	IsLast        bool
	Content       []T
}
