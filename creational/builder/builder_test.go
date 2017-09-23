package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	// Instantiate once the Director
	director := AnimalDirector{}

	// Create an empty dog
	dogBuilder := &DogBuilder{}
	// Construct a dog
	director.SetBuilder(dogBuilder)
	director.Construct()

	// Retrieve the constructed dog
	dog := dogBuilder.GetAnimal()

	// Test if it's a dog
	if dog.Type != "dog" {
		t.Error("This should be a dog")
	}
	if dog.Legs != true {
		t.Error("Dogs should have legs")
	}
	if dog.Wings != false {
		t.Error("Dogs shouldn't have wings")
	}
}
