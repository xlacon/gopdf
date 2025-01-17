package core

import "fmt"

// the units of below is pixels.
type Config struct {
	StartX, StartY float64 // PDF page start position
	EndX, EndY     float64 // PDF page end postion
	Width, Height  float64 // PDF page width and height

	ContentWidth, ContentHeight float64 // PDF page content width and height
}

// Params width, height is pdf page width and height
// Params padingH, padingV is pdf horizontal and vertical padding
// The units of the above parameters are pixels.
// Params width must more than 2*padingH, and height must more 2*padingV
func NewConfig(width, height float64, padingH, padingV float64) (*Config, error) {
	if width <= 0 || height <= 0 || padingH < 0 || padingV < 0 {
		return nil, fmt.Errorf("params must more than zero")
	}

	if width <= 2*padingH || height <= 2*padingV {
		return nil, fmt.Errorf("this config params invalid")
	}

	c := &Config{
		Width:  width,
		Height: height,

		StartX: padingH,
		StartY: padingV,

		ContentWidth:  width - 2*padingH,
		ContentHeight: height - 2*padingV,
	}

	c.EndX = c.StartX + c.ContentWidth
	c.EndY = c.StartY + c.ContentHeight

	return c, nil
}

func (config *Config) GetWidthAndHeight() (width, height float64) {
	return config.Width, config.Width
}

// Get pdf page start position, from the position you can write the pdf body content.
func (config *Config) GetStart() (x, y float64) {
	return config.StartX, config.StartY
}

func (config *Config) GetEnd() (x, y float64) {
	return config.EndX, config.EndY
}

/*
*************************************
A0 ~ A5 page width and height config:

	'A0': [2383.94, 3370.39],
	'A1': [1683.78, 2383.94],
	'A2': [1190.55, 1683.78],
	'A3': [841.89, 1190.55],
	'A4': [595.28, 841.89],
	'A5': [419.53, 595.28],

**************************************
*/
