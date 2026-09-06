package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"gonum.org/v1/gonum/dsp/window"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var windows = map[string]*builtin{
	"rectangular": {
		name: func(_ float64) string { return "Rectangular" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Rectangular },
		ok:   func(_ float64) bool { return true },
	},
	"sine": {
		name: func(_ float64) string { return "Sine" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Sine },
		ok:   func(_ float64) bool { return true },
	},
	"lanczos": {
		name: func(_ float64) string { return "Lanczos" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Lanczos },
		ok:   func(_ float64) bool { return true },
	},
	"triangular": {
		name: func(_ float64) string { return "Triangular" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Triangular },
		ok:   func(_ float64) bool { return true },
	},
	"hann": {
		name: func(_ float64) string { return "Hann" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Hann },
		ok:   func(_ float64) bool { return true },
	},
	"bartletthann": {
		name: func(_ float64) string { return "BartlettHann" },
		fn:   func(_ float64) func([]float64) []float64 { return window.BartlettHann },
		ok:   func(_ float64) bool { return true },
	},
	"hamming": {
		name: func(_ float64) string { return "Hamming" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Hamming },
		ok:   func(_ float64) bool { return true },
	},
	"blackman": {
		name: func(_ float64) string { return "Blackman" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Blackman },
		ok:   func(_ float64) bool { return true },
	},
	"blackmanharris": {
		name: func(_ float64) string { return "BlackmanHarris" },
		fn:   func(_ float64) func([]float64) []float64 { return window.BlackmanHarris },
		ok:   func(_ float64) bool { return true },
	},
	"nuttall": {
		name: func(_ float64) string { return "Nuttall" },
		fn:   func(_ float64) func([]float64) []float64 { return window.Nuttall },
		ok:   func(_ float64) bool { return true },
	},
	"blackmannuttall": {
		name: func(_ float64) string { return "BlackmanNuttall" },
		fn:   func(_ float64) func([]float64) []float64 { return window.BlackmanNuttall },
		ok:   func(_ float64) bool { return true },
	},
	"flattop": {
		name: func(_ float64) string { return "FlatTop" },
		fn:   func(_ float64) func([]float64) []float64 { return window.FlatTop },
		ok:   func(_ float64) bool { return true },
	},
	"gaussian": {
		name: func(p float64) string { return fmt.Sprintf("Gaussian σ=%v", p) },
		fn:   func(p float64) func([]float64) []float64 { return window.Gaussian{Sigma: p}.Transform },
		ok:   func(p float64) bool { return !math.IsNaN(p) },
	},
	"tukey": {
		name: func(p float64) string { return fmt.Sprintf("Tukey α=%v", p) },
		fn:   func(p float64) func([]float64) []float64 { return window.Tukey{Alpha: p}.Transform },
		ok:   func(p float64) bool { return !math.IsNaN(p) },
	},
}

type builtin struct {
	name func(float64) string
	fn   func(float64) func([]float64) []float64
	ok   func(float64) bool
}

func main() {
	name := flag.String("window", "", "specify a built-in window name (required if csv not given)")
	param := flag.Float64("param", math.NaN(), "specify parameter for parametric windows")
	symm := flag.Bool("symm", true, "specify whether the window is symmetric")
	n := flag.Int("n", 20, "specify window length")
	m := flag.Int("m", 1000, "specify sample length (must be greater than n)")
	csv := flag.String("csv", "", "specify an input file of dB transformed frequency amplitudes (required if window not given)")
	out := flag.String("o", "", "specify output file for plot (required, formats eps, jpg, jpeg, pdf, png, svg, tex or tif)")
	width := flag.Float64("width", 16, "specify plot width (cm)")
	height := flag.Float64("height", 8, "specify plot height (cm)")
	comment := flag.Bool("comment", false, "output a comment line for the window (only for window function)")
	flag.Parse()

	win := windows[strings.ToLower(*name)]
	if win == nil && *csv == "" {
		fmt.Fprintln(os.Stderr, "missing function name or csv input")
		flag.Usage()
		os.Exit(2)
	}
	if *csv == "" && !win.ok(*param) {
		fmt.Fprintln(os.Stderr, "missing parameter")
		flag.Usage()
		os.Exit(2)
	}
	if *out == "" {
		fmt.Fprintln(os.Stderr, "missing output filename")
		flag.Usage()
		os.Exit(2)
	}

	p := plot.New()
	p.X.Label.Text = "DFT bin"
	p.Y.Label.Text = "Amplitude [dB]"
	p.Add(plotter.NewGrid())

	var (
		c        *characteristics
		spectrum plotter.XYs
		min      float64
		err      error
	)
	if win != nil {
		symmetry := "(symmetric)"
		if !*symm {
			symmetry = "(periodic)"
		}
		p.Title.Text = fmt.Sprintf("Spectral Leakage for %s %s", win.name(*param), symmetry)
		c, spectrum, min, err = funcCharacteristics(win.fn(*param), *n, *m, *symm)
		if err != nil {
			log.Fatal(err)
		}
		if *comment {
			fmt.Printf("// Spectral leakage parameters: ΔF_0 = %2f, ΔF_0.5 = %.2f, K = %.2f, ɣ_max = %2f, β = %2f.\n",
				c.deltaF0, c.deltaFhalf, c.k(), c.gammaMax, c.beta)
		}
	} else {
		f, err := os.Open(*csv)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		p.Title.Text = fmt.Sprintf("Spectral Leakage for %s", *csv)
		c, spectrum, min, err = csvCharacteristics(f, *n, *m)
		if err != nil {
			log.Fatal(err)
		}
	}

	spectrumLines, err := plotter.NewLine(spectrum)
	if err != nil {
		log.Fatalf("spectrum: %v", err)
	}

	gammaLine, err := plotter.NewLine(plotter.XYs{
		{X: 0, Y: c.gammaMax}, {X: float64(*n) / 2, Y: c.gammaMax},
	})
	if err != nil {
		log.Fatalf("ɣ_max: %v", err)
	}
	gammaLine.Color = color.RGBA{R: 0x40, G: 0x80, B: 0xff, A: 0xff}

	deltaF0Line, err := plotter.NewLine(plotter.XYs{
		{X: c.deltaF0 / 2, Y: 0}, {X: c.deltaF0 / 2, Y: min},
	})
	if err != nil {
		log.Fatalf("ΔF_0: %v", err)
	}
	deltaF0Line.Color = color.RGBA{R: 0xff, A: 0xff}

	deltaFhalfLine, err := plotter.NewLine(plotter.XYs{
		{X: c.deltaFhalf / 2, Y: 0}, {X: c.deltaFhalf / 2, Y: min},
	})
	if err != nil {
		log.Fatalf("ΔF_0.5: %v", err)
	}
	deltaFhalfLine.Color = color.RGBA{G: 0xff, A: 0xff}

	var blank plotter.Line

	p.Add(
		gammaLine,
		deltaF0Line,
		deltaFhalfLine,
		spectrumLines,
	)
	p.Legend.Add(fmt.Sprintf("ΔF_0 = %.3v", c.deltaF0), deltaF0Line)
	p.Legend.Add(fmt.Sprintf("ΔF_0.5 = %.3v", c.deltaFhalf), deltaFhalfLine)
	p.Legend.Add(fmt.Sprintf("K = %.3v", c.k()), &blank)
	p.Legend.Add(fmt.Sprintf("ɣ_max = %.3v", c.gammaMax), gammaLine)
	p.Legend.Add(fmt.Sprintf("β = %.3v", c.beta), &blank)
	p.Legend.Top = true
	p.Legend.XOffs = -10
	p.Legend.YOffs = -10

	err = p.Save(vg.Length(*width)*vg.Centimeter, vg.Length(*height)*vg.Centimeter, *out)
	if err != nil {
		log.Fatal(err)
	}
}

type characteristics struct {
	deltaF0    float64
	deltaFhalf float64
	gammaMax   float64
	beta       float64
}

func (c *characteristics) k() float64 { _ = "STUB: not implemented"; return 0 }

func funcCharacteristics(fn func([]float64) []float64, n, m int, symm bool) (c *characteristics, xy plotter.XYs, min float64, err error) {
	_ = "STUB: not implemented"
	return nil, *new(plotter.XYs), 0, nil
}

func csvCharacteristics(r io.Reader, n, m int) (c *characteristics, xy plotter.XYs, min float64, err error) {
	_ = "STUB: not implemented"
	return nil, *new(plotter.XYs), 0, nil
}

func parameters(spectrum []float64, n, m int) (deltaF0, deltaFhalf, gammaMax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func db(m float64) float64 { _ = "STUB: not implemented"; return 0 }
