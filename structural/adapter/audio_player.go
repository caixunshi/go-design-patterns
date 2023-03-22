package adapter

import "fmt"

type AudioPlayer struct {
}

func (a *AudioPlayer) play(fileName string) {
	fmt.Println("audio player：" + fileName)
}
