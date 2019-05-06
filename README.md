# scarlet

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

Generates CSS from `.scarlet` files. Very similar to Stylus, but with higher compression.

## Installation

```shell
go get -u github.com/aerogo/scarlet/...
```

## CLI

If you're looking for the official compiler, please install [pack](https://github.com/aerogo/pack).

The CLI tool included in this repo offers a check to see if your classes are referenced or not via `scarlet -check`.

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

Mixins can be used like this:

```scarlet
#sidebar
	vertical
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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://www.patreon.com/eduardurbach)

[godoc-image]: https://godoc.org/github.com/aerogo/scarlet?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/scarlet
[report-image]: https://goreportcard.com/badge/github.com/aerogo/scarlet
[report-url]: https://goreportcard.com/report/github.com/aerogo/scarlet
[tests-image]: https://cloud.drone.io/api/badges/aerogo/scarlet/status.svg
[tests-url]: https://cloud.drone.io/aerogo/scarlet
[coverage-image]: https://codecov.io/gh/aerogo/scarlet/graph/badge.svg
[coverage-url]: https://codecov.io/gh/aerogo/scarlet
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
