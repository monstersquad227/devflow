package service

import (
	"devflow/model"
	"devflow/repository"
)

type ImageService struct {
	ImageRepository *repository.ImageRepository
}

func (i *ImageService) FetchImages(pageNumber, pageSize int) ([]*model.Image, error) {
	return i.ImageRepository.GetImages(pageNumber, pageSize)
}

func (i *ImageService) Fetch(pageNumber, pageSize int) ([]*model.Image, error) {
	return i.ImageRepository.GetImages(pageNumber, pageSize)
}

func (i *ImageService) FetchImagesCount() (int, error) {
	return i.ImageRepository.GetImagesCount()
}

func (i *ImageService) SaveImage(image model.Image) (int64, error) {
	return i.ImageRepository.CreateImage(image)
}

func (i *ImageService) RemoveImage(id int) (int64, error) {
	return i.ImageRepository.DeleteImage(id)
}

func (i *ImageService) ModifyImage(image model.Image) (int64, error) {
	return i.ImageRepository.UpdateImage(image)
}
