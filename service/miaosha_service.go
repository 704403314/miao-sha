package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"log"
	"miaosha/common"
	"miaosha/pb"
	"time"
	"strconv"
)

type StockService struct {
}

func NewStockService() *StockService {
	return &StockService{}
}

func (self *StockService) GetStock(ctx context.Context, req *pb.SearchStock) (*pb.StockInfo, error) {
	goodsId := req.GoodsId
	log.Printf("GetStock receive goods id: %s", goodsId)
	if goodsId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "goods id is empty %v", goodsId)
	}
	
	StockNum,err := common.Get(fmt.Sprintf("real_stock_%s", goodsId))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "redis get stock_%s error", goodsId)
	}
	StockNum64, err := strconv.ParseInt(StockNum, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "stock num change to int64 error %s ", StockNum)
	}
	StockNum32 := int32(StockNum64)
	res := &pb.StockInfo{
		GoodsId:  goodsId,
		StockNum: StockNum32,
	}
	log.Printf("GetStock response info: %v", res)
	return res, nil
}

func (self *StockService) PreDeduct(ctx context.Context, req *pb.PreDeductRequest) (response *pb.PreDeductResponse, err error) {
	goodsId := req.GoodsId
	if goodsId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "goods id is empty %v", goodsId)
	}

	guid := uuid.New()
	guidStr := guid.String()
	setRes, err := common.SetNX("stock_"+goodsId, guidStr, 10*time.Second)

	if err != nil || setRes == false {
		fmt.Printf("set nx failed err:%v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "set nx failed err %v ", err)
	}

	err = self.pushList(goodsId)

	if err != nil {
		err = self.delKey(goodsId, guidStr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "del nx failed err %v ", err)
		}
		return nil, status.Errorf(codes.Unavailable, "deal Stock pre_stock_%s error", goodsId)
	}

	err = self.delKey(goodsId, guidStr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "del nx failed err %v ", err)
	}

	res := &pb.PreDeductResponse{
		GoodsId:  goodsId,
		UserId: req.UserId,
	}

	log.Printf("PreDeduct response info: %v", res)
	return res, nil
}

func (self *StockService) delKey(goodsId string, guidStr string) error {
	// write lua script to del key
	script :=   "if redis.call('get',KEYS[1]) == ARGV[1] then \n" +
					"return redis.call('del',KEYS[1]) \n" +
				"else\n " +
					"return false\n" +
				"end; "
	err := common.Exec(script, []string{"stock_"+goodsId}, []string{guidStr})
	if err != nil {
		fmt.Printf("exec lua script failed err:%v", err)
		return status.Errorf(codes.Internal, "exec lua script failed err:%v", err)
	}
	return nil
}
func (self *StockService) CancelTransaction(ctx context.Context, req *pb.CancelTransactionRequest) (*pb.PreDeductResponse, error) {
	goodsId := req.GoodsId
	if goodsId == "" {
		log.Printf("goods id is empty: %v", goodsId)
		return nil, status.Errorf(codes.InvalidArgument, "goods id is empty %v", goodsId)
	}

	err := common.LPush(fmt.Sprintf("list_stock_%s", goodsId), "1")
	if err != nil {
		fmt.Printf("push list failed err:%v", err)
		return nil, status.Errorf(codes.InvalidArgument, "push list failed err %v ", err)
	}

	//afterIncrVal, err := common.Incr(fmt.Sprintf("pre_stock_%s", goodsId))
	//if err != nil {
	//	log.Printf("CancelTransaction incr error: %v", err)
	//	return nil, status.Errorf(codes.InvalidArgument, "incr error %v ", err)
	//}

	res := &pb.PreDeductResponse{
		GoodsId:  goodsId,
		UserId: req.UserId,
	}

	log.Printf("cancel response info:%v", res)
	return res, nil
}

func (self *StockService) decrStock(goodsId string) error {
	// search stock deduct pre stock
	beforeStockStr, err := common.Get(fmt.Sprintf("pre_stock_%s", goodsId))
	if err != nil {
		log.Printf("get beforeStock failed err: %v", err)
		return status.Errorf(codes.Internal, "get beforeStock  err %v ", err)
	}
	beforeStock64, err := strconv.ParseInt(beforeStockStr,10,64)
	if err != nil {
		log.Printf("incrStock strconv err:%v", err)
		return status.Errorf(codes.Internal, "incrStock strconv err: %v ", err)
	}
	beforeStock32 := int32(beforeStock64)

	if beforeStock32 <= 0 {
		return status.Errorf(codes.Internal, "stock is not enough: pre_stock_%s ", goodsId)
	}

	PreStockNum, err := common.Decr(fmt.Sprintf("pre_stock_%s", goodsId))
	if err != nil {
		return status.Errorf(codes.Unavailable, "redis get pre_stock_%s error", goodsId)
	}

	PreStockNum32 := int32(PreStockNum)
	if PreStockNum32 < 0 {
		common.SetInt32(fmt.Sprintf("pre_stock_%s", goodsId), 0)
		return status.Errorf(codes.InvalidArgument, "pre stock num is negative %s ", PreStockNum32)
	}
	return nil
}

func (self *StockService) pushList(goodsId string) error {
	listLen, err := common.LLen(fmt.Sprintf("list_stock_%s", goodsId))
	if err != nil {
		fmt.Printf("get len failed err:%v", err)
		return status.Errorf(codes.InvalidArgument, "get len failed: %v", goodsId)
	}
	if listLen <= 0 {
		return status.Errorf(codes.InvalidArgument, "len is empty %v ", goodsId)
	}
	list,err := common.Pop(fmt.Sprintf("list_stock_%s", goodsId))
	if err != nil {
		fmt.Printf("get list failed err:%v", err)
		return status.Errorf(codes.InvalidArgument, "get list failed err %v ", err)
	}
	if list == "" {
		fmt.Printf("list is empty:%v", err)
		return status.Errorf(codes.InvalidArgument, "list is empty %v ", list)
	}
	return nil
}


