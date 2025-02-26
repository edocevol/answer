package unique

import (
	"context"
	"fmt"
	"strconv"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// uniqueIDRepo Unique id repository
type uniqueIDRepo struct {
	data *data.Data
}

// NewUniqueIDRepo new repository
func NewUniqueIDRepo(data *data.Data) unique.UniqueIDRepo {
	return &uniqueIDRepo{
		data: data,
	}
}

// GenUniqueID generate unique id
// 1 + 00x(objectType) + 000000000000x(id)
func (ur *uniqueIDRepo) GenUniqueID(ctx context.Context, key string) (uniqueID int64, err error) {
	idStr, err := ur.GenUniqueIDStr(ctx, key)
	if err != nil {
		return 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	uniqueID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return uniqueID, nil
}

// GenUniqueIDStr generate unique id string
func (ur *uniqueIDRepo) GenUniqueIDStr(ctx context.Context, key string) (uniqueID string, err error) {
	objectType := constant.ObjectTypeStrMapping[key]
	bean := &entity.Uniqid{UniqidType: objectType}
	_, err = ur.data.DB.Insert(bean)
	if err != nil {
		return "", errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return fmt.Sprintf("1%03d%013d", objectType, bean.ID), nil
}
