package adapter

import "testing"

func TestAdapter(t *testing.T) {
	// user audio play
	var mediaPlayer MediaPlayer
	mediaPlayer = &AudioPlayer{}
	mediaPlayer.play("map3")

	// user VlcMediaAdapter play
	mediaPlayer = &VlcMediaAdapter{
		advancedMediaPlayer: &VlcPlayer{},
	}
	mediaPlayer.play("vlc")

	// user Mp4MediaAdapter play
	mediaPlayer = &Mp4MediaAdapter{
		advancedMediaPlayer: &Mp4Player{},
	}
	mediaPlayer.play("mp4")
}
