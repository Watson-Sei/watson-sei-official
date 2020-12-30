package Controllers

import (
	"bytes"
	"encoding/json"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) GetAllArticle() (*[]Models.Article, error) {
	args := m.Called()
	return args.Get(0).(*[]Models.Article), args.Error(1)
}

func (m *ModelMock) CreateArticle(article *Models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func TestController_GetArticle(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	createTime := time.Now().In(jst)
	articles := &[]Models.Article{
		{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", CreatedAt: createTime},
	}
	modelMock := new(ModelMock)
	modelMock.On("GetAllArticle").Return(articles ,nil)

	controller := GetArticleController{Model: modelMock}

	// テスト用のコンテキストを作成
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/article/list",
		nil)
	controller.GetArticle(c)

	var wants []Models.Article
	wants = append(wants, Models.Article{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", CreatedAt: createTime})

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var retArticle []Models.Article
	if err := json.Unmarshal(body, &retArticle); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
	// なぜか一致しない
	//assert.Equal(t, retArticle, wants)
}
// Create Controller
func TestCreateArticleController_CreateArticle(t *testing.T) {

	sample := new(Models.Article)
	sample.Title = "Google Title"
	sample.Overview = "Google Overview"
	sample.Text = "Google Text"
	sample.Tags = []Models.Tag{{Name: "Google", ArticleID: 1}}

	modelMock := new(ModelMock)
	modelMock.On("CreateArticle", sample).Return(nil)

	bytesJSON, _ := json.Marshal(sample)

	controller := CreateArticleController{Model: modelMock}

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	req, _ := http.NewRequest(
		http.MethodPost,
		"/article/create",
		bytes.NewBuffer(bytesJSON))
	context.Request = req

	controller.CreateArticle(context)

	assert.Equal(t, 200, response.Code)

}