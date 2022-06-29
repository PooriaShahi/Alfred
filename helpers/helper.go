package helpers

import (
	"encoding/xml"

	"github.com/vifraa/gopom"
)

func GetJarFileName(data []byte) string {
	var parsedPom gopom.Project
	err := xml.Unmarshal(data, &parsedPom)
	if err != nil {
		panic(err)
	}
	if parsedPom.Version == "" {
		name := parsedPom.ArtifactID + ".jar"
		return name
	}
	name := parsedPom.ArtifactID + "-" + parsedPom.Version
	return name
}
