package paintimage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPainter(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		wantErr    bool
		wantConfig PaintConfig
	}{
		{
			name:       "valid_file_path",
			configPath: "./config_test.json",
			wantErr:    false,
			wantConfig: PaintConfig{
				Template: "yeezy-outilined.png",
				Masks: []MaskConfig{
					{Path: "mask1.png", Color: "random"},
					{Path: "mask2.png", Color: "random"},
					{Path: "mask3.png", Color: "random"},
				},
			},
		},
		{
			name:       "invalid_file_path",
			configPath: "../invalid_config_test.json",
			wantErr:    true,
		},
		{
			name:       "inexistent_file_path",
			configPath: "../inexistent_config_test.json",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			painter, err := NewPainter(tt.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Painter.NewPainter() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !cmp.Equal(painter.Config, tt.wantConfig) {
				t.Errorf("Painter.NewPainter() configRecived = %v, wantedConfig = %v", painter, tt.wantConfig)
			}
		})
	}
}
