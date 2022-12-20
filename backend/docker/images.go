package images

import (
	"os/exec"
)

const dockerImagesCommand = "docker images --format {{.Repository}}:{{.Tag}}"

func GetDockerImages() []string {
	images := exec.Command("sh", "-c", dockerImagesCommand)
	out,
		err := images.Output()
	if err != nil {
		panic(err)
	}
	return []string{string(out)}
}
