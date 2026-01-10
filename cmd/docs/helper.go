package docs

import (
	"embed"

	"github.com/swaggo/swag"
)

const swaggerErrMsg = "the swagger could not be loaded"
const docsPath = "docs.json"

// SpecEmbed implements swag.Swagger.
type SpecEmbed struct {
	// Name is the name of the api-doc file.
	Name string
	// Docs is the embed file to be used as API Doc Spec.
	Docs *embed.FS
}

func (y *SpecEmbed) ReadDoc() string {
	d, err := y.Docs.ReadFile(y.Name)
	if err != nil {
		return swaggerErrMsg
	}
	return string(d)
}

// Init registries the specification into Swaggo.
func (y *SpecEmbed) Init() {
	swag.Register(swag.Name, y)
}

func NewSpec(d *embed.FS) *SpecEmbed {
	return &SpecEmbed{Docs: d, Name: docsPath}
}
