package container

type IContainer interface {
	Exists(string) (bool, error)
	YoungerThan(string, float64) (bool, error)
	CreateContainerName(string, string, string) (string, error)
}

func Exists(i IContainer, containerID string) (bool, error) {
	return i.Exists(containerID)
}

func YoungerThan(i IContainer, containerID string, age float64) (bool, error) {
	return i.YoungerThan(containerID, age)
}

func CreateContainerName(i IContainer, containerID string, containerTag string, containerTagReplacementValue string) (string, error) {
	return i.CreateContainerName(containerID, containerTag, containerTagReplacementValue)
}
