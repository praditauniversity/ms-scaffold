package repository

import (
	"errors"
	"ms-scaffold/api/models/domain"

	"gorm.io/gorm"
)

type ScaffoldRepository interface {
	All() []domain.Scaffold
	Create(scaffold domain.Scaffold) domain.Scaffold
	Update(scaffold domain.Scaffold) domain.Scaffold
	Delete(scaffold domain.Scaffold)
	FindById(id uint) (domain.Scaffold, error)
}

type ScaffoldConnectDB struct {
	dbConnect *gorm.DB //connect to database
}

func NewScaffoldRepository(db *gorm.DB) ScaffoldRepository {
	return &ScaffoldConnectDB{dbConnect: db} //connect database to interface
}

func (conn *ScaffoldConnectDB) All() []domain.Scaffold {
	var scaffold []domain.Scaffold
	conn.dbConnect.Find(&scaffold)
	return scaffold
}

func (conn *ScaffoldConnectDB) Create(scaffold domain.Scaffold) domain.Scaffold {
	conn.dbConnect.Save(&scaffold)
	return scaffold
}

func (conn *ScaffoldConnectDB) Update(scaffold domain.Scaffold) domain.Scaffold {
	conn.dbConnect.Omit("created_at").Save(&scaffold)
	return scaffold
}

func (conn *ScaffoldConnectDB) Delete(scaffold domain.Scaffold) {
	conn.dbConnect.Delete(&scaffold)
}

func (conn *ScaffoldConnectDB) FindById(id uint) (domain.Scaffold, error) {
	var scaffold domain.Scaffold
	conn.dbConnect.Find(&scaffold, "id = ?", id)
	if scaffold.ID == 0 {
		return scaffold, errors.New("id not found")
	}
	return scaffold, nil
}
