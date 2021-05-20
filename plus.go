package main

import (
	_ "embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func PlusPage() g.Node {
	return Page(
		"gomponents +",
		"gomponents really shines if we start building reusable components together. Add yours here!",
		"/plus/",

		Div(
			H1(g.Text("gomponents +")),
			P(g.Raw(`gomponents really shines if we start building reusable components together. 🌟 Have you made a project with gomponents that other developers can use? <a href="https://github.com/maragudk/gomponents.com/issues/new">Create an issue on Github</a> and we can work together to add it here. 😎`)),

			H2(g.Text(`gomponents + TailwindCSS`)),
			P(g.Raw(`gomponents works great together with <a href="https://tailwindcss.com">TailwindCSS</a>. In fact, this page is built using gomponents and TailwindCSS! Check out <a href="https://github.com/maragudk/gomponents.com">the source of this page</a> or <a href="https://github.com/maragudk/gomponents-tailwind-example">a gomponents + TailwindCSS example project</a>.`)),
		),
	)
}
