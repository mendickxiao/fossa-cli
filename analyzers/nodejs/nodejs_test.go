package nodejs_test

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fossas/fossa-cli/analyzers"
	"github.com/fossas/fossa-cli/module"
	"github.com/fossas/fossa-cli/pkg"
)

// TestNoDependencies checks that there is no error even when `package.json` is
// missing a `dependencies` key or has an empty object as the value for
// `dependencies`.
func TestNoDependencies(t *testing.T) {
	m := module.Module{
		Dir:  filepath.Join("testdata", "empty"),
		Type: pkg.NodeJS,
	}

	a, err := analyzers.New(m)
	assert.NoError(t, err)

	deps, err := a.Analyze()
	assert.NoError(t, err)
	assert.Empty(t, deps.Direct)
	assert.Empty(t, deps.Transitive)
}

// TestDuplicateDependencies checks that analysis correctly handles duplicate
// dependencies, even in the case where duplicates may not have the same set of
// imports listed.
//
// For example, try running `npm ls --json` in the `testdata/duplicates` folder.
// Notice that `babel-runtime` is included as a dependency twice: once by
// `babel-polyfill` and once by `jira-client`. However, the _dependencies_ of
// `babel-runtime` are only listed _once,_ when it's imported by
// `babel-polyfill`. This means that we must ensure that we get transitive
// dependencies from the original dependency entry, not the deduplicated entry.
//
// See #257 for details.
func TestDuplicateDependencies(t *testing.T) {
	m := module.Module{
		BuildTarget: filepath.Join("testdata", "duplicates", "package.json"),
		Dir:         filepath.Join("testdata", "duplicates"),
		Type:        pkg.NodeJS,
	}

	a, err := analyzers.New(m)
	assert.NoError(t, err)

	// We run this multiple times because this bug may flake; map traversal order
	// is random in Go.
	var failed pkg.Deps
	for i := 0; i < 10; i++ {
		deps, err := a.Analyze()
		assert.NoError(t, err)
		id := pkg.ID{
			Type:     pkg.NodeJS,
			Name:     "regenerator-runtime",
			Revision: "0.11.1",
			Location: "https://registry.npmjs.org/regenerator-runtime/-/regenerator-runtime-0.11.1.tgz",
		}
		ok := assert.Contains(t, deps.Transitive, id)
		if !ok {
			failed = deps.Transitive
		}
	}

	if t.Failed() {
		log.Printf("%#v", failed)
	}
}
