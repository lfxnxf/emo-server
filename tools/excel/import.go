package excel

import (
	"context"
	"fmt"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

func ReadFromExcel(ctx context.Context, file, sheet string, keys []string, hasTitle bool) ([]map[string]interface{}, error) {
	log := logging.For(ctx, "func", "ReadFromExcel",
		zap.String("file", file),
		zap.String("sheet", sheet),
		zap.Any("keys", keys),
		zap.Bool("has_title", hasTitle),
	)

	resp := make([]map[string]interface{}, 0)

	// 打开文件
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		log.Errorw("excelize.OpenFile error", zap.Error(err))
		return resp, err
	}

	if len(sheet) <= 0 {
		sheet = fmt.Sprintf("%s%d", "Sheet", 1)
	}

	// 获取excel中具体的列的值
	rows, err := xlsx.GetRows(sheet)
	if err != nil {
		log.Errorw("xlsx.GetRows error", zap.Error(err))
		return nil, err
	}
	// 循环刚刚获取到的表中的值
	for k, row := range rows {
		if hasTitle && k == 0 {
			continue
		}
		info := make(map[string]interface{})
		for key, v := range keys {
			if len(row) <= key {
				continue
			}
			info[v] = row[key]
		}
		resp = append(resp, info)
	}

	log.Infow("success!", zap.Any("resp", resp))

	return resp, nil
}
