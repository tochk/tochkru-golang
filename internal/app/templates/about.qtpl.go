// Code generated by qtc from "about.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line about.qtpl:1
package templates

//line about.qtpl:1
import "tochkru-golang/internal/app/datastruct"

//line about.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line about.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line about.qtpl:2
func StreamAboutPage(qw422016 *qt422016.Writer, topTags []datastruct.Tags) {
//line about.qtpl:2
	qw422016.N().S(`
`)
//line about.qtpl:3
	qw422016.N().S(Header("About"))
//line about.qtpl:3
	qw422016.N().S(`
`)
//line about.qtpl:4
	qw422016.N().S(Menu(topTags))
//line about.qtpl:4
	qw422016.N().S(`
<article id='block_news'>
Aleksei Andreev </br>
Software engineer at Vio.com </br>
</article>
`)
//line about.qtpl:9
	qw422016.N().S(Footer())
//line about.qtpl:9
	qw422016.N().S(`
`)
//line about.qtpl:10
}

//line about.qtpl:10
func WriteAboutPage(qq422016 qtio422016.Writer, topTags []datastruct.Tags) {
//line about.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line about.qtpl:10
	StreamAboutPage(qw422016, topTags)
//line about.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line about.qtpl:10
}

//line about.qtpl:10
func AboutPage(topTags []datastruct.Tags) string {
//line about.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line about.qtpl:10
	WriteAboutPage(qb422016, topTags)
//line about.qtpl:10
	qs422016 := string(qb422016.B)
//line about.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line about.qtpl:10
	return qs422016
//line about.qtpl:10
}