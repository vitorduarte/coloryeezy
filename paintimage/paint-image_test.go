package paintimage

import (
	"image"
	"image/color"
	"reflect"
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

func TestPainter_Paint(t *testing.T) {
	type args struct {
		outputPath string
	}
	tests := []struct {
		name    string
		p       *Painter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Paint(tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("Painter.Paint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_openTemplateImage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantCanva *image.RGBA
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCanva, err := openTemplateImage(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("openTemplateImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCanva, tt.wantCanva) {
				t.Errorf("openTemplateImage() = %v, want %v", gotCanva, tt.wantCanva)
			}
		})
	}
}

func Test_openImage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openImage(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("openImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomColor(t *testing.T) {
	tests := []struct {
		name string
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToRGBA(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantC   color.RGBA
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := hexToRGBA(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToRGBA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("hexToRGBA() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func Test_fillImageWithColorWithMask(t *testing.T) {
	type args struct {
		canva    *image.RGBA
		maskPath string
		color    color.RGBA
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fillImageWithColorWithMask(tt.args.canva, tt.args.maskPath, tt.args.color); (err != nil) != tt.wantErr {
				t.Errorf("fillImageWithColorWithMask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saveImage(t *testing.T) {
	type args struct {
		img        image.Image
		outputPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveImage(tt.args.img, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("saveImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
