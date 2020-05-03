package services

import (
	"time"

	"github.com/pkg/errors"

	"golang.org/x/net/context"
)

// Manufacturer of the parts
type Manufacturer struct {
	tableName struct{} `pg:"manufacturer"`
	ID        int32
	Name      string
}

// Part of the automobiles
type Part struct {
	tableName      struct{} `pg:"part"`
	ID             int64
	ManufacturerID int32
	Manufacturer   *Manufacturer
	VendorCode     string
	CreatedAt      *time.Time
}

// Create new part
func (ps *Part) Create(ctx context.Context, r *CreateRequest) (*Response, error) {
	db := getConnection()
	ps.VendorCode = r.VendorCode
	ps.ManufacturerID = r.ManufacturerID

	// insert into part(manufacturer_id, vendor_code) values (?, ?)
	if err := db.Insert(ps); err != nil {
		return nil, errors.Wrap(err, "cannot create part")
	}

	return &Response{Success: true, Message: "part created successfully"}, nil
}

// Update part by id
func (ps *Part) Update(ctx context.Context, r *UpdateRequest) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	db := getConnection()
	// update part set manufacturer_id = ?, vendor_code = ? where id = ? and deleted_at is null
	if _, err := db.Model(ps).Context(ctx).Where("id = ? and deleted_at is null", r.ID).
		Set("manufacturer_id = ?, vendor_code = ?", r.ManufacturerID, r.VendorCode).Update(); err != nil {
		return nil, errors.Wrap(err, "cannot create part")
	}

	return &Response{
		Success: true,
		Message: "part was updated",
	}, nil
}

// Delete part by id
func (ps *Part) Delete(ctx context.Context, r *DeleteRequest) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	db := getConnection()
	// "update part set deleted_at = now() where id = ?",
	if _, err := db.Model(ps).Context(ctx).Set("deleted_at = now()").Where("id = ?", r.ID).Update(); err != nil {
		return nil, errors.Wrap(err, "cannot create part")
	}

	return &Response{
		Success: true,
		Message: "deleted successfully",
	}, nil
}

// Get part by id
func (ps *Part) Get(ctx context.Context, r *GetRequest) (*SingleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	db := getConnection()

	// select p.id, m.name, p.vendor_code, p.created_at from part p join manufacturer m on m.id = p.manufacturer_id where p.id = ? and p.deleted_at is null;
	if err := db.Model(ps).Context(ctx).Relation("Manufacturer").Where("id = ? and deleted_at is null", r.ID).Select(); err != nil {
		return nil, errors.Wrap(err, "cannot get part")
	}

	return &SingleResponse{
		Success: true,
		Body: &PartResponse{
			ID:           ps.ID,
			Manufacturer: ps.Manufacturer.Name,
			VendorCode:   ps.VendorCode,
			CreatedAt:    ps.CreatedAt.Unix(),
		},
	}, nil
}

// List of the part with paging
func (ps *Part) List(ctx context.Context, r *ListRequest) (*ListResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	if r.Page < 1 {
		return nil, errors.New("page should not be less than 1")
	}

	var page int32
	if r.Page > 0 {
		page = r.Page - 1
	}

	if r.PageSize < 1 {
		return nil, errors.New("page size should not be less than 1")
	}
	offset := int(page * r.PageSize)
	limit := int(r.PageSize)

	db := getConnection()

	// select count(*) from part where deleted_at is null offset ? limit ?
	total, err := db.Model(ps).Context(ctx).Where("deleted_at is null").Limit(limit).Offset(offset).Count()
	if err != nil {
		return nil, errors.Wrap(err, "cannot count of parts")
	}

	var parts []Part
	// select p.id, m.name, p.vendor_code, p.created_at from part p join manufacturer m on m.id = p.manufacturer_id where p.deleted_at is null offset ? limit ?
	err = db.Model(&parts).Relation("Manufacturer").Where("deleted_at is null").Offset(offset).Limit(limit).Select()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get list of parts")
	}

	var prr []*PartResponse
	for _, p := range parts {
		pr := PartResponse{
			ID:           p.ID,
			Manufacturer: p.Manufacturer.Name,
			VendorCode:   p.VendorCode,
			CreatedAt:    p.CreatedAt.Unix(),
		}

		prr = append(prr, &pr)
	}

	pageCount := int32(1)
	if total > 0 && r.PageSize > 0 {
		pageCount = int32(total) / r.PageSize
	}

	return &ListResponse{
		Body:      prr,
		PageCount: pageCount,
		Success:   true,
	}, nil
}

// BatchCreate parts
func (ps *Part) BatchCreate(ctx context.Context, r *BatchCreateRequest) (*Response, error) {
	return nil, nil
}
