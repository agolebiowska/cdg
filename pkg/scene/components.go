package scene

type componentType string

var (
	Physics   componentType = "physics"
	Animation componentType = "animation"
)

type Component interface {
	GetType() componentType
	SetRef(a *Actor)
	Update()
	// TODO something like this
}
