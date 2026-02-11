package pocketbase_ts_generator

import (
	"github.com/Vogeslu/pocketbase-ts-generator/internal/cmd"
	"github.com/pocketbase/pocketbase"
	pbcore "github.com/pocketbase/pocketbase/core"
)

type GeneratorOptions struct {
	AllCollections     bool
	CollectionsInclude []string
	CollectionsExclude []string

	Output string

	IndentSize int
	UseInterface bool
}

func RegisterHook(app *pocketbase.PocketBase, options *GeneratorOptions) {
	indentSize := options.IndentSize
	if indentSize == 0 {
		indentSize = 2
	}

	generatorFlags := &cmd.GeneratorFlags{
		AllCollections:     options.AllCollections,
		CollectionsInclude: options.CollectionsInclude,
		CollectionsExclude: options.CollectionsExclude,

		Output:     options.Output,
		IndentSize: indentSize,
		UseInterface: options.UseInterface,
	}

	app.OnCollectionAfterCreateSuccess().BindFunc(func(e *pbcore.CollectionEvent) error {
		_ = processFileGeneration(app, generatorFlags)

		return e.Next()
	})

	app.OnCollectionAfterUpdateSuccess().BindFunc(func(e *pbcore.CollectionEvent) error {
		_ = processFileGeneration(app, generatorFlags)

		return e.Next()
	})

	app.OnCollectionAfterDeleteSuccess().BindFunc(func(e *pbcore.CollectionEvent) error {
		_ = processFileGeneration(app, generatorFlags)

		return e.Next()
	})
}
