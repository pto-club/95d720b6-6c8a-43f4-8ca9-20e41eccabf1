//go:generate fyne bundle -o binary.go assets
package main

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type smet4ikTheme struct{
	fyne.Theme
}

func NewSmet4ikTheme() fyne.Theme{
	return &smet4ikTheme{Theme: theme.DefaultTheme()}
}

func (s *smet4ikTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return s.Theme.Color(name, theme.VariantLight) // set always theme-VariantLight
}

func(s *smet4ikTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText{
		return 10
	}
	return s.Theme.Size(name)
}