package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/model"
	"github.com/lfxnxf/emo-server/model/model_posting"
	"github.com/lfxnxf/emo-server/model/model_users"
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

func (d *Dao) SearchPosting(ctx context.Context, where map[string]interface{}, orderBy []model.OrderBy, page, limit int64, withPaging bool) ([]model_posting.SearchPosting, int64, error) {
	log := logging.For(ctx, "func", "SearchPosting",
		zap.Any("where", where),
	)

	var (
		list  []model_posting.SearchPosting
		total int64
	)

	sql := " from `%s` t1 inner join `%s` t2 on t1.uid = t2.id " +
		" left join `%s` t3 on t1.id = t3.posting_id " +
		" left join `%s` t4 on t3.subject_id = t4.id " +
		" left join `%s` t5 on t1.posting_id = t5.business_id and business_type = ? and statistics_type = ? " +
		" left join `%s` t6 on t1.posting_id = t6.business_id and business_type = ? and statistics_type = ? " +
		" left join `%s` t7 on t1.posting_id = t6.business_id and business_type = ? and statistics_type = ? " +
		" where t1.status <> ? and t3.status <> ? and t4.status <> ?"

	searchKeys := make([]interface{}, 0)

	searchKeys = append(searchKeys,
		model_posting.BusinessTypePosting, model_posting.TypeAllThumbNum,
		model_posting.BusinessTypePosting, model_posting.TypeNormalCommentNum,
		model_posting.BusinessTypePosting, model_posting.TypeAllCommentNum,
		model_posting.StatusDeleted, model_posting.StatusDeleted, model_posting.StatusDeleted,
	)

	if id, ok := where["id"]; ok {
		sql += " and t1.id = ?"
		searchKeys = append(searchKeys, id)
	}

	if uid, ok := where["uid"]; ok {
		sql += " and t1.uid = ?"
		searchKeys = append(searchKeys, uid)
	}

	if userType, ok := where["user_type"]; ok {
		sql += " and t1.user_type = ?"
		searchKeys = append(searchKeys, userType)
	}

	if postingType, ok := where["posting_type"]; ok {
		sql += " and t1.posting_type = ?"
		searchKeys = append(searchKeys, postingType)
	}

	if content, ok := where["content"]; ok {
		sql += " and t1.content like ?"
		searchKeys = append(searchKeys, fmt.Sprintf("%%%s%%", content))
	}

	if startAt, ok := where["start_at"]; ok {
		sql += " and create_time >= ?"
		searchKeys = append(searchKeys, startAt)
	}
	if endAt, ok := where["end_at"]; ok {
		sql += " and create_time <= ?"
		searchKeys = append(searchKeys, endAt)
	}

	sql = fmt.Sprintf(
		sql,
		model_posting.TablePosting,
		model_users.TableUsers,
		model_posting.TablePostingSubject,
		model_posting.TableSubject,
		model_posting.TablePostingStatistics,
		model_posting.TablePostingStatistics,
		model_posting.TablePostingStatistics,
	)

	if withPaging {
		// 获取总数
		sqlCount := fmt.Sprintf("select count(DISTINCT t1.id) total %s", sql)
		var totalList []struct {
			Total int64 `gorm:"column:total"`
		}
		err := d.db.Slave().Raw(sqlCount, searchKeys...).Scan(&totalList).Error
		if err != nil {
			log.Errorw("get total error", zap.Error(err))
			return nil, 0, err
		}
		if len(totalList) > 0 {
			total = totalList[0].Total
		}
	}

	sqlList := fmt.Sprintf("select t1.*, t2.nickname, GROUP_CONCAT(DISTINCT t4.name) subjects, "+
		" COALESCE(t5.num, 0) like_num, "+
		" COALESCE(t6.num, 0) human_comment_num "+
		" COALESCE(t7.num, 0) all_comment_num "+
		" %s group by t1.id ", sql)

	if len(orderBy) > 0 {
		sqlList = fmt.Sprintf("%s order by ", sqlList)
		for k, v := range orderBy {
			if k == 0 {
				sqlList = fmt.Sprintf("%s %s %s ", sqlList, v.Key, v.Order)
			} else {
				sqlList = fmt.Sprintf("%s, %s %s ", sqlList, v.Key, v.Order)
			}
		}
	} else {
		sqlList = fmt.Sprintf("%s order by id desc ", sqlList)
	}

	if withPaging {
		sqlList = fmt.Sprintf("%s limit %d, %d", sql, (page-1)*limit, limit)
	}

	err := d.db.Slave().Raw(sqlList, searchKeys...).Scan(&list).Error
	if err != nil {
		log.Errorw("get list error", zap.Error(err))
		return nil, 0, err
	}

	log.Infow("success!", zap.Any("list", list))
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
		"status": model_posting.StatusDeleted,
	}).Error
	if err != nil {
		log.Errorw("find error", zap.Error(err))
		return err
	}

	log.Infow("success!")
	return nil
}

// GetAllSubject 获取全部话题
func (d *Dao) GetAllSubject(ctx context.Context) ([]model_posting.Subject, error) {
	log := logging.For(ctx, "func", "GetAllSubject")
	var list []model_posting.Subject

	err := d.db.Slave().Table(model_posting.TableSubject).Where("status <> ?", model_posting.SubjectStatusDeleted).Find(&list).Error
	if err != nil {
		log.Errorw("find error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return list, nil
}
