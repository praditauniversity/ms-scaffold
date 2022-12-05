package service

import (
	"ms-scaffold/api/models/domain"
	"ms-scaffold/api/models/web"
	"ms-scaffold/api/repository"

	"github.com/mashingan/smapping"
)

type ScaffoldService interface {
	All(scaffold *domain.Scaffold, pagination *domain.Pagination) (*[]domain.Scaffold, error)
	Create(request web.ScaffoldRequest) (domain.Scaffold, error)
	FindById(id uint) (domain.Scaffold, error)
	Update(request web.ScaffoldUpdateRequest) (domain.Scaffold, error)
	Delete(id uint) error
}

type scaffoldService struct {
	scaffoldRepository repository.ScaffoldRepository
}

func NewScaffoldService(scaffoldRepo repository.ScaffoldRepository) ScaffoldService {
	return &scaffoldService{scaffoldRepository: scaffoldRepo}
}

func (service *scaffoldService) All(scaffold *domain.Scaffold, pagination *domain.Pagination) (*[]domain.Scaffold, error) {
	return service.scaffoldRepository.All(scaffold, pagination)
}

func (service *scaffoldService) Create(request web.ScaffoldRequest) (domain.Scaffold, error) {
	scaffold := domain.Scaffold{}

	err := smapping.FillStruct(&scaffold, smapping.MapFields(&request))
	if err != nil {

		return scaffold, err
	}

	return service.scaffoldRepository.Create(scaffold), nil
}

func (service *scaffoldService) Update(request web.ScaffoldUpdateRequest) (domain.Scaffold, error) {
	scaffold := domain.Scaffold{}
	res, err := service.scaffoldRepository.FindById(request.ID)
	if err != nil {
		return scaffold, err
	}
	err = smapping.FillStruct(&scaffold, smapping.MapFields(&request))
	if err != nil {
		return scaffold, err
	}
	scaffold.UserID = res.UserID

	return service.scaffoldRepository.Update(scaffold), nil
}

func (service *scaffoldService) FindById(id uint) (domain.Scaffold, error) {
	scaffold, err := service.scaffoldRepository.FindById(id)
	if err != nil {
		return scaffold, err
	}
	return scaffold, nil
}

func (service *scaffoldService) Delete(id uint) error {
	scaffold, err := service.scaffoldRepository.FindById(id)
	if err != nil {
		return err
	}
	service.scaffoldRepository.Delete(scaffold)
	return nil
}
