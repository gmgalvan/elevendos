package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"lab/productLab/internal/entity"
	"lab/productLab/internal/transport/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func setupUC(t *testing.T) (*gomock.Controller, *mocks.MockProductUC) {
	mockCtrl := gomock.NewController(t)
	mockProductUC := mocks.NewMockProductUC(mockCtrl)
	return mockCtrl, mockProductUC
}

func TestProductTransport_CreateProduct(t *testing.T) {
	now := time.Time{}
	comment := "comment"
	fakeCreatedProduct := &entity.Product{
		ID:        1,
		Name:      "prod-name",
		Price:     1.0,
		Comments:  &comment,
		Timestamp: &now,
	}
	fakeProduct := entity.Product{
		Name:      "prod-name",
		Price:     1.0,
		Comments:  &comment,
		Timestamp: &now,
	}
	fakeProductBytes, _ := json.Marshal(fakeProduct)
	fakeBody := bytes.NewBuffer(fakeProductBytes)
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *entity.Product
	}{
		{
			name: "Success Create a new product",
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest(http.MethodPost, "/products", fakeBody),
				ctx: context.TODO(),
			},
			want:    fakeCreatedProduct,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl, uc := setupUC(t)
			defer mockCtrl.Finish()
			pt := ProductTransport{
				Product: uc,
			}
			uc.EXPECT().Create(gomock.Any(), &fakeProduct).Return(fakeCreatedProduct, nil)
			pt.CreateProduct(tt.args.w, tt.args.r)
		})
	}
}
