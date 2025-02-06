package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	randomalias "url-short/internal/lib/randomAlias"
	"url-short/internal/service/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T){
	ctx := context.Background()
	
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	t.Run("success with provider alias", func(t *testing.T) {
		mockRepo := new(mocks.UrlRepository)
		service := urlService{
			repo: mockRepo,
			log: logger,
		}
		url := "http://example.com"
		alias := "puska"


		t.Cleanup(func() {
            mockRepo.AssertExpectations(t)
        })

		mockRepo.On("Create", ctx, url, alias).Return(nil)


		result, err := service.Create(ctx, url, alias)
		assert.NoError(t, err)
		assert.Equal(t, alias, result)
		
	})

	t.Run("success without alias", func(t *testing.T) {
		mockRepo := new(mocks.UrlRepository)
		service := urlService{
			repo: mockRepo,
			log: logger,
		}
		url := "http://example.com"
		extendedAlias := "puska"


		t.Cleanup(func() {
            mockRepo.AssertExpectations(t)
        })

		randomalias.RandomAlias = func(length int) string {
            return extendedAlias
        }

		mockRepo.On("Create", ctx, url, extendedAlias).Return(nil)



		result, err := service.Create(ctx, url, "")
		assert.NoError(t, err)
		assert.Equal(t, extendedAlias, result)
		
	})

	t.Run("db return error", func(t *testing.T) {
		mockRepo := new(mocks.UrlRepository)
		service := urlService{
			repo: mockRepo,
			log: logger,
		}
		url := "http://example.com"
		alias := "alias"
		expectedError := errors.New("database error")


		t.Cleanup(func() {
            mockRepo.AssertExpectations(t)
        })
	
		mockRepo.On("Create", ctx, url, alias).Return(expectedError)

		result, err := service.Create(ctx, url, alias)

		fmt.Println("result:", result)
		fmt.Println("err:", err)

		assert.Error(t, err)
		assert.Equal(t, "", result) 
		assert.Equal(t, expectedError, err) 
	
		
	})
	
}


func TestGetUrl(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	t.Run("success get url", func(t *testing.T) {
		mockRepo := new(mocks.UrlRepository)
		alias := "puska"
		extentedUrl := "www.google.com"
		service := urlService{
			repo: mockRepo,
			log: logger,
		}

		mockRepo.On("GetByAlias", ctx, alias).Return(extentedUrl, nil)

		url, err := service.GetUrl(ctx, alias)
		assert.NoError(t, err)
		assert.Equal(t, url, extentedUrl)
	})

	t.Run("get url with error", func(t *testing.T) {
		mockRepo := new(mocks.UrlRepository)
		alias := "puska"
		extendedErr := errors.New("database error")
		service := urlService{
			repo: mockRepo,
			log: logger,
		}

		mockRepo.On("GetByAlias", ctx, alias).Return("", extendedErr)

		url, err := service.GetUrl(ctx, alias)
		assert.Error(t, err)
		assert.Equal(t, "", url)
		assert.Equal(t, extendedErr, err)
	})
}