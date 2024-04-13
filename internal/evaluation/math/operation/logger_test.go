package operation

import (
	"GoAlgorithmShuntingYard/kit/logger"
	"context"
	"reflect"
	"testing"
)

func TestLogger_Encrypt(t *testing.T) {
	type fields struct {
		next   IEvaluation
		logger logger.ILogger
	}
	type args struct {
		ctx     context.Context
		request Request
	}

	ctx := context.Background()
	newLogger := logger.NewLoggerDebug()
	service := ServiceMock{WantError: false}
	errService := ServiceMock{WantError: true}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp Response
		wantErr  bool
	}{
		{
			name: "Test Logger Successful",
			fields: fields{
				next:   service,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "5+4"},
			},
			wantResp: Response{
				Infix:   "5+4",
				Postfix: "5 4 +",
				Result:  1,
			},
			wantErr: false,
		},
		{
			name: "Test Logger Failure",
			fields: fields{
				next:   errService,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "5+4"},
			},
			wantResp: Response{},
			wantErr:  true,
		},
		{
			name: "Test Logger Failure PointerNil",
			fields: fields{
				next:   nil,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "5+4"},
			},
			wantResp: Response{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				next:   tt.fields.next,
				logger: tt.fields.logger,
			}
			gotResp, err := l.Evaluate(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Encrypt() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		next   IEvaluation
		logger logger.ILogger
	}

	service := ServiceMock{WantError: false}

	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "Test New Logger",
			args: args{
				next:   service,
				logger: nil,
			},
			want: NewLogger(service, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.next, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
