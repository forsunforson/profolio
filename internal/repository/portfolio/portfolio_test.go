package portfolio

import (
	"fmt"
	"testing"

	"github.com/forsunforson/profolio/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_portfolioImpl_Save(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("/home/ywy/my_project/my_profolio/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect db fail: %s\n", err)
		return
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		porfolio *model.Portfolio
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "case1: save",
			fields: fields{db: db},
			args: args{
				porfolio: &model.Portfolio{
					ID:       0,
					Total:    20,
					Cash:     20,
					Accounts: []int64{0},
					Holders: map[int64]*model.Holder{
						0: {
							UserID:     0,
							Percentage: 1,
							Total:      20,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portfolioImpl{
				db: tt.fields.db,
			}
			if err := p.Save(tt.args.porfolio); (err != nil) != tt.wantErr {
				t.Errorf("portfolioImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
