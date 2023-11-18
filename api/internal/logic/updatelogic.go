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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type M struct {
	Id      int32  `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Version int32  `gorm:"column:version" json:"version"`
}

func (l *UpdateLogic) Update(req *types.Update) (resp *types.State, err error) {
	version, err := utils.ParseLongToken(req.Atoken)
	m := &model.User{
		Name:     sql.NullString{String: req.Name, Valid: true},
		Password: sql.NullInt64{Int64: int64(req.Password), Valid: true},
		Version:  sql.NullInt64{Int64: int64(version.Version) + 1, Valid: true},
	}

	if err != nil {
		errors.New("生成错误")
	}

	db, err := DB.Init()
	if err != nil {
		errors.New("连接数据库失败")
		return
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("关闭数据库失败, error=" + err.Error())
		}
		sqlDB.Close()
	}()

	err = db.Debug().Model(&model.User{}).Table("User").Where("name = ?", m.Name).Updates(&m).Error
	if err != nil {
		errors.New("插入失败")
	}
	tokenUpdate := &utils.JWTClaims{
		Username: m.Name.String,
		Version:  int(m.Version.Int64),
	}

	res, err := utils.GenLongToken(tokenUpdate)
	if err != nil {
		errors.New("生成错误")
	}
	fmt.Println(res)

	return &types.State{Onestring: "success"}, nil
}
