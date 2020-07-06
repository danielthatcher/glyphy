# glyphy
Generate unicode homoglyphs for a given string.

Example usage:
```
$ glyphy example.com

example.ᶜom
example.ℂom
example.ℭom
example.Ⅽom
example.ⅽom
example.Ⓒom
example.ⓒom
example.Ｃom
example.ｃom
example.𝐂om
example.𝐜om
example.𝐶om
example.𝑐om
example.𝑪om
example.𝒄om
example.𝒞om
example.𝒸om
example.𝓒om
example.𝓬om
example.𝔠om
example.𝕔om
example.𝕮om
example.𝖈om
example.𝖢om
example.𝖼om
example.𝗖om
example.𝗰om
example.𝘊om
example.𝘤om
example.𝘾om
example.𝙘om
example.𝙲om
example.𝚌om
example.🄫om
example.🄲om
exampˡe.com
exampᴸe.com
exampₗe.com
exampℒe.com
exampℓe.com
exampⅬe.com
exampⅼe.com
...
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
