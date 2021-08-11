package paintimage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadConfigFile(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		wantErr      bool
		wantedConfig ConfigFile
	}{
		{
			name:    "valid_file_path",
			path:    "./config_test.json",
			wantErr: false,
			wantedConfig: ConfigFile{
				Template: "yeezy-outilined.png",
				Masks:    []string{"mask1.png", "mask2.png", "mask3.png"},
			},
		},
		{
			name:    "valid_file_path",
			path:    "../invalid_config.json",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configFile, err := ReadConfigFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if !cmp.Equal(configFile, tt.wantedConfig) {
				t.Errorf("configRecived = %v, wantedConfig = %v", configFile, tt.wantedConfig)
			}
		})
	}
}
