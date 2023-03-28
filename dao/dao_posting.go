package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/model/model_posting"
	"go.uber.org/zap"
)

// InsertPosting 新增帖子
func (d *Dao) InsertPosting(ctx context.Context, tx *gorm.DB, m model_posting.Posting) (model_posting.Posting, error) {
	log := logging.For(ctx, "func", "InsertPosting",
		zap.Any("m", m),
	)
	err := d.Tx(tx).Create(&m).Error
	if err != nil {
		log.Errorw("d.db.Master().Create(res) error", zap.Error(err))
		return model_posting.Posting{}, err
	}
	return m, nil
}

// EditPosting 修改Posting
func (d *Dao) EditPosting(ctx context.Context, tx *gorm.DB, m model_posting.Posting) error {
	log := logging.For(ctx, "func", "EditPosting",
		zap.Any("m", m),
	)
	tx = d.Tx(tx)
	err := tx.Save(&m).Error
	if err != nil {
		log.Errorw("d.db.Master().Save(&m) error", zap.Error(err))
		return err
	}
	return nil
}

// SearchPosting 获取Posting列表
func (d *Dao) SearchPosting(ctx context.Context, where map[string]interface{}, page, limit int64, withPaging bool) ([]model_posting.Posting, int64, error) {
	log := logging.For(ctx, "func", "SearchPosting",
		zap.Any("where", where),
	)

	var (
		list  []model_posting.Posting
		total int64
	)

	query := d.db.Slave().Table(model_posting.TablePosting).Where("status <> ?", model_posting.PostingStatusDeleted)

	if name, ok := where["name"]; ok {
		query = query.Where("name = ?", name)
	}


	// 获取全部数量
	if withPaging {
		err := query.Count(&total).Error
		if err != nil {
			log.Errorw("d.db.Slave().Count(&total) error", zap.Error(err))
			return list, 0, err
		}

		query = query.Offset((page - 1) * limit).Limit(limit)
	}

	// 获取列表
	err := query.Order("id desc").Find(&list).Error
	if err != nil {
		log.Errorw("d.db.Slave().Find(&res) error", zap.Error(err))
		return list, 0, err
	}

	return list, total, nil
}

// FetchOnePosting 通过id获取数据
func (d *Dao) FetchOnePosting(ctx context.Context, id int64) (model_posting.Posting, error) {
	log := logging.For(ctx, "func", "FetchOnePosting",
		zap.Int64("id", id),
	)
	var list []model_posting.Posting

	err := d.db.Slave().Table(model_posting.TablePosting).Where("id = ?", id).Find(&list).Error
	if err != nil {
		log.Errorw("find error", zap.Error(err))
		return model_posting.Posting{}, err
	}

	if len(list) <= 0 {
		return model_posting.Posting{}, nil
	}

	log.Infow("success!")
	return list[0], nil
}

// MGetPosting 批量获取Posting
func (d *Dao) MGetPosting(ctx context.Context, ids []int64) ([]model_posting.Posting, error) {
	log := logging.For(ctx, "func", "MGetPosting",
		zap.Any("ids", ids),
	)
	var list []model_posting.Posting

	err := d.db.Slave().Table(model_posting.TablePosting).Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		log.Errorw("find error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return list, nil
}

func (d *Dao) DeletePostingByIds(ctx context.Context, ids []int64) error {
	log := logging.For(ctx, "func", "DeletePostingByIds",
		zap.Any("ids", ids),
	)
	err := d.db.Slave().Table(model_posting.TablePosting).Where("id in (?)", ids).Update(map[string]interface{}{
		"status": model_posting.PostingStatusDeleted,
	}).Error
	if err != nil {
		log.Errorw("find error", zap.Error(err))
		return err
	}

	log.Infow("success!")
	return nil
}