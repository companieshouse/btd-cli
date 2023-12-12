package btd

type Renderer interface {
	Render(data TagData) string
}
