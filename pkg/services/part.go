package services

import (
	"time"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

// Part
type Part struct{}

// Create new part
func (ps *Part) Create(ctx context.Context, req *CreateRequest) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	if err := ps.create(ctx, req); err != nil {
		return failResponse(), err
	}

	return successResponse(), nil
}

func (ps *Part) create(ctx context.Context, r *CreateRequest) error {
	db := getConnection()
	if _, err := db.QueryContext(ctx, "insert into part(manufacturer_id, vendor_code) values (?, ?)",
		r.ManufacturerID, r.VendorCode); err != nil {
		return errors.Wrap(err, "cannot create part")
	}

	return nil
}

// Update part by id
func (ps *Part) Update(ctx context.Context, req *UpdateRequest) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	if err := ps.update(ctx, req); err != nil {
		return failResponse(), err
	}

	return successResponse(), nil
}

func (ps *Part) update(ctx context.Context, r *UpdateRequest) error {
	db := getConnection()
	if _, err := db.QueryContext(ctx,
		"update part set manufacturer_id = ?, vendor_code = ? where id = ?",
		r.ID, r.ManufacturerID, r.VendorCode); err != nil {
		return errors.Wrap(err, "cannot create part")
	}

	return nil
}

// Delete part by id
func (ps *Part) Delete(ctx context.Context, req *DeleteRequest) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	if err := ps.delete(ctx, req); err != nil {
		return failResponse(), err
	}

	return successResponse(), nil
}

func (ps *Part) delete(ctx context.Context, r *DeleteRequest) error {
	db := getConnection()
	if _, err := db.QueryContext(ctx,
		"update part set deleted_at = now() where id = ?",
		r.ID); err != nil {
		return errors.Wrap(err, "cannot create part")
	}

	return nil
}

// Get part by id
func (ps *Part) Get(ctx context.Context, req *GetRequest) (*SingleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	return ps.get(ctx, req)
}

func (ps *Part) get(ctx context.Context, r *GetRequest) (*SingleResponse, error) {
	res := &SingleResponse{}
	db := getConnection()
	err := db.QueryRowContext(ctx,
		"select id, manufacturer_id, vendor_code, created_at from part where id = ? and deleted_at is null",
		r.ID).Scan(res.Body.ID, res.Body.ManufacturerID, res.Body.VendorCode, res.Body.CreatedAt)
	if err != nil {
		res.Response = failResponse()
		return res, err
	}

	res.Response = successResponse()
	return res, nil
}

func (ps *Part) BatchCreate(ctx context.Context, req *BatchCreateRequest) (*Response, error) {
	return nil, nil
}

func (ps *Part) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	return ps.list(ctx, req)
}

func (ps *Part) list(ctx context.Context, r *ListRequest) (*ListResponse, error) {
	page := int32(1)
	if r.Page > 1 {
		page = r.Page
	}

	if r.PageSize < 1 {
		return nil, errors.New("page size should not be less than 1")
	}

	offset := page * r.PageSize
	limit := r.PageSize

	db := getConnection()

	var total int64
	if err := db.QueryRowContext(ctx, "select count(*) from part where deleted_at is not null offset ? limit ?", offset, limit).Scan(&total); err != nil {
		return nil, errors.Wrap(err, "cannot count of parts")
	}

	var prl []*PartResponse
	rows, err := db.QueryContext(ctx, "select id, manufacturer_id, vendor_code, created_at from part where deleted_at is not null offset ? limit ?", offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get list of parts")
	}

	for rows.Next() {
		pr := PartResponse{}
		if err := rows.Scan(&pr.ID, &pr.ManufacturerID, &pr.VendorCode, &pr.CreatedAt); err != nil {
			return nil, errors.Wrap(err, "cannot scant ")
		}

		prl = append(prl, &pr)
	}

	res := &ListResponse{
		Body:      prl,
		PageCount: 100,
		Response:  successResponse(),
	}

	return res, nil
}
