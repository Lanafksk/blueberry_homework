package repo_interface

import (
	"blueberry_homework/internal/domain/entities"
)

type CompanyRepository interface {
	CreateCompany(company entities.CompanyEntity) error
	GetCompanies() ([]entities.CompanyEntity, error)
}
