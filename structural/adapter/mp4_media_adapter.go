package adapter

type Mp4MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *Mp4MediaAdapter) play(fileName string) {
	m.advancedMediaPlayer.playMp4(fileName)
}
