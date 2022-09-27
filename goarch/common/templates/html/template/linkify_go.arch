package template

import (
	"github.com/mvdan/xurls"
	"html/template"
)

/*Linkify detects the links inside given string and
return a html fragment to directly render in template.

Ex: Suppose that your template like;
...
<p> {.Content}</p>
...
and you want to detect if there is a link inside the Content variable then
convert it to a link element which is in a different color then normal
text and clickable usually.
After you pass that function to the template before parse it
like template.New("ex").Funcs(template.FuncMap{ "linkify": Linkify}),
you can use it in your template like;
...
<p> {linkify .Content}
...

*/
func Linkify(s string) template.HTML {
	s = template.HTMLEscapeString(s)
	return template.HTML(xurls.Relaxed.
		ReplaceAllString(s, `<a href ="$0" target="_blank" rel="noopener noreferror">$0</a>`))
}
