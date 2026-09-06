package optimize

import (
	"io"
	"time"
)

var printerHeadings = [...]string{
	"Iter",
	"Runtime",
	"FuncEvals",
	"Func",
	"GradEvals",
	"|Gradient|∞",
	"HessEvals",
}

const (
	printerBaseTmpl = "%9v  %16v  %9v  %22v"
	printerGradTmpl = "  %9v  %22v"
	printerHessTmpl = "  %9v"
)

var _ Recorder = (*Printer)(nil)

type Printer struct {
	Writer          io.Writer
	HeadingInterval int
	ValueInterval   time.Duration

	lastHeading int
	lastValue   time.Time
}

func NewPrinter() *Printer { _ = "STUB: not implemented"; return nil }

func (p *Printer) Init() error { _ = "STUB: not implemented"; return nil }

func (p *Printer) Record(loc *Location, op Operation, stats *Stats) error {
	_ = "STUB: not implemented"
	return nil
}
