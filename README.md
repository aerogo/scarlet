# scarlet

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Travis build][travis-image]][travis-url]
[![Code coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]

Generates CSS from `.scarlet` files. Very similar to Stylus, but with higher compression.

Example:

```styl
text-color = black
link-color = blue
link-hover-color = red
transition-speed = 200ms

body
	font-size 100%
	color text-color

a
	color link-color
	transition all transition-speed ease
	
	:hover
		color link-hover-color
```

[godoc-image]: https://godoc.org/github.com/aerogo/scarlet?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/scarlet
[goreportcard-image]: https://goreportcard.com/badge/github.com/aerogo/scarlet
[goreportcard-url]: https://goreportcard.com/report/github.com/aerogo/scarlet
[travis-image]: https://travis-ci.org/aerogo/scarlet.svg?branch=master
[travis-url]: https://travis-ci.org/aerogo/scarlet
[codecov-image]: https://codecov.io/gh/aerogo/scarlet/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/aerogo/scarlet
[license-image]: https://img.shields.io/badge/license-MIT-blue.svg
[license-url]: https://github.com/aerogo/scarlet/blob/master/LICENSE
