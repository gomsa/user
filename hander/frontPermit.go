package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user/proto/frontPermit"
	"github.com/gomsa/user/service"
)

// FrontPermit 前端权限结构
type FrontPermit struct {
	Repo service.FPRepository
}

// UpdateOrCreate 创建或者更新
func (srv *FrontPermit) UpdateOrCreate(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	p := pb.FrontPermit{}
	p.App = req.App
	p.Service = req.Service
	p.Method = req.Method
	frontPermit, err := srv.Repo.Get(&p)
	if frontPermit == nil {
		_, err = srv.Repo.Create(req)
	} else {
		req.Id = frontPermit.Id
		_, err = srv.Repo.Update(req)
	}
	return err
}

// All 获取所有前端权限
func (srv *FrontPermit) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	frontPermits, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.FrontPermits = frontPermits
	return err
}

// List 获取所有前端权限
func (srv *FrontPermit) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	frontPermits, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.FrontPermits = frontPermits
	res.Total = total
	return err
}

// Get 获取前端权限
func (srv *FrontPermit) Get(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	frontPermit, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.FrontPermit = frontPermit
	return err
}

// Create 创建前端权限
func (srv *FrontPermit) Create(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加前端权限失败")
	}
	res.Valid = true
	return err
}

// Update 更新前端权限
func (srv *FrontPermit) Update(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新前端权限失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除前端权限
func (srv *FrontPermit) Delete(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除前端权限失败")
	}
	res.Valid = valid
	return err
}