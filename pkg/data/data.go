package data

import "student/pkg/dto"

var Todos []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Name: "A"},
		{ID: 2, Name: "B"},
		{ID: 3, Name: "C"},
		{ID: 4, Name: "D"},
		{ID: 5, Name: "E"},
	}
}
