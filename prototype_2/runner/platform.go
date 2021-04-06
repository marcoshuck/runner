package runner

import "log"

type Platform interface {
	LaunchInstance() int
	RemoveInstance() error
}

type platform struct{}

func (p platform) RemoveInstance() error {
	log.Println("Removing instance...")
	return nil
}

func (p platform) LaunchInstance() int {
	log.Println("Launching instance...")
	return 1
}

func NewPlatform() Platform {
	return &platform{}
}
