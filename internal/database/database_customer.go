package database

import (
	"context"

	"github.com/farshidboroomand/goPostgresCrud/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where("email LIKE ?", "%"+emailAddress+"%").
		Find(&customers)
	return customers, result.Error
}
