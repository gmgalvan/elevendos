package store

import (
	"context"
	"lab/productLab/internal/entity"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func setup(t *testing.T) (*Store, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	st := NewStore(db)
	return st, mock
}

func TestStore_Create(t *testing.T) {
	now := time.Now()
	comment := "comment"
	type fields struct {
		create func() *Store
	}
	type args struct {
		ctx     context.Context
		product *entity.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Product
		wantErr bool
	}{
		{
			name: "Succes create a product",
			fields: fields{
				create: func() *Store {
					s, mock := setup(t)
					rows := sqlmock.NewRows([]string{"id"}).
						AddRow("1")
					mock.ExpectQuery("INSERT").
						WithArgs("prod-name", 1.0, comment, now).
						WillReturnRows(rows)
					return s
				},
			},
			args: args{
				ctx: context.TODO(),
				product: &entity.Product{
					Name:      "prod-name",
					Price:     1.0,
					Comments:  &comment,
					Timestamp: &now,
				},
			},
			want: &entity.Product{
				ID:        1,
				Name:      "prod-name",
				Price:     1.0,
				Comments:  &comment,
				Timestamp: &now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.create()
			defer store.client.Close()
			s := &Store{
				client: store.client,
			}
			got, err := s.Create(tt.args.ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
