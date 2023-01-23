package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/widget"
)

/**
 * Main struct of image button widget.
 */
type ImageButton struct {
	fyne.WidgetRenderer
	widget.BaseWidget
	objects    []fyne.CanvasObject
	OnTapped   func(*ImageButton, *canvas.Image) `json:"-"`
	imagePath  string
	image      *canvas.Image
	size       fyne.Size
	hovered    bool
	background *canvas.Rectangle
}

/**
 * Public method for create image button.
 * @param {string} image
 * @param {func(*ImageButton)} OnTapped
 * @param {fyne.Size} size
 * @return {*ImageButton}
 */
func NewImageButton(image string, OnTapped func(*ImageButton, *canvas.Image), size fyne.Size) *ImageButton {
	ib := &ImageButton{
		imagePath: image,
		OnTapped:  OnTapped,
		size:      size,
	}

	ib.image = ib.CreateImage(ib.imagePath)
	widget.NewSimpleRenderer(ib).Layout(ib.MinSize())
	return ib
}

/**
 * Public method to set cursor.
 * @return {desktop.Cursor}
 */
func (ib *ImageButton) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

/**
 * Public method for create render of custom widget.
 * @return {fyne.WidgetRenderer}
 */
func (ib *ImageButton) CreateRenderer() fyne.WidgetRenderer {
	ib.background = canvas.NewRectangle(theme.ButtonColor())
	ib.background.Resize(ib.MinSize())

	ib.objects = []fyne.CanvasObject{
		ib.background,
		ib.image,
	}

	return ib
}

/**
 * Public method for create image.
 * @param {string} path
 * @return {*canvas.Image}
 */
func (ib *ImageButton) CreateImage(path string) *canvas.Image {
	image := canvas.NewImageFromFile(path) // this is an image.Image
	image.Resize(ib.Size())
	image.Translucency = 0

	return image
}

/**
 * Public method for get min size.
 * @return {fyne.Size}
 */
func (ib *ImageButton) MinSize() fyne.Size {
	return fyne.Size{
		Width:  ib.size.Width,
		Height: ib.size.Height,
	}
}

/**
 * Public method for refresh canvas.
 */
func (ib *ImageButton) Refresh() {
	canvas.Refresh(ib)
}

/**
 * Public method get size.
 * @return {fyne.Size}
 */
func (ib *ImageButton) Size() fyne.Size {
	return fyne.Size{
		Width:  ib.size.Width,
		Height: ib.size.Height,
	}
}

/**
 * Public method get objects in widget.
 * @return {[]fyne.CanvasObject}
 */
func (ib *ImageButton) Objects() []fyne.CanvasObject {
	return ib.objects
}

/**
 * Public method to set opacity for image.
 * @params {float64} opacity
 */
func (ib *ImageButton) SetOpacity(opacity float64) {
	if ib.image != nil {
		opacity = 1 - opacity

		anim := fyne.NewAnimation(canvas.DurationStandard, func(done float32) {
			if ib.image.Translucency < opacity {
				ib.image.Translucency += 0.1
			} else {
				ib.image.Translucency -= 0.1
			}

			ib.Refresh()
		})

		anim.Start()
	}
}

/**
 * Public method fired when tapped on widget.
 * @param {*fyne.PointEvent}
 */
func (ib *ImageButton) Tapped(*fyne.PointEvent) {
	if ib.OnTapped != nil {
		ib.OnTapped(ib, ib.image)
	}
}

/**
 * Public method fired when mouse in.
 * @param {*desktop.MouseEvent}
 */
func (ib *ImageButton) MouseIn(*desktop.MouseEvent) {
	ib.hovered = true

	ib.ApplyButtonTheme()
}

/**
 * Public method fired when mouse out.
 * @param {*desktop.MouseEvent}
 */
func (ib *ImageButton) MouseOut() {
	ib.hovered = false

	ib.ApplyButtonTheme()
}

/**
 * Public method for apply button theme.
 */
func (ib *ImageButton) ApplyButtonTheme() {
	if ib.background == nil {
		return
	}

	ib.background.FillColor = ib.ButtonColor()
	ib.background.Refresh()
}

/**
 * Public method for get button color.
 * @return {color.Color}
 */
func (ib *ImageButton) ButtonColor() color.Color {
	switch {
	case ib.hovered:
		return BlendColor(theme.ButtonColor(), theme.HoverColor())
	default:
		return theme.ButtonColor()
	}
}

/**
 * Public method fired when mouse moved.
 * @param {*desktop.MouseEvent}
 */
func (ib *ImageButton) MouseMoved(*desktop.MouseEvent) {
	//...
}

/**
 * Public method for set layout.
 * @param {fyne.Size} size
 */
func (ib *ImageButton) Layout(size fyne.Size) {
	//...
}

/**
 * Public method for apply theme.
 */
func (*ImageButton) ApplyTheme() {
	//...
}

/**
 * Public method fired when destory widget.
 */
func (ib *ImageButton) Destroy() {
	//...
}

/**
 * Function to apply color.
 * @param {color.Color} under
 * @param {color.Color} over
 * @return {color.Color}
 */
func BlendColor(under, over color.Color) color.Color {
	dstR, dstG, dstB, dstA := under.RGBA()
	srcR, srcG, srcB, srcA := over.RGBA()

	srcAlpha := float32(srcA) / 0xFFFF
	dstAlpha := float32(dstA) / 0xFFFF

	outAlpha := srcAlpha + dstAlpha*(1-srcAlpha)
	outR := srcR + uint32(float32(dstR)*(1-srcAlpha))
	outG := srcG + uint32(float32(dstG)*(1-srcAlpha))
	outB := srcB + uint32(float32(dstB)*(1-srcAlpha))

	return color.RGBA64{R: uint16(outR), G: uint16(outG), B: uint16(outB), A: uint16(outAlpha * 0xFFFF)}
}
