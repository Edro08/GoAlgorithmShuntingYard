package operation

import (
	"GoAlgorithmShuntingYard/internal/evaluation/math"
	"GoAlgorithmShuntingYard/kit/config"
	"GoAlgorithmShuntingYard/kit/logger"
	"context"
	"reflect"
	"testing"
)

func TestNewEnvExprMathService(t *testing.T) {
	type args struct {
		config config.IConfig
		logger logger.ILogger
	}
	newConfig := config.NewConfig("../../../../application.yaml")
	newLogger := logger.NewLoggerDebug()

	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "TEST NewEnvExprMathService Successful",
			args: args{
				config: newConfig,
				logger: newLogger,
			},
			want: Service{
				config: newConfig,
				logger: newLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnvExprMathService(tt.args.config, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvExprMathService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Evaluate(t *testing.T) {
	type fields struct {
		config config.IConfig
		logger logger.ILogger
	}
	type args struct {
		ctx     context.Context
		request Request
	}

	ctx := context.Background()
	newConfig := config.NewConfig("../../../../application.yaml")
	newLogger := logger.NewLoggerDebug()

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse Response
		wantErr      bool
	}{
		{
			name: "TEST Evaluate Successful",
			fields: fields{
				config: newConfig,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "3 + 4 * 2 / 1 - 5"},
			},
			wantResponse: Response{
				Infix:   "3 + 4 * 2 / 1 - 5",
				Postfix: "3 4 2 * 1 / + 5 -",
				Result:  6,
			},
			wantErr: false,
		},
		{
			name: "TEST Evaluate Infix Empty",
			fields: fields{
				config: newConfig,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: ""},
			},
			wantResponse: Response{},
			wantErr:      true,
		},
		{
			name: "TEST Evaluate Infix Invalid ()",
			fields: fields{
				config: newConfig,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "(5+4)"},
			},
			wantResponse: Response{},
			wantErr:      true,
		},
		{
			name: "TEST Evaluate Infix Invalid",
			fields: fields{
				config: newConfig,
				logger: newLogger,
			},
			args: args{
				ctx:     ctx,
				request: Request{Infix: "5+4 Hola"},
			},
			wantResponse: Response{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Service{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			gotResponse, err := h.Evaluate(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Evaluate() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestService_evaluateExpressionPostfix(t *testing.T) {
	type fields struct {
		config config.IConfig
		logger logger.ILogger
	}
	type args struct {
		ctx        context.Context
		expression math.ExpressionPostfix
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Service{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			got, err := h.evaluateExpressionPostfix(tt.args.ctx, tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluateExpressionPostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("evaluateExpressionPostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_createExpressionPostfix(t *testing.T) {
	type fields struct {
		config config.IConfig
		logger logger.ILogger
	}
	type args struct {
		ctx   context.Context
		infix math.ExpressionInfix
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    math.ExpressionPostfix
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Service{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			got, err := h.createExpressionPostfix(tt.args.ctx, tt.args.infix)
			if (err != nil) != tt.wantErr {
				t.Errorf("execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
