package Models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName: "mysql",
			Conn: db,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		return gormDB, mock, err
	}
	return gormDB, mock, err
}

func TestModel_GetAllArticle(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	createTime := time.Now()

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `tag` WHERE `tag`.`article_id` = ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"Name","ArticleID"}).
			AddRow("Google", 1))
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `article`")).
		WillReturnRows(sqlmock.NewRows([]string{"ID","Title","Overview","Text","CreatedAt"}).
			AddRow(1, "Google Title", "Google Overview", "Google Text", createTime))
	mock.ExpectCommit()

	m := Model{Db: db}
	var articles []Article
	err = m.GetAllArticle(&articles)

	// 比べるためのデータ作成処理
	var wants []Article
	wants = append(wants, Article{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", Tags: []Tag{{Name: "Google", ArticleID: 1}}, CreatedAt: createTime})

	//
	assert.Nil(t, err)
	assert.Equal(t, wants, articles)
}

func TestModel_GetArticleByID(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	createTime := time.Now()

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `tag` WHERE `tag`.`article_id` = ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"Name","ArticleID"}).
			AddRow("Google", 1))
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `article` WHERE id = ? ORDER BY `article`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"ID","Title","Overview","Text","CreatedAt"}).
			AddRow(1, "Google Title", "Google Overview", "Google Text", createTime))
	mock.ExpectCommit()

	m := Model{Db: db}
	var article Article
	err = m.GetArticleByID(&article, 1)

	assert.Nil(t, err)
	assert.Equal(t, Article{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", Tags: []Tag{{Name: "Google", ArticleID: 1}}, CreatedAt: createTime}, article)
}

func TestModel_CreateArticle(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	createTime := time.Now()

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `article` (`title`,`overview`,`text`,`created_at`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs( "Google Title", "Google Overview", "Google Text", createTime, 1).
		WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	m := Model{Db: db}
	err = m.CreateArticle(&Article{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", CreatedAt: createTime})

	assert.Nil(t, err)
}

func TestModel_UpdateArticle(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `article`")).
		WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	m := Model{Db: db}
	err = m.UpdateArticle(&Article{ID: 1, Title: "Google Title", Overview: "Google Overview", Text: "Google Text", CreatedAt: time.Now()})

	assert.Nil(t, err)
}

func TestModel_DeleteArticle(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}


	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `article` WHERE id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `tag` WHERE article_id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	m := Model{Db: db}
	var article Article
	err = m.DeleteArticle(&article, 1)

	assert.Nil(t, err)
}

func TestModel_GetAllTag2(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT `name` FROM `tag` GROUP BY `name`")).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("Google").
			AddRow("FaceBook"))
	mock.ExpectCommit()

	m := Model{Db: db}
	var tags []Tag
	err = m.GetAllTag(&tags)

	assert.Nil(t, err)
}

func TestModel_GetArticleByTag(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	createTime := time.Now()

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `tag` WHERE `tag`.`article_id` = ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"name","article_id"}).
			AddRow("Google", 1))
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT `article`.`id`,`article`.`title`,`article`.`overview`,`article`.`text`,`article`.`created_at` FROM `article` inner join tag on article.id = tag.article_id WHERE tag.name = ?")).
		WithArgs("Google").
		WillReturnRows(sqlmock.NewRows([]string{"id","title","overview","text","create"}).
			AddRow(1, "Google Title", "Google Overview", "Google Text", createTime))
	mock.ExpectCommit()

	m := Model{Db: db}
	var articles []Article
	err = m.GetArticleByTag(&articles, "Google")

	assert.Nil(t, err)
}