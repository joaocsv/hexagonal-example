package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joaocsv/hexagonal-example/app"
	mock_app "github.com/joaocsv/hexagonal-example/app/mock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	product := mock_app.NewMockProductInterface(crtl)
	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	product := mock_app.NewMockProductInterface(crtl)
	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 2.22)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	product := mock_app.NewMockProductInterface(crtl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()

	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
