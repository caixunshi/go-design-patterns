package adapter

type VlcMediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *VlcMediaAdapter) play(fileName string) {
	m.advancedMediaPlayer.playVlc(fileName)
}
