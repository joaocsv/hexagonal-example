package app_test

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/joaocsv/hexagonal-example/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "the price must be zero to have the product disabled", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "hello"
	product.Status = app.ENABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "Invalid"

	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED

	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10

	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
