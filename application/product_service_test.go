package application_test

import (
	mock_application "github.com/vitor7001/go-hexagonal/application/mocks"
	"github.com/vitor7001/go-hexagonal/application"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require")

func TestProductService_Get(t *testing.T){

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")
	require.Nil(t, err)

	require.Equal(t, product , result)

}