# scarlet
Generates CSS from `.scarlet` files.

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
