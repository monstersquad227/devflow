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

func (i *ImageService) FetchImagesCount() (int, error) {
	return i.ImageRepository.GetImagesCount()
}
