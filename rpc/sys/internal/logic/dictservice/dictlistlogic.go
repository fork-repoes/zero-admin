package dictservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictList 字典列表
func (l *DictListLogic) DictList(in *sysclient.DictListReq) (*sysclient.DictListResp, error) {
	q := query.SysDict.WithContext(l.ctx)
	if len(in.Type) > 0 {
		q = q.Where(query.SysDict.Type.Like("%" + in.Type + "%"))
	}
	if len(in.Label) > 0 {
		q = q.Where(query.SysDict.Type.Like("%" + in.Label + "%"))
	}

	if in.DelFlag != 2 {
		q = q.Where(query.SysDict.DelFlag.Eq(in.DelFlag))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询字典列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*sysclient.DictListData
	for _, dict := range result {
		list = append(list, &sysclient.DictListData{
			Id:          dict.ID,
			Value:       dict.Value,
			Label:       dict.Label,
			Type:        dict.Type,
			Description: dict.Description,
			Sort:        dict.Sort,
			Remarks:     *dict.Remarks,
			CreateBy:    dict.CreateBy,
			CreateTime:  dict.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:    *dict.UpdateBy,
			UpdateTime:  dict.UpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:     dict.DelFlag,
		})
	}

	logc.Infof(l.ctx, "查询字典列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.DictListResp{
		Total: count,
		List:  list,
	}, nil

}
