package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-contrib/cors"
)

// TestCorsConfig tests the CorsConfig function.
func TestCorsConfig(t *testing.T) {
	type args struct {
		allowOrigins []string
		allowMethods []string
		allowHeaders []string
	}

	validHeaders := []string{"Allow"}
	validMethods := []string{"OPTIONS"}
	validOrigins := []string{"http://localhost"}
	emptySlice := []string{}

	validConfig := cors.DefaultConfig()
	validConfig.AllowOrigins = validOrigins
	validConfig.AllowHeaders = validHeaders
	validConfig.AllowMethods = validMethods
	validConfig.AllowCredentials = true

	emptyConfig := cors.DefaultConfig()
	emptyConfig.AllowOrigins = emptySlice
	emptyConfig.AllowHeaders = emptySlice
	emptyConfig.AllowMethods = emptySlice
	emptyConfig.AllowCredentials = true

	tests := []struct {
		name string
		args args
		want cors.Config
	}{
		{
			"Valid arguments",
			args{allowOrigins: validOrigins, allowMethods: validMethods, allowHeaders: validHeaders},
			validConfig,
		},
		{
			"Empty slices",
			args{emptySlice, emptySlice, emptySlice},
			emptyConfig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorsConfig(tt.args.allowOrigins, tt.args.allowMethods, tt.args.allowHeaders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CorsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
