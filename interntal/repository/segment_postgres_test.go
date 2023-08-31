package repository

import (
	"log"
	"testing"

	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestSegmentPostgres_CreateSegment(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewSegmentPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   ent.Segment
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO segments").WithArgs("Test").WillReturnRows(rows)
			},
			input: ent.Segment{
				Name: "Test",
			},
			want: 1,
		},
		{
			name: "Empty field",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO segments").WithArgs("").
					WillReturnRows(rows)
			},
			input: ent.Segment{
				Name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.CreateSegment(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}
