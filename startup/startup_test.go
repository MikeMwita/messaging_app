package startup

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful initialization",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up test environment variables
			os.Setenv("DB_PASSWORD", "1A5hhh3qsjQQdUA6IajljFTXoDQKEcwo")
			os.Setenv("DB_CONNECTION_STRING", "postgres://mike:1A5hhh3qsjQQdUA6IajljFTXoDQKEcwo@dpg-clk5aoeg1b2c739gus30-a.oregon-postgres.render.com/mike_dy9f?sslmode=require")

			got, got1, err := Initialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				// If an error is expected, assert that both return values are nil
				assert.Nil(t, got, "Unexpected non-nil router when error is expected")
				assert.Nil(t, got1, "Unexpected non-nil repository when error is expected")
			} else {
				// If no error is expected, assert that both return values are not nil
				assert.NotNil(t, got, "Router should not be nil")
				assert.NotNil(t, got1, "Repository should not be nil")
			}
		})
	}
}
