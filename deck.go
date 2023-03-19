package cotalk

import (
	"fmt"
	"strings"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

// newDeck returns a deck with default styling and navigation on bottom
func newDeck() *deck {
	return &deck{
		Styles: []*CSS{},
	}
}

type deck struct {
	Title  string // header title
	Slides []*Element
	Styles []*CSS // first one is the deck default styling
}

// Slide appends a new slide to the deck. elements can be anything
// supported by the web package.
func (d *deck) Slide(elements ...interface{}) {
	d.Slides = append(d.Slides, Wrap(elements...))
}

// Page returns a web page ready for use.
func (d *deck) Page() *web.Page {
	styles := Style()
	for _, s := range d.Styles {
		styles.With(s)
	}
	body := Body()
	for i, content := range d.Slides {
		j := i + 1
		id := fmt.Sprintf("%v", j)
		slide := Div(Class("slide"), Id(id))

		slide.With(content)
		body.With(slide)
	}

	return NewFile("index.html",
		Html(
			Head(
				Title(d.Title),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				styles,
			),
			body,
		),
	)
}

type navbar struct {
	max     int // number of slides
	current int // current slide
}

// BuildElement is called at time of rendering
func (b *navbar) BuildElement() *Element {
	ul := Ul()
	for i := 0; i < b.max; i++ {
		j := i + 1
		hash := fmt.Sprintf("#%v", j)
		li := Li(A(Href(hash), j))
		if j == b.current {
			li.With(Class("current"))
		}
		ul.With(li)
	}
	b.current++
	return Nav(ul)
}

type algorithms struct {
	current int // current slide
}

func (a *algorithms) BuildElement() *Element {
	name := fmt.Sprintf("Alg%v", a.current)
	a.current++
	return Wrap(
		loadFunc("../alg.go", name),
		shell(
			fmt.Sprintf("$ go test -benchmem -run=Benchmark%s .", name),
			fmt.Sprintf("testdata/%s_bench.html", strings.ToLower(name)),
		),
	)
}
