package platform

type IPlatform interface {
	CreateNamespace(string) error
}

func CreateNamespace(i IPlatform, namespace string) error {
	return i.CreateNamespace(namespace)
}
