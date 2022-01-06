package checkout

import (
	"context"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-server/internal/pkg/logger"
	rpc "github.com/ZAF07/tigerlily-e-bakery-server/internal/pkg/protos"
	"github.com/ZAF07/tigerlily-e-bakery-server/internal/repository/checkout"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Service struct {
	db *gorm.DB
	base checkout.CheckoutRepo
	logs logger.Logger
}

func NewCheckoutService(DB *gorm.DB) *Service {
	return&Service{
		db: DB,
		base: *checkout.NewCheckoutRepo(DB),
		logs: *logger.NewLogger(),
	}
}

func (srv Service) Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	srv.logs.InfoLogger.Printf(" [SERVICE] Checkout service ran %+v", req)
	fmt.Printf("request : %+v", &req)

	checkoutSuccess, err := srv.base.CreateNewOrder(req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
		srv.logs.ErrorLogger.Printf(" [SERVICE] RESUKT FROM DS : %+v", checkoutSuccess)
	}

	srv.logs.InfoLogger.Printf(" [SERVICE] RESULT : %+v\n",checkoutSuccess)

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success: checkoutSuccess,
	}

	return 
} 
