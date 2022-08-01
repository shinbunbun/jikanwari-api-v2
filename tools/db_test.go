package tools

import (
	"reflect"
	"testing"

	"github.com/guregu/dynamo"
)

func getTable() *dynamo.Table {
	local := Local{}
	var config Config = &local
	table := config.GetDynamoDB()
	return &table
}

func TestGetByIDDataType(t *testing.T) {
	table := getTable()

	type args struct {
		id    string
		table *dynamo.Table
	}
	tests := []struct {
		name    string
		args    args
		want    *DynamoItem
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id:    "test_user",
				table: table,
			},
			want: &DynamoItem{
				ID:   "test_user",
				Date: "2022年07月30日 (土) 0時0分0秒",
				Flag: "25",
				Fri:  "friday",
				Mon:  "monday",
				Sat:  "saturday",
				Thu:  "thursday",
				Tue:  "tuesday",
				Uuid: "test_id",
				Wed:  "wednesday",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				id:    "failed_user",
				table: table,
			},
			want:    &DynamoItem{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByID(tt.args.id, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIDDataType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIDDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPut(t *testing.T) {
	table := getTable()

	type args struct {
		item  *DynamoItem
		table *dynamo.Table
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				item: &DynamoItem{
					ID:   "add_test_user",
					Date: "2022年07月30日 (土) 0時0分0秒",
					Flag: "25",
					Fri:  "friday",
					Mon:  "monday",
					Sat:  "saturday",
					Thu:  "thursday",
					Tue:  "tuesday",
					Uuid: "test_id",
					Wed:  "wednesday",
				},
				table: table,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Put(tt.args.item, tt.args.table); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}

			item, error := GetByID(tt.args.item.ID, tt.args.table)
			if error != nil {
				t.Errorf("GetByID() error = %v", error)
			}
			if !reflect.DeepEqual(item, tt.args.item) {
				t.Errorf("GetByID() = %v, want %v", item, tt.args.item)
			}
		})
	}
}
