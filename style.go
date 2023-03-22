package cotalk

import . "github.com/gregoryv/web"

func themeOldstyle() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".slide",
		"padding: 10px 20%",
		"text-align: center",
		"height: calc( 100vh - 50px)",
	)
	css.Style("p, pre",
		"text-align: left",
	)
	css.Style(".slide ul",
		"text-align: left",
	)
	// navbar
	css.Style(".slide nav", "text-align: center")
	css.Style(".slide nav ul",
		"list-style-type: none",
		"margin: 0 0",
		"padding: 0 0",
		"text-align: center",
	)
	css.Style("nav ul li",
		"margin: 0 4px",
		"cursor: pointer",
		"display: inline",
	)
	css.Style("nav ul li.current",
		"text-decoration: underline",
	)

	// goish colors
	css.Style("a:link, a:visited",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	css.Style(".header",
		"width: 100%",
		"border-bottom: 1px solid #727272",
		"text-align: right",
		"margin-top: -2em",
		"margin-bottom: 1em",
	)
	css.Style("h1, h2, h3",
		"text-align: center",
	)
	css.Style(".srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"padding-left: 36px",
		"background-color: #fafafa",
		"tab-size: 4",
		"-moz-tab-size: 4",
	)

	css.Style(".filename",
		"float: left",
		"margin-right: 1.6em",
		"margin-top: -1.6em",
	)

	css.Style(".srcfile code",
		"padding: .6em 0 .6em 0",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"background-position: right top",
		"display: block",
		"text-align: left",
	)

	css.Style("nav",
		"float: left",
		"clear: both",
		"width: 100%",
		"display: block",
		"margin-top: 1em",
		"text-align: center",
	)
	css.Style("nav a:link, nav a:visited",
		"color: #727272",
		"padding: 0 5px",
		"margin: 0 2px",
		"text-decoration: none",
	)
	css.Style("nav a.current, nav a:hover",
		"color: black",
		"border-bottom: 1px solid black",
	)
	// toc
	css.Style("li.h3",
		"margin-left: 2em",
	)
	css.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
	)
	css.Style(".dark",
		"background-color: #2e2e34",
		"color: aliceblue",
	)
	css.Style(".light",
		"background-color: #ffffff",
		"color: #3b2616",
	)
	css.Style("img.center",
		"display: block",
		"margin-left: auto",
		"margin-right: auto",
	)
	css.Style("img.left",
		"float: left",
		"margin-right: 2em",
	)
	css.Style(".group",
		"float: left",
		"text-align: left",
		"margin-right: 3em",
	)

	return css
}
