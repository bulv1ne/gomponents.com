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
			P(Class("lead"), g.Text(`gomponents really shines if we build reusable components together. 🌟`)),
			P(g.Raw(`Have you made a project with gomponents that other developers can use? <a href="https://github.com/maragudk/gomponents.com/issues/new">Create an issue on Github</a> and we can work together to add it here. 😎`)),

			H2(g.Text(`gomponents + TailwindCSS`)),
			P(g.Raw(`gomponents works great together with <a href="https://tailwindcss.com">TailwindCSS</a>. In fact, this page is built using gomponents and TailwindCSS! Check out <a href="https://github.com/maragudk/gomponents.com">the source of this page</a> or <a href="https://github.com/maragudk/gomponents-tailwind-example">a gomponents + TailwindCSS example project</a>.`)),

			H2(g.Text(`gomponents + Heroicons`)),
			P(g.Raw(`<a href="https://heroicons.com">Heroicons</a> are a collection of handcrafted SVG icons, by the makers of TailwindCSS. They work great together with gomponents! Check out <a href="https://github.com/maragudk/gomponents-heroicons">the gomponents-heroicons library on Github</a>.`)),
		),
	)
}
