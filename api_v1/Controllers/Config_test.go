package Controllers

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
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

	controller := GetArticleController{Model: modelMock}

	// テスト用のコンテキストを作成
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	controller.GetArticle(context)

	// Model.GetAllArticle が呼ばれているかどうかをチェックできると良い

	// 今回は、正常に article が取れたパターン
	// Status Code が 200
	assert.Equal(t, recorder.Code, http.StatusOK)

	// レスポンスボディが正しい
	wantJSON := "{\"a\":\"b\",...}"
	assert.JSONEq(t, recorder.Body.String(), wantJSON)
}

// error の方もテストしたほうが良い
.Retun(error) して
assert.Equal(t, recoder.Code, http.StatusNotFound)
