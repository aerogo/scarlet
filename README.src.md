# {name}

{go:header}

Generates CSS from `.scarlet` files. Very similar to Stylus, but with higher compression.

{go:install}

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

{go:footer}
