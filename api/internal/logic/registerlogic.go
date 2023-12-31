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

func (l *RegisterLogic) Register(req *types.Register) (resp *types.Token, err error) {
	db, err := DB.Init()
	if err != nil {
		errors.New("连接数据库失败")
		return
	}
	m := &model.User{
		Name:     sql.NullString{String: req.Name, Valid: true},
		Password: sql.NullInt64{Int64: int64(req.Password), Valid: true},
		Version:  sql.NullInt64{Int64: 0, Valid: true},
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
	//这里要加一个唯一id
	token := &utils.JWTClaims{
		Id:       int(m.Id),
		Username: m.Name.String,
		Version:  int(m.Version.Int64),
	}
	shortToken, err := utils.GenShortToken(token)
	if err != nil {
		errors.New("生成错误")
	}
	logx.Info(shortToken)
	//生成长token
	longToken, err := utils.GenLongToken(token)
	if err != nil {
		errors.New("生成错误")
	}
	logx.Info(longToken)

	return &types.Token{Atoken: shortToken, Rtoken: longToken}, nil
}
