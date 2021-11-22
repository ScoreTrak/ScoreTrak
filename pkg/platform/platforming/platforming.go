package platforming

type Config struct {
	Use    string
	Docker struct {
		Name    string
		Host    string
		Network string
	}
	Kubernetes struct {
		Namespace string
	}
}
