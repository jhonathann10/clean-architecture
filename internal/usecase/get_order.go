package usecase

import "github.com/jhonathann10/clean-architecture/internal/entity"

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.Get()
	if err != nil {
		return nil, err
	}

	var orderOutputDTO []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orderOutputDTO = append(orderOutputDTO, dto)
	}

	return orderOutputDTO, nil
}
