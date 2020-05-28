package dao

import (
	"backend/app/server/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

func (d *dao) Info(address string) (res *model.Info, err error) {
	res = &model.Info{}

	// DB
	DBres := d.db.Where(&model.Info{Address: address}).First(res)
	err = DBres.Error
	if DBres.RecordNotFound() {
		err = ecode.Errorf(codes.NotFound, "Can not found server (%s)", address)
	}
	return
}

func (d *dao) UpdateRcon(server *model.Info) (err error) {
	// DB
	err = d.db.Model(server).Update("rcon", server.Rcon).Error
	return
}
