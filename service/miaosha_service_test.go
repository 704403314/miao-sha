package service

import (
	"context"
	"miaosha/pb"

	// "miaosha/service"
	"testing"

	"github.com/stretchr/testify/require"
)

var mockGoodsId string = "99"
var mockStockNum int32 = 33

func TestGetStock(t *testing.T) {
	t.Parallel()
	stock := NewStockService()
	req := &pb.SearchStock{
		GoodsId: mockGoodsId,
	}
	expectedID := mockGoodsId
	res, err := stock.GetStock(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.GoodsId)
	require.Equal(t, mockStockNum, res.StockNum)
}
