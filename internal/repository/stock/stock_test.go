package stock

import (
	"reflect"
	"testing"

	"github.com/forsunforson/profolio/internal/model"
)

func Test_stockImpl_GetStockByCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		i       *stockImpl
		args    args
		want    *model.Stock
		wantErr bool
	}{
		{
			name:    "case1",
			i:       &stockImpl{},
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &stockImpl{}
			got, err := i.GetStockByCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("stockImpl.GetStockByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stockImpl.GetStockByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
