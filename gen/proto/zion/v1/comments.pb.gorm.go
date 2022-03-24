package v1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type CommentORM struct {
	Created int64
	Deleted bool
	Id      int64 `gorm:"primary_key;unique"`
	Text    string
	Updated int64
	Zid     string
}

// TableName overrides the default tablename generated by GORM
func (CommentORM) TableName() string {
	return "comments"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Comment) ToORM(ctx context.Context) (CommentORM, error) {
	to := CommentORM{}
	var err error
	if prehook, ok := interface{}(m).(CommentWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Zid = m.Zid
	to.Text = m.Text
	to.Deleted = m.Deleted
	to.Created = m.Created
	to.Updated = m.Updated
	if posthook, ok := interface{}(m).(CommentWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CommentORM) ToPB(ctx context.Context) (Comment, error) {
	to := Comment{}
	var err error
	if prehook, ok := interface{}(m).(CommentWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Zid = m.Zid
	to.Text = m.Text
	to.Deleted = m.Deleted
	to.Created = m.Created
	to.Updated = m.Updated
	if posthook, ok := interface{}(m).(CommentWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Comment the arg will be the target, the caller the one being converted from

// CommentBeforeToORM called before default ToORM code
type CommentWithBeforeToORM interface {
	BeforeToORM(context.Context, *CommentORM) error
}

// CommentAfterToORM called after default ToORM code
type CommentWithAfterToORM interface {
	AfterToORM(context.Context, *CommentORM) error
}

// CommentBeforeToPB called before default ToPB code
type CommentWithBeforeToPB interface {
	BeforeToPB(context.Context, *Comment) error
}

// CommentAfterToPB called after default ToPB code
type CommentWithAfterToPB interface {
	AfterToPB(context.Context, *Comment) error
}

// DefaultCreateComment executes a basic gorm create call
func DefaultCreateComment(ctx context.Context, in *Comment, db *gorm.DB) (*Comment, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CommentORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadComment(ctx context.Context, in *Comment, db *gorm.DB) (*Comment, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &CommentORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CommentORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CommentORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CommentORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteComment(ctx context.Context, in *Comment, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CommentORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CommentORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteCommentSet(ctx context.Context, in []*Comment, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&CommentORM{})).(CommentORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&CommentORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CommentORM{})).(CommentORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CommentORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Comment, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Comment, *gorm.DB) error
}

// DefaultStrictUpdateComment clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateComment(ctx context.Context, in *Comment, db *gorm.DB) (*Comment, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateComment")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &CommentORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type CommentORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchComment executes a basic gorm update call with patch behavior
func DefaultPatchComment(ctx context.Context, in *Comment, updateMask *field_mask.FieldMask, db *gorm.DB) (*Comment, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Comment
	var err error
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadComment(ctx, &Comment{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskComment(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateComment(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CommentWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CommentWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Comment, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommentWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Comment, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommentWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Comment, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommentWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Comment, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetComment executes a bulk gorm update call with patch behavior
func DefaultPatchSetComment(ctx context.Context, objects []*Comment, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Comment, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Comment, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchComment(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskComment patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskComment(ctx context.Context, patchee *Comment, patcher *Comment, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Comment, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Zid" {
			patchee.Zid = patcher.Zid
			continue
		}
		if f == prefix+"Text" {
			patchee.Text = patcher.Text
			continue
		}
		if f == prefix+"Deleted" {
			patchee.Deleted = patcher.Deleted
			continue
		}
		if f == prefix+"Created" {
			patchee.Created = patcher.Created
			continue
		}
		if f == prefix+"Updated" {
			patchee.Updated = patcher.Updated
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListComment executes a gorm list call
func DefaultListComment(ctx context.Context, db *gorm.DB) ([]*Comment, error) {
	in := Comment{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &CommentORM{}, &Comment{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CommentORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Comment{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CommentORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommentORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]CommentORM) error
}