package models

import "math/rand"

type Button struct {
	id int
	label string
	image string
	audio string
}

func NewButton(label string) *Button {
	return &Button {
		id: rand.Intn(10000),
		label: label,
		image: "",
		audio: "",
	}
}
