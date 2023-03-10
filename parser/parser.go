package parser

import (
	"gorm.io/gorm"
	gorm_models "main/internal/datetime"
	"main/internal/session"
	"main/internal/utils/db"
	"main/internal/utils/vars"
	"main/models"
	"time"
)

type Parser struct {
	Session session.Session
	Db      *gorm.DB
	Date    gorm_models.Date
}

func NewParser() Parser {
	now := time.Now()
	return Parser{
		Session: session.NewSession(vars.DefaultUserName, vars.DefaultUserPassword),
		Db:      db.Connect(),
		Date:    gorm_models.NewDateYM(now.Year(), int(now.Month())),
	}
}

func (p *Parser) SaveToDB(model models.DbModel) {
	model.SaveToDB(p.Db)
}

func (p *Parser) DeleteFromDB(model models.DbModel) {
	model.DeleteFromDB(p.Db)
}
