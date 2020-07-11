package scene

type componentType string

var (
	Physics     componentType = "physics"
	Animation   componentType = "animation"
	Combat      componentType = "combat"
	Dialogue    componentType = "dialogue"
	Interaction componentType = "interaction"
)

type Component interface {
	GetType() componentType
	SetRef(a *Actor)
	Update()
	// TODO something like this
}
