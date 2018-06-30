# scarlet

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Travis build][travis-image]][travis-url]
[![Code coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]

Generates CSS from `.scarlet` files. Very similar to Stylus, but with higher compression.

## Basic usage

```scarlet
body
	color black
	font-size 100%
	padding 1rem
```

## State

```scarlet
a
	color blue

	:hover
		color red
```

## Classes

```scarlet
a
	color blue

	// "active" elements inside a link
	.active
		color red

	// links that have the "active" class
	&.active
		color red
```

## Multiple selectors

```scarlet
// All in one line
h1, h2, h3
	color orange

// Split over multiple lines
h4,
h5,
h6
	color purple
```

## Variables

```scarlet
text-color = black
transition-speed = 200ms

body
	font-size 100%
	color text-color

a
	color blue
	transition color transition-speed ease
	
	:hover
		color red
```

## Mixins

```scarlet
mixin horizontal
	display flex
	flex-direction row

mixin vertical
	display flex
	flex-direction column
```

## Animations

```scarlet
animation rotate
	0%
		transform rotateZ(0)
	100%
		transform rotateZ(360deg)

animation pulse
	0%, 100%
		transform scale3D(0.4, 0.4, 0.4)
	50%
		transform scale3D(0.9, 0.9, 0.9)
```

## Quick media queries

```scarlet
body
	vertical

> 800px
	body
		horizontal
```

## Author

| [![Eduard Urbach on Twitter](https://gravatar.com/avatar/16ed4d41a5f244d1b10de1b791657989?s=70)](https://twitter.com/eduardurbach "Follow @eduardurbach on Twitter") |
|---|
| [Eduard Urbach](https://eduardurbach.com) |

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
