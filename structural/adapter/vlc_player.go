package adapter

import "fmt"

type VlcPlayer struct {
}

func (*VlcPlayer) playVlc(fileName string) {
	fmt.Println("Play vlc file: " + fileName)
}
func (*VlcPlayer) playMp4(fileName string) {
	panic("unsupported file format")
}
