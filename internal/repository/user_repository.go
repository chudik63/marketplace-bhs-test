package repository

import "marketplace-bhs-test/internal/entity"

type UserRepository interface {
	Create(*entity.User) error
}
