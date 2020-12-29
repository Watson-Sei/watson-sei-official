package Controllers

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) GetAllArticle(article *[]Models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func TestController_GetArticle(t *testing.T) {
	var article []Models.Article
	modelMock := new(ModelMock)
	modelMock.On("GetAllArticle", article).Return(nil)

	controller := Controller{Model: modelMock}
	err := controller.GetAllArticle(&article)

	assert.Nil(t, err)
}