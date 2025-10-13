package data

type Section interface {
	Title() string
	RenderTeX() string
}