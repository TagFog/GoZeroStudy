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

func (l *UpdateLogic) Update(req *types.Update) (resp *types.State, err error) {

	version, err := utils.ParseShortToken(req.Atoken)
	if err != nil {
		errors.New("生成错误")
	}

	m := &model.User{
		Id:       int64(req.Id),
		Name:     sql.NullString{String: req.Name, Valid: true},
		Password: sql.NullInt64{Int64: int64(req.Password), Valid: true},
		Version:  sql.NullInt64{Int64: int64(version.Version) + 1, Valid: true},
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

	err = db.Debug().Model(&model.User{}).Table("User").Where("name = ? and id = ?", m.Name, m.Id).Updates(&m).Error
	if err != nil {
		errors.New("更新失败")
	}
	//短token生成
	shortTokenUpdate := &utils.JWTClaims{
		Id:       int(m.Id),
		Username: m.Name.String,
		Version:  int(m.Version.Int64),
	}
	res, err := utils.GenShortToken(shortTokenUpdate)
	if err != nil {
		errors.New("生成错误")
	}
	logx.Info(res)

	longTokenUpdate := &utils.JWTClaims{
		Id:       int(m.Id),
		Username: m.Name.String,
	}
	//更新长token,此时应该没有version
	longtoken, err := utils.GenLongToken(longTokenUpdate)
	if err != nil {
		errors.New("生成错误")
	}
	logx.Info(longtoken)

	return &types.State{Onestring: "success"}, nil
}
