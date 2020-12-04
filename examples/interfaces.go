package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphone int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Declaring interface and function which could use both types as input (TapePlayer and TapeRecorder)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Survivor","Horror","Numb"}
	var player Player = TapePlayer{}
	playList(player, mixtape)
	player = TapeRecorder{}
	playList(player, mixtape)
}
