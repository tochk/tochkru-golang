// Code generated by qtc from "main.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line main.qtpl:1
package templates

//line main.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line main.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line main.qtpl:1
func StreamHeader(qw422016 *qt422016.Writer) {
//line main.qtpl:1
	qw422016.N().S(`
head
`)
//line main.qtpl:3
}

//line main.qtpl:3
func WriteHeader(qq422016 qtio422016.Writer) {
//line main.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:3
	StreamHeader(qw422016)
//line main.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:3
}

//line main.qtpl:3
func Header() string {
//line main.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:3
	WriteHeader(qb422016)
//line main.qtpl:3
	qs422016 := string(qb422016.B)
//line main.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:3
	return qs422016
//line main.qtpl:3
}
