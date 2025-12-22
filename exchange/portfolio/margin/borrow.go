package margin

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) MarginMaxBorrow(asset string) (*MaxBorrow, error) {
	borrow, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginMaxBorrowService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &MaxBorrow{
		Amount:      borrow.Amount,
		BorrowLimit: borrow.BorrowLimit,
	}, nil
}

func (c *Client) MarginMaxWithdraw(asset string) (string, error) {
	withdraw, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginMaxWithdrawService().Asset(asset).Do(context.Background())
	if err != nil {
		return "", err
	}
	return withdraw.Amount, nil
}
