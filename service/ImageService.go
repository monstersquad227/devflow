package service

import (
	"devflow/model"
	"devflow/repository"
)

type ImageService struct {
	ImageRepository *repository.ImageRepository
}

func (i *ImageService) List(pageNumber, pageSize int) ([]*model.Image, error) {
	return i.ImageRepository.ListImages(pageNumber, pageSize)
}

func (i *ImageService) Count() (int, error) {
	return i.ImageRepository.CountImages()
}

func (i *ImageService) Create(image *model.Image) (int64, error) {
	return i.ImageRepository.CreateImage(image)
}

func (i *ImageService) Update(image *model.Image) (int64, error) {
	return i.ImageRepository.UpdateImage(image)
}

func (i *ImageService) Delete(id int) (int64, error) {
	return i.ImageRepository.DeleteImage(id)
}
