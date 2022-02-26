package v1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type ContactORM struct {
	Id   int64  `gorm:"primary_key;unique"`
	Name string `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (ContactORM) TableName() string {
	return "contacts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Contact) ToORM(ctx context.Context) (ContactORM, error) {
	to := ContactORM{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	if posthook, ok := interface{}(m).(ContactWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ContactORM) ToPB(ctx context.Context) (Contact, error) {
	to := Contact{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	if posthook, ok := interface{}(m).(ContactWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Contact the arg will be the target, the caller the one being converted from

// ContactBeforeToORM called before default ToORM code
type ContactWithBeforeToORM interface {
	BeforeToORM(context.Context, *ContactORM) error
}

// ContactAfterToORM called after default ToORM code
type ContactWithAfterToORM interface {
	AfterToORM(context.Context, *ContactORM) error
}

// ContactBeforeToPB called before default ToPB code
type ContactWithBeforeToPB interface {
	BeforeToPB(context.Context, *Contact) error
}

// ContactAfterToPB called after default ToPB code
type ContactWithAfterToPB interface {
	AfterToPB(context.Context, *Contact) error
}

// DefaultCreateContact executes a basic gorm create call
func DefaultCreateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ContactORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
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
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &ContactORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ContactORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ContactORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ContactORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteContact(ctx context.Context, in *Contact, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ContactORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ContactORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteContactSet(ctx context.Context, in []*Contact, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&ContactORM{})).(ContactORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ContactORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ContactORM{})).(ContactORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ContactORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Contact, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Contact, *gorm.DB) error
}

// DefaultStrictUpdateContact clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ContactORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithAfterStrictUpdateSave); ok {
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

type ContactORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchContact executes a basic gorm update call with patch behavior
func DefaultPatchContact(ctx context.Context, in *Contact, updateMask *field_mask.FieldMask, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Contact
	var err error
	if hook, ok := interface{}(&pbObj).(ContactWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadContact(ctx, &Contact{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ContactWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskContact(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ContactWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateContact(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ContactWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ContactWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Contact, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ContactWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Contact, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ContactWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Contact, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ContactWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Contact, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetContact executes a bulk gorm update call with patch behavior
func DefaultPatchSetContact(ctx context.Context, objects []*Contact, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Contact, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Contact, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchContact(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskContact patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskContact(ctx context.Context, patchee *Contact, patcher *Contact, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Contact, error) {
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
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListContact executes a gorm list call
func DefaultListContact(ctx context.Context, db *gorm.DB) ([]*Contact, error) {
	in := Contact{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &ContactORM{}, &Contact{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ContactORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ContactORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Contact{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ContactORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ContactORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ContactORM) error
}