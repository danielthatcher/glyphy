# glyphy
Generate unicode homoglyphs for a given string.

Example usage:
```
$ glyphy example.com
exampˡe.com
exampℓe.com
exampⅼe.com
exampⓛe.com
exampｌe.com
example․com
example﹒com
example．com
example.cºm
example.cᵒm
example.cₒm
example.cℴm
example.cⓞm
example.cｏm
eˣample.com
eₓample.com
eⅹample.com
eⓧample.com
eｘample.com
ᵉxample.com
ₑxample.com
ℯxample.com
ⅇxample.com
ⓔxample.com
ｅxample.com
examplᵉ.com
examplₑ.com
examplℯ.com
examplⅇ.com
examplⓔ.com
examplｅ.com
example.ᶜom
example.ⅽom
example.ⓒom
example.ｃom
examᵖle.com
examⓟle.com
examｐle.com
exaᵐple.com
exaⅿple.com
exaⓜple.com
exaｍple.com
example.coᵐ
example.coⅿ
example.coⓜ
example.coｍ
exªmple.com
exᵃmple.com
exₐmple.com
exⓐmple.com
exａmple.com
```

Help text:
```
Usage of glyphy:
  -n, --max-replacements int   The maximum number of positions to replace in (default 1)
  -u, --urlencode              URL encode special characters in the data
```

## Installation
If you have Go installed, you can run
```
go get github.com/danielthatcher/glyphy
```

## Building
This project uses [packr2](https://github.com/gobuffalo/packr/tree/master/v2). Once you have installed packr2, from inside the project's directory run
```
packr2
```
Then you can build or install the project with `go build` or `go install` respectively.

## Sources
The `replacements.json` file which defines the replacments glyphy uses is build using the following sources:
* https://appcheck-ng.com/wp-content/uploads/unicode_normalization.html.
* Testing which unicode codepoints get replaced with ASCII character when using in domain names for JavaScript's `URL` constructor in Chrome and Firefox.
