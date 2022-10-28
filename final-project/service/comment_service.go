package service

import "github.com/cperdiansyah/golang-h8-dts/final-project/model/domain"

type CommentService interface {
	Create(comment *domain.Comment) (err error)
	GetAll() (comments []domain.Comment, err error)
	GetOne(id uint) (comment domain.Comment, err error)
	Update(comment domain.Comment) (photo domain.Photo, err error)
	Delete(id uint) (err error)
}
