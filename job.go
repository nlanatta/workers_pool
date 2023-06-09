package main

type Job interface {
	ID() int
	Process() error
}
