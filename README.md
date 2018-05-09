# Genetic Al-go-rithm

Generates random abstract images using mutation to maintain genetic diversity from one generation to the next. Written in Go / Golang and SDL2. Abstract syntax tree outputs an RBG image and valid Scheme / Lisp code.

## Example

```
R: ( SimplexNoise ( Lerp ( + X X ) X ( Sine ( Sine X ) ) ) ( + ( Atan ( * ( + Y X ) Y ) ) ( / ( + Y X ) ( Sine X ) ) ) )
X ) Y ) ) ( / ( + Y X ) ( Sine X ) ) ) ) 

G: ( Atan2 ( Atan ( Atan ( Atan ( Sine ( Atan ( Cos ( * X ( Lerp ( Lerp ( + X X
) -0.679606915 Y ) -0.936428785 0.151277661 ) ) ) ) ) ) ) ) ( Sine ( + ( - ( + YX ( Lerp -0.311703146 0.281130552 Y ) ) ( * -0.895853817 -0.415774405 ) ) ( Lerp 0.134992719 Y -0.430226684 ) ( Atan ( + Y ( Cos ( -0.452641904 ) ( * ( Atan -0.725865364 ) Y ) ) ( + Y ( Cos Y ) ) ) ) )

B: ( Sine ( * ( * ( * ( Atan X ) ( Atan X ) ) X ) ( Lerp ( - ( Lerp 0.332089782 .349048853 ) ( Sine ( Atan2 ( + ( Sine X ) X ) Y ) ) ) ) ) ( Atan2 ( + 0.521236300 ( + -0.462310731 ( * Y 0.732817292 ) ) ) ( / X (X ( Lerp -0.311703146 0.281130552 Y ) ) ( * -0.895853817 -0.415774405 ) ) ( Lerp 0.134992719 Y -0.430226684 ) ( Atan ( + Y ( Cos ( Lerp 0.142749071 0.506971121  ) X ) ( + ( - Y -0.201029778 ) ( Lerp -0.857006669 X -0.032399476 ) ) ) ( Atan ( + ( Cos ( Atan ( Lerp Y -0.403355420 ( / -0.60870Y ) ) ) ) ) ) )  
```

![example](https://i.imgur.com/J46z5nc.gif)

### Prerequisites 

Requires the SDL2 binding for Go.

```
go get -v github.com/scottferg/Go-SDL2/sdl
```

### References

Episodes 15 - 18 of Games With Go by Jack Mott (https://gameswithgo.org/).

SimplexNoise1234 by Stefan Gustavson, ported to Go by Jack Mott (https://github.com/stegu/perlin-noise).

SDL2 binding for Go by Ve & Co (https://github.com/veandco/go-sdl2).