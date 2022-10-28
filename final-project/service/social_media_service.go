package service

import "github.com/cperdiansyah/golang-h8-dts/final-project/model/domain"

type SocialMediaService interface {
	Create(socialMedia *domain.SocialMedia) (err error)
	GetAll() (socialMedias []domain.SocialMedia, err error)
	GetOne(id uint) (socialMedia domain.SocialMedia, err error)
	Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error)
	Delete(id uint) (err error)
}
