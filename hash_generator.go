package main

import "github.com/google/uuid"

type StringHashGenerator interface {
	generate() string
}

type UUIDHashGenerator struct {
}

func (g UUIDHashGenerator) generate() string {
	return uuid.New().String()
}
