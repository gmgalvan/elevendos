package usecase

import (
	"context"
	"lab/productLab/internal/entity"
	"lab/productLab/internal/usecase/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func setupProductStore(t *testing.T) (*gomock.Controller, *mocks.MockProductStore) {
	mockCtrl := gomock.NewController(t)
	mockProductStore := mocks.NewMockProductStore(mockCtrl)
	return mockCtrl, mockProductStore
}

func TestProductUC_Create(t *testing.T) {
	now := time.Now()
	comment := "comment"
	fakeProduct := &entity.Product{
		Name:      "prod-name",
		Price:     1.0,
		Comments:  &comment,
		Timestamp: &now,
	}
	fakeCreatedProduct := &entity.Product{
		ID:        1,
		Name:      "prod-name",
		Price:     1.0,
		Comments:  &comment,
		Timestamp: &now,
	}
	type args struct {
		ctx     context.Context
		product *entity.Product
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Product
		wantErr bool
	}{
		{
			name: "Success: New Product",
			args: args{
				ctx:     context.TODO(),
				product: fakeProduct,
			},
			want:    fakeCreatedProduct,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl, store := setupProductStore(t)
			defer mockCtrl.Finish()
			testUS := &ProductUC{
				Store: store,
			}
			store.EXPECT().Create(tt.args.ctx, fakeProduct).Return(fakeCreatedProduct, nil)
			got, err := testUS.Create(tt.args.ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductUC.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductUC.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
