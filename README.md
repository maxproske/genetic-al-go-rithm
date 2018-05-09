# Genetic Al-go-rithm

Generates random abstract images using mutation to maintain genetic diversity from one generation to the next. Written in Go / Golang and SDL2.

## Example

Abstract syntax tree outputs valid Scheme / Lisp code and an RBG image.

```
R: ( Cos ( / ( SimplexNoise ( Atan2 -0.698849916 0.478828907 ) ( Atan ( Cos X ) ) ) ( Sine ( - ( SimplexNoise Y ( Cos X ) ) ( Atan2 0.384923816 -0.703229308 ) ) ) ) )
G: ( Sine ( + ( Atan2 ( + ( / X -0.311322391 ) X ) ( Cos ( * 0.309438229 Y ) ) ) ( Atan2 ( Atan ( Atan Y ) ) ( Atan2 X ( Atan2 ( Cos X ) Y ) ) ) ) )
B: ( SimplexNoise ( Sine ( + ( Cos X ) ( Atan2 -0.504971743 X ) ) ) ( + Y X ) )
```

![example](https://i.imgur.com/oMrVQnm.png)

### Prerequisites 

Requires the SDL2 binding for Go.

```
go get -v github.com/scottferg/Go-SDL2/sdl
```

### References

Episodes 15 - 18 of Games With Go by Jack Mott (https://gameswithgo.org/).

SimplexNoise1234 by Stefan Gustavson, ported to Go by Jack Mott (https://github.com/stegu/perlin-noise).

SDL2 binding for Go by Ve & Co (https://github.com/veandco/go-sdl2).