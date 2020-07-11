package scene

type componentType string

var (
	Physics     componentType = "physics"
	Animation   componentType = "animation"
	Combat      componentType = "combat"
	Dialogue    componentType = "dialogue"
	Interaction componentType = "interaction"
)

type component interface {
	getType() componentType
	setRef(a *actor)
	update()
	// TODO something like this
}
