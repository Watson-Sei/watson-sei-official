package Models

import "github.com/Watson-Sei/watson-sei-official/api_v1/Config"
import _ "github.com/jinzhu/gorm/dialects/mysql"

func GetAllArticle(article *[]Article) (err error) {
	if err = Config.DB.Find(article).Error; err != nil {
		return err
	}
	return nil
}