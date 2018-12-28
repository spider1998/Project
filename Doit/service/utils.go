package service

import (
	"Project/Doit/app"
	"Project/Doit/code"
	"Project/Doit/util"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

// 数据库错误处理
func DbErrorHandler(err error, allowNotFound bool) error {
	if err != nil {
		if util.IsDBNotFound(err) {
			if !allowNotFound {

				err = code.New(http.StatusNotFound, code.CodeRecordNotExist).Err(err)
			} else {
				//允许不存在数据库记录
				err = nil
			}
		} else {
			err = errors.WithStack(err)
		}
	}
	return err
}

func PretreatmentQueryPage(c *routing.Context, query *dbx.SelectQuery) (err error) {
	//记录总数
	var cnt int
	err = query.Row(&cnt)
	if err = DbErrorHandler(err, true); err != nil {
		return
	}

	c.Response.Header().Set("X-Total-Count", strconv.Itoa(cnt))
	//解析页数和每页数量
	page, pageSize, err := util.ParsePagination(c)
	if err != nil {
		err = errors.Wrap(err, "fail to parse pagination.")
		return
	}
	query.Limit(pageSize).Offset((page - 1) * pageSize).OrderBy("create_time desc")
	return
}

//获取一个 id-xxx 的图，根据：源slice，提取字段，查询表名，返回map包含(查询)的字段
func GetIdXxxMap(source interface{}, field string, tableName string, selectField string, mapFields ...string) (maps map[string]dbx.NullStringMap, err error) {
	conditionList := util.SliceAnyToSliceInterface(util.Map(source, field))
	mapFields = append(mapFields, selectField)
	rows, err := app.DB.Select(mapFields...).From(tableName).Where(dbx.In(selectField, conditionList...)).Rows()
	if err = DbErrorHandler(err, true); err != nil {
		return
	}
	maps = make(map[string]dbx.NullStringMap, 0)
	for rows.Next() {
		temp := dbx.NullStringMap{}
		rows.ScanMap(temp)
		maps[temp["id"].String] = temp
	}
	return
}

/*//字符串分割为图片列表
func GetPhotoList(url string) (list []entity.BasePhoto) {
	list = make([]entity.BasePhoto, 0)
	urls := strings.Split(url, entity.PhotoSeparator)
	for k := range urls {
		var temp entity.BasePhoto
		temp.Url = urls[k]
		if temp.Url != "" {
			list = append(list, temp)
		}
	}
	return list
}

//图片数组组合为字符串
func GetPhotoString(list []entity.BasePhoto) (s string) {
	//避免图片为空查询报错的情况
	s = entity.PhotoSeparator
	for k := range list {
		s += list[k].Url + entity.PhotoSeparator
	}
	return s
}
*/
