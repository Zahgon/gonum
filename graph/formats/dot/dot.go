//go:generate ./makeinternal.bash

package dot

import (
	"io"

	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

func ParseFile(path string) (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }

func Parse(r io.Reader) (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseBytes(b []byte) (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseString(s string) (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }
