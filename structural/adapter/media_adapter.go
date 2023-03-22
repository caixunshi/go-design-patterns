package adapter

type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) play(fileName string) {
	m.advancedMediaPlayer.playVlc(fileName)
}
