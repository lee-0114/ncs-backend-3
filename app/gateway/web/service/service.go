package service

import "backend/pkg/rpc"

func RegService(gws *rpc.Gateways) {
	regDonateService(gws)
	regUserService(gws)
}
