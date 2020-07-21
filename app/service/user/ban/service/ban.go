package service

import (
	serverService "backend/app/game/server/api/grpc"
	pb "backend/app/service/user/ban/api/grpc"
	"backend/app/service/user/ban/model"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"time"
)

const (
	BanNotifyCMD = "ncs_ban_notify %d %d \"%s\""
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	res, err := s.dao.Info(req.Uid)
	if res != nil {
		resp.Info = &pb.Info{
			Id:         res.ID,
			Uid:        res.UID,
			CreateTime: res.CreateTime,
			ExpireTime: res.ExpireTime,
			Type:       int32(res.Type),
			ServerId:   res.ServerID,
			ModId:      res.ModID,
			GameId:     res.GameID,
			Reason:     res.Reason,
		}
	}
	return
}

func (s *Service) Add(ctx context.Context, req *pb.AddReq) (resp *pb.AddResp, err error) {
	resp = &pb.AddResp{}
	now := time.Now().UTC()

	if req.Info == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info")
		return
	}
	if req.Info.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info.UID")
		return
	}
	if req.Info.Type == model.BanTypeNone {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info.Type")
		return
	}
	if req.Info.ExpireTime <= now.Unix() {
		err = ecode.Errorf(codes.InvalidArgument, "Expire time must be after %s", now.Format(time.UnixDate))
		return
	}

	err = s.dao.Add(&model.Ban{
		UID:        req.Info.Uid,
		CreateTime: now.Unix(),
		ExpireTime: req.Info.ExpireTime,
		Type:       int(req.Info.Type),
		ServerID:   req.Info.ServerId,
		ModID:      req.Info.ModId,
		GameID:     req.Info.GameId,
		Reason:     req.Info.Reason,
	})

	go func() {
		res, err := s.server.RconAll(context.Background(), &serverService.RconAllReq{
			Cmd: fmt.Sprintf(BanNotifyCMD, req.Info.Uid, req.Info.Type, req.Info.Reason),
		})
		if res.Success == 0 {
			log.Warn("None server notify!")
		}
		if err != nil {
			log.Error(err)
		}
	}()
	return
}

func (s *Service) Remove(ctx context.Context, req *pb.RemoveReq) (resp *pb.RemoveResp, err error) {
	resp = &pb.RemoveResp{}

	err = s.dao.Remove(&model.Ban{
		ID: req.Id,
	})
	return
}

func (s *Service) BanCheck(ctx context.Context, req *pb.Info2Req) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	res, err := s.dao.Info(req.Uid)
	if res != nil {
		if res.IsBanned(req.ServerId, req.ModId, req.GameId) {
			resp.Info = &pb.Info{
				Id:         res.ID,
				Uid:        res.UID,
				CreateTime: res.CreateTime,
				ExpireTime: res.ExpireTime,
				Type:       int32(res.Type),
				ServerId:   res.ServerID,
				ModId:      res.ModID,
				GameId:     res.GameID,
				Reason:     res.Reason,
			}
			return
		}
	}
	err = ecode.Errorf(codes.NotFound, "No ban recode")
	return
}
