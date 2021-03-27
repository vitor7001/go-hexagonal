package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitor7001/go-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t,"The price must be greather than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T){

	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t,"The price must be zero in order to have the product disabled", err.Error())
}