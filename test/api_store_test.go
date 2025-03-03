/*
PetStore API

Testing StoreAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package testsdkgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_testsdkgo_StoreAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StoreAPIService GetOrderById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId int64

		resp, httpRes, err := apiClient.StoreAPI.GetOrderById(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreAPIService PlaceOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StoreAPI.PlaceOrder(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
