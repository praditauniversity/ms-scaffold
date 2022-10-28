package service

import (
	"ms-scaffold/api/models/domain"
	"ms-scaffold/api/models/web"
	"ms-scaffold/api/repository"

	"github.com/mashingan/smapping"
)

type ScaffoldService interface {
	All() []domain.Scaffold
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

func (as *scaffoldService) All() []domain.Scaffold {
	return as.scaffoldRepository.All()
}

func (as *scaffoldService) Create(request web.ScaffoldRequest) (domain.Scaffold, error) {
	scaffold := domain.Scaffold{}

	err := smapping.FillStruct(&scaffold, smapping.MapFields(&request))
	if err != nil {

		return scaffold, err
	}

	return as.scaffoldRepository.Create(scaffold), nil
}

func (as *scaffoldService) Update(request web.ScaffoldUpdateRequest) (domain.Scaffold, error) {
	scaffold := domain.Scaffold{}
	res, err := as.scaffoldRepository.FindById(request.ID)
	if err != nil {
		return scaffold, err
	}
	err = smapping.FillStruct(&scaffold, smapping.MapFields(&request))
	if err != nil {
		return scaffold, err
	}
	scaffold.User_id = res.User_id

	return as.scaffoldRepository.Update(scaffold), nil
}

func (as *scaffoldService) FindById(id uint) (domain.Scaffold, error) {
	scaffold, err := as.scaffoldRepository.FindById(id)
	if err != nil {
		return scaffold, err
	}
	return scaffold, nil
}

func (as *scaffoldService) Delete(id uint) error {
	scaffold, err := as.scaffoldRepository.FindById(id)
	if err != nil {
		return err
	}
	as.scaffoldRepository.Delete(scaffold)
	return nil
}
