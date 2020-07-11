package scene

type componentType string

// available component types
var (
	Physics     componentType = "physics"
	Animation   componentType = "animation"
	Combat      componentType = "combat"
	Dialogue    componentType = "dialogue"
	Interaction componentType = "interaction" // todo
)

type component interface {
	getType() componentType
	setRef(a *actor)
	update()
	draw()
}
