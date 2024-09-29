package facade

type MultimediaFacade struct {
	audioPlayer   *AudioPlayer
	videoPlayer   *VideoPlayer
	screenManager *ScreenManager
}

func NewMultimediaFacade() *MultimediaFacade {
	return &MultimediaFacade{
		audioPlayer:   &AudioPlayer{},
		videoPlayer:   &VideoPlayer{},
		screenManager: &ScreenManager{},
	}
}

func (m *MultimediaFacade) PlayMovie() {
	m.audioPlayer.PlayAudio()
	m.videoPlayer.PlayVideo()
	m.screenManager.ShowScreen()
}
