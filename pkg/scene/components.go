package scene

type Component interface {
	GetType() string
	SetRef(a *Actor)
	Update()
	// TODO something like this
}
