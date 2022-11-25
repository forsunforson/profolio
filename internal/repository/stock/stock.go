package stock

import (
	"fmt"
	"net/http"

	"github.com/forsunforson/profolio/internal/model"
)

type Repository interface {
	GetStockByCode(code string) (*model.Stock, error)
}

type stockImpl struct {
}

func NewRepository() Repository {
	return &stockImpl{}
}

func (i *stockImpl) GetStockByCode(code string) (*model.Stock, error) {
	resp, _ := http.Get("http://hq.sinajs.cn/list=sh601500")
	fmt.Printf("get resp: %v\n", resp)
	return nil, nil
}
