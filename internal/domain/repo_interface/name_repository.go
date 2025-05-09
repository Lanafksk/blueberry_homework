package repo_interface

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/request"
)

// NameRepository는 이름을 관리하는 인터페이스입니다.
type NameRepository interface {
	// CreateName은 새로운 이름을 저장소에 추가합니다.
	CreateName(entity entities.NameEntity) error

	// GetNames는 저장된 모든 이름을 반환합니다.
	GetNames() ([]entities.NameEntity, error)

	// DeleteByName 은 이름을 받아 해당하는 이름을 삭제하고 재정렬합니다
	DeleteByName(name string) error

	// 이름 변경 함수
	ChangeName(req request.ChangeNameRequest) error

	// 내부 로직용 이름 찾기 함수
	FindByName(name string) bool
}
