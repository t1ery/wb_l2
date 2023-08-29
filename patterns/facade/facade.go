package main

import "fmt"

// AudioFile - класс для работы с аудиофайлами
type AudioFile struct {
	filename string
}

func NewAudioFile(filename string) *AudioFile {
	return &AudioFile{filename: filename}
}

func (a *AudioFile) Load() {
	fmt.Println("Загрузка аудиофайла", a.filename)
}

func (a *AudioFile) Play() {
	fmt.Println("Воспроизведение аудиофайла", a.filename)
}

// VideoFile - класс для работы с видеофайлами
type VideoFile struct {
	filename string
}

func NewVideoFile(filename string) *VideoFile {
	return &VideoFile{filename: filename}
}

func (v *VideoFile) Load() {
	fmt.Println("Загрузка видеофайла", v.filename)
}

func (v *VideoFile) Play() {
	fmt.Println("Воспроизведение видеофайла", v.filename)
}

// ImageFile - класс для работы с изображениями
type ImageFile struct {
	filename string
}

func NewImageFile(filename string) *ImageFile {
	return &ImageFile{filename: filename}
}

func (i *ImageFile) Load() {
	fmt.Println("Загрузка изображения", i.filename)
}

func (i *ImageFile) Display() {
	fmt.Println("Отображение изображения", i.filename)
}

// MultimediaFacade - фасад для упрощенного взаимодействия с мультимедийными файлами
type MultimediaFacade struct {
	audioFile *AudioFile
	videoFile *VideoFile
	imageFile *ImageFile
}

func NewMultimediaFacade(audioFilename, videoFilename, imageFilename string) *MultimediaFacade {
	return &MultimediaFacade{
		audioFile: NewAudioFile(audioFilename),
		videoFile: NewVideoFile(videoFilename),
		imageFile: NewImageFile(imageFilename),
	}
}

func (m *MultimediaFacade) PlayAudio() {
	m.audioFile.Load()
	m.audioFile.Play()
}

func (m *MultimediaFacade) PlayVideo() {
	m.videoFile.Load()
	m.videoFile.Play()
}

func (m *MultimediaFacade) DisplayImage() {
	m.imageFile.Load()
	m.imageFile.Display()
}

func main() {
	facade := NewMultimediaFacade("song.mp3", "movie.mp4", "image.jpg")

	fmt.Println("Проигрывание аудио:")
	facade.PlayAudio()

	fmt.Println("\nПроигрывание видео:")
	facade.PlayVideo()

	fmt.Println("\nОтображение изображения:")
	facade.DisplayImage()
}
