package repository

type itemShopRepositoryImpl struct{}

func NewitemShopRepositoryImpl() ItemShopRepository {
	return &itemShopRepositoryImpl{}
}
