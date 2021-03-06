package proteus

import (
	"github.com/kennyzhu/go-os/tools/proteus/protobuf"
	"github.com/kennyzhu/go-os/tools/proteus/resolver"
	"github.com/kennyzhu/go-os/tools/proteus/rpc"
	"github.com/kennyzhu/go-os/tools/proteus/scanner"
)

// Options are all the available options to configure proto generation.
type Options struct {
	BasePath string
	Packages []string
}

// store every proto file names  for every package...
// Todo:
type PackageFiles struct {
	Files map[string][]string // key = package name, value = files...
}

type generator func(*scanner.Package, *protobuf.Package) error

// generate for rpc generator...
// can't find import: "github.com/golang/protobuf/proto"
// go get github.com/golang/protobuf/proto
func transformToProtobuf(packages []string, generate generator) error {
	scanner, err := scanner.New(packages...)
	if err != nil {
		return err
	}

	pkgs, err := scanner.Scan()
	if err != nil {
		return err
	}

	r := resolver.New()
	r.Resolve(pkgs)

	t := protobuf.NewTransformer()
	t.SetStructSet(createStructTypeSet(pkgs))
	t.SetEnumSet(createEnumTypeSet(pkgs))
	for _, p := range pkgs {
		pkg := t.Transform(p)
		if err := generate(p, pkg); err != nil {
			return err
		}
	}

	return nil
}

func createStructTypeSet(pkgs []*scanner.Package) protobuf.TypeSet {
	ts := protobuf.NewTypeSet()
	for _, p := range pkgs {
		for _, s := range p.Structs {
			ts.Add(p.Path, s.Name)
		}
	}
	return ts
}

func createEnumTypeSet(pkgs []*scanner.Package) protobuf.TypeSet {
	ts := protobuf.NewTypeSet()
	for _, p := range pkgs {
		for _, e := range p.Enums {
			ts.Add(p.Path, e.Name)
		}
	}
	return ts
}

// GenerateProtos generates proto files for the given options.
// BasePath = -f path...
func GenerateProtos(options Options) error {
	g := protobuf.NewGenerator(options.BasePath)
	return transformToProtobuf(options.Packages, func(_ *scanner.Package, pkg *protobuf.Package) error {
		// return g.GeneratorDb(pkg)
		return g.Generator( pkg )
	})
}

// GenerateRPCServer generates the gRPC server implementation of the given
// packages.
func GenerateRPCServer(packages []string) error {
	// g := rpc.NewGenerator()
	g := rpc.NewMicroGenerator() // change to micro service ....
	return transformToProtobuf(packages, func(p *scanner.Package, pkg *protobuf.Package) error {
		return g.Generate(pkg, p.Path)
	})
}
