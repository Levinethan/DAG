package main

type Job struct {
	tasks []func()
	squential bool
	onComplete func()
}