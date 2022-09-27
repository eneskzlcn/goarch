This directory contains common ui elements
like reusable components and layout elements which
are common for every page.

Suppose that you want to render home.gohtml page
then actually the html content of home.gohtml
taken by template engine and get put into layout.gohtml
file which expects a template "content". Every page
you want to parse and render, the same thing continually happens.
