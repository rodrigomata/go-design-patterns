package singleton

type singleton struct {
}

var instance *singleton

// GetInstance checks if an instance exists or creates one
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton) // or &singleton{} <-- Both are not thread safe
	}
	return instance
}
