package builder

// BuildProcess provides the interface of the builder
type BuildProcess interface {
	SetType() BuildProcess
	SetLegs() BuildProcess
	SetWings() BuildProcess
}

// AnimalDirector is the one in charge of providing the builder
type AnimalDirector struct {
	builder BuildProcess
}

// Construct concatenates the methods and abstracts the object creation
func (f *AnimalDirector) Construct() {
	f.builder.SetType().SetLegs().SetWings()
}

// SetBuilder provides the individual access to the builder
func (f *AnimalDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Animal structure
type Animal struct {
	Type  string
	Legs  bool
	Wings bool
}

// DogBuilder type
type DogBuilder struct {
	a Animal
}

// SetType sets the properties for a dog
func (d *DogBuilder) SetType() BuildProcess {
	d.a.Type = "dog"
	return d
}

// SetLegs sets the properties for a dog
func (d *DogBuilder) SetLegs() BuildProcess {
	d.a.Legs = true
	return d
}

// SetWings sets the properties for a dog
func (d *DogBuilder) SetWings() BuildProcess {
	d.a.Wings = false
	return d
}

// GetAnimal retrieves the constructed object
func (d *DogBuilder) GetAnimal() Animal {
	return d.a
}
