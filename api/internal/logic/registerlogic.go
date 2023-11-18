package logic

import (
	"GoZeroStudy/api/internal/logic/DB"
	"GoZeroStudy/api/internal/logic/utils"
	"GoZeroStudy/api/internal/model"
	"GoZeroStudy/api/internal/svc"
	"GoZeroStudy/api/internal/types"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//type M struct {
//	Id      int32  `gorm:"column:id" json:"id"`
//	Name    string `gorm:"column:name" json:"name"`
//	Version int32  `gorm:"column:version" json:"version"`
//}

func (l *RegisterLogic) Register(req *types.Register) (resp *types.Token, err error) {
	db, err := DB.Init()
	//data := new(m)
	if err != nil {
		errors.New("连接数据库失败")
		return
	}
	m := &model.User{
		Name: sql.NullString{String: req.Name, Valid: true},
		Password: sql.NullInt64{
			Int64: int64(req.Password),
			Valid: true,
		},
		Version: sql.NullInt64{Int64: 0, Valid: true},
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("关闭数据库失败, error=" + err.Error())
		}
		sqlDB.Close()
	}()
	err = db.Debug().Model(&model.User{}).Table("User").Create(&m).Error
	if err != nil {
		errors.New("插入失败")
	}
	token := &utils.JWTClaims{
		UserID:   int(m.Id),
		Username: m.Name.String,
		Version:  int(m.Version.Int64),
	}
	res, err := utils.GenLongToken(token)
	if err != nil {
		errors.New("生成错误")
	}

	fmt.Println(res)
	data, err := utils.ParseLongToken(res)
	if err != nil {
		errors.New("生成错误")
	}
	fmt.Println(data.Version)
	return
}
