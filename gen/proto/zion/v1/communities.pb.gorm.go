package v1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type CommunityORM struct {
	Created         int64
	Deleted         bool
	Description     string `gorm:"not null"`
	EscrowAmount    int64  `gorm:"not null"`
	Id              int64  `gorm:"primary_key;unique"`
	Img             string
	LastActive      int64
	Name            string `gorm:"not null"`
	OwnerDid        string `gorm:"not null"`
	OwnerUsername   string `gorm:"not null"`
	PricePerMessage int64  `gorm:"not null"`
	PriceToJoin     int64  `gorm:"not null"`
	Public          bool
	Updated         int64
	ZionId          string
}

// TableName overrides the default tablename generated by GORM
func (CommunityORM) TableName() string {
	return "communities"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Community) ToORM(ctx context.Context) (CommunityORM, error) {
	to := CommunityORM{}
	var err error
	if prehook, ok := interface{}(m).(CommunityWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ZionId = m.ZionId
	to.Name = m.Name
	to.OwnerDid = m.OwnerDid
	to.OwnerUsername = m.OwnerUsername
	to.Description = m.Description
	to.Img = m.Img
	// Repeated type string is not an ORMable message type
	to.PriceToJoin = m.PriceToJoin
	to.PricePerMessage = m.PricePerMessage
	to.EscrowAmount = m.EscrowAmount
	to.LastActive = m.LastActive
	to.Public = m.Public
	to.Deleted = m.Deleted
	to.Created = m.Created
	to.Updated = m.Updated
	if posthook, ok := interface{}(m).(CommunityWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CommunityORM) ToPB(ctx context.Context) (Community, error) {
	to := Community{}
	var err error
	if prehook, ok := interface{}(m).(CommunityWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ZionId = m.ZionId
	to.Name = m.Name
	to.OwnerDid = m.OwnerDid
	to.OwnerUsername = m.OwnerUsername
	to.Description = m.Description
	to.Img = m.Img
	// Repeated type string is not an ORMable message type
	to.PriceToJoin = m.PriceToJoin
	to.PricePerMessage = m.PricePerMessage
	to.EscrowAmount = m.EscrowAmount
	to.LastActive = m.LastActive
	to.Public = m.Public
	to.Deleted = m.Deleted
	to.Created = m.Created
	to.Updated = m.Updated
	if posthook, ok := interface{}(m).(CommunityWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Community the arg will be the target, the caller the one being converted from

// CommunityBeforeToORM called before default ToORM code
type CommunityWithBeforeToORM interface {
	BeforeToORM(context.Context, *CommunityORM) error
}

// CommunityAfterToORM called after default ToORM code
type CommunityWithAfterToORM interface {
	AfterToORM(context.Context, *CommunityORM) error
}

// CommunityBeforeToPB called before default ToPB code
type CommunityWithBeforeToPB interface {
	BeforeToPB(context.Context, *Community) error
}

// CommunityAfterToPB called after default ToPB code
type CommunityWithAfterToPB interface {
	AfterToPB(context.Context, *Community) error
}

// DefaultCreateCommunity executes a basic gorm create call
func DefaultCreateCommunity(ctx context.Context, in *Community, db *gorm.DB) (*Community, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CommunityORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadCommunity(ctx context.Context, in *Community, db *gorm.DB) (*Community, error) {
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
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &CommunityORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CommunityORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CommunityORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CommunityORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteCommunity(ctx context.Context, in *Community, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CommunityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CommunityORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteCommunitySet(ctx context.Context, in []*Community, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&CommunityORM{})).(CommunityORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&CommunityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CommunityORM{})).(CommunityORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CommunityORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Community, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Community, *gorm.DB) error
}

// DefaultStrictUpdateCommunity clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateCommunity(ctx context.Context, in *Community, db *gorm.DB) (*Community, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateCommunity")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &CommunityORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithAfterStrictUpdateSave); ok {
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

type CommunityORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchCommunity executes a basic gorm update call with patch behavior
func DefaultPatchCommunity(ctx context.Context, in *Community, updateMask *field_mask.FieldMask, db *gorm.DB) (*Community, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Community
	var err error
	if hook, ok := interface{}(&pbObj).(CommunityWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadCommunity(ctx, &Community{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CommunityWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskCommunity(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CommunityWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateCommunity(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CommunityWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CommunityWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Community, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommunityWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Community, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommunityWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Community, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CommunityWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Community, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetCommunity executes a bulk gorm update call with patch behavior
func DefaultPatchSetCommunity(ctx context.Context, objects []*Community, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Community, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Community, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchCommunity(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskCommunity patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskCommunity(ctx context.Context, patchee *Community, patcher *Community, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Community, error) {
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
		if f == prefix+"ZionId" {
			patchee.ZionId = patcher.ZionId
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"OwnerDid" {
			patchee.OwnerDid = patcher.OwnerDid
			continue
		}
		if f == prefix+"OwnerUsername" {
			patchee.OwnerUsername = patcher.OwnerUsername
			continue
		}
		if f == prefix+"Description" {
			patchee.Description = patcher.Description
			continue
		}
		if f == prefix+"Img" {
			patchee.Img = patcher.Img
			continue
		}
		if f == prefix+"Tags" {
			patchee.Tags = patcher.Tags
			continue
		}
		if f == prefix+"PriceToJoin" {
			patchee.PriceToJoin = patcher.PriceToJoin
			continue
		}
		if f == prefix+"PricePerMessage" {
			patchee.PricePerMessage = patcher.PricePerMessage
			continue
		}
		if f == prefix+"EscrowAmount" {
			patchee.EscrowAmount = patcher.EscrowAmount
			continue
		}
		if f == prefix+"LastActive" {
			patchee.LastActive = patcher.LastActive
			continue
		}
		if f == prefix+"Public" {
			patchee.Public = patcher.Public
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

// DefaultListCommunity executes a gorm list call
func DefaultListCommunity(ctx context.Context, db *gorm.DB) ([]*Community, error) {
	in := Community{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &CommunityORM{}, &Community{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CommunityORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommunityORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Community{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CommunityORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CommunityORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]CommunityORM) error
}
type CommunitiesServiceDefaultServer struct {
	DB *gorm.DB
}

// CreateCommunity ...
func (m *CommunitiesServiceDefaultServer) CreateCommunity(ctx context.Context, in *CreateCommunityRequest) (*CreateCommunityResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithBeforeCreateCommunity); ok {
		var err error
		if db, err = custom.BeforeCreateCommunity(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateCommunity(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateCommunityResponse{Result: res}
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithAfterCreateCommunity); ok {
		var err error
		if err = custom.AfterCreateCommunity(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// CommunitiesServiceCommunityWithBeforeCreateCommunity called before DefaultCreateCommunityCommunity in the default CreateCommunity handler
type CommunitiesServiceCommunityWithBeforeCreateCommunity interface {
	BeforeCreateCommunity(context.Context, *gorm.DB) (*gorm.DB, error)
}

// CommunitiesServiceCommunityWithAfterCreateCommunity called before DefaultCreateCommunityCommunity in the default CreateCommunity handler
type CommunitiesServiceCommunityWithAfterCreateCommunity interface {
	AfterCreateCommunity(context.Context, *CreateCommunityResponse, *gorm.DB) error
}

// DeleteCommunity ...
func (m *CommunitiesServiceDefaultServer) DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest) (*DeleteCommunityResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithBeforeDeleteCommunity); ok {
		var err error
		if db, err = custom.BeforeDeleteCommunity(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteCommunity(ctx, &Community{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteCommunityResponse{}
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithAfterDeleteCommunity); ok {
		var err error
		if err = custom.AfterDeleteCommunity(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// CommunitiesServiceCommunityWithBeforeDeleteCommunity called before DefaultDeleteCommunityCommunity in the default DeleteCommunity handler
type CommunitiesServiceCommunityWithBeforeDeleteCommunity interface {
	BeforeDeleteCommunity(context.Context, *gorm.DB) (*gorm.DB, error)
}

// CommunitiesServiceCommunityWithAfterDeleteCommunity called before DefaultDeleteCommunityCommunity in the default DeleteCommunity handler
type CommunitiesServiceCommunityWithAfterDeleteCommunity interface {
	AfterDeleteCommunity(context.Context, *DeleteCommunityResponse, *gorm.DB) error
}

// EditCommunity ...
func (m *CommunitiesServiceDefaultServer) EditCommunity(ctx context.Context, in *EditCommunityRequest) (*EditCommunityResponse, error) {
	out := &EditCommunityResponse{}
	return out, nil
}

// JoinCommunity ...
func (m *CommunitiesServiceDefaultServer) JoinCommunity(ctx context.Context, in *JoinCommunityRequest) (*JoinCommunityResponse, error) {
	out := &JoinCommunityResponse{}
	return out, nil
}

// ListCommunity ...
func (m *CommunitiesServiceDefaultServer) ListCommunity(ctx context.Context, in *ListCommunityRequest) (*ListCommunityResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithBeforeListCommunity); ok {
		var err error
		if db, err = custom.BeforeListCommunity(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListCommunity(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListCommunityResponse{Results: res}
	if custom, ok := interface{}(in).(CommunitiesServiceCommunityWithAfterListCommunity); ok {
		var err error
		if err = custom.AfterListCommunity(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// CommunitiesServiceCommunityWithBeforeListCommunity called before DefaultListCommunityCommunity in the default ListCommunity handler
type CommunitiesServiceCommunityWithBeforeListCommunity interface {
	BeforeListCommunity(context.Context, *gorm.DB) (*gorm.DB, error)
}

// CommunitiesServiceCommunityWithAfterListCommunity called before DefaultListCommunityCommunity in the default ListCommunity handler
type CommunitiesServiceCommunityWithAfterListCommunity interface {
	AfterListCommunity(context.Context, *ListCommunityResponse, *gorm.DB) error
}
