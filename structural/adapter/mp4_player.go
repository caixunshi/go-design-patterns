package adapter

import "fmt"

type Mp4Player struct {
}

func (*Mp4Player) playVlc(fileName string) {
	panic("unsupported file format: " + fileName)
}
func (*Mp4Player) playMp4(fileName string) {
	fmt.Println("Play mp4 file: " + fileName)
}
