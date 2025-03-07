# QUAD

## README.md

### UPDATE :

`quad.go`: `firstEdge` and `lastEdge` replaced by `leftCorner` and `rightCorner` for more readability.

`quad.go`: A condition have been added to filter better when `x`or `y` are null in order to display nothing.

### procédure de test:

`main.go :`

```go
package main

import "piscine"

func main() {
	piscine.QuadA(5, 3)
	//piscine.QuadA(5, 1)
	//piscine.QuadA(1, 1)
	//piscine.QuadA(1, 5)
	//piscine.QuadB(5, 3)
	//piscine.QuadB(5, 1)
	//piscine.QuadB(1, 1)
	//piscine.QuadB(1, 5)
	//piscine.QuadC(5, 3)
	//piscine.QuadC(5, 1)
	//piscine.QuadC(1, 1)
	//piscine.QuadC(1, 5)
	//piscine.QuadD(5, 3)
	//piscine.QuadD(5, 1)
	//piscine.QuadD(1, 1)
	//piscine.QuadD(1, 5)
	//piscine.QuadE(5, 3)
	//piscine.QuadE(5, 1)
	//piscine.QuadE(1, 1)
	//piscine.QuadE(1, 5)
}
```

Dans le dossier du `main.go`:

```bash
go mod init (nom du dossier)

& toujours dans le dossier du main.go :

go work init (dans le dossier du main.go)
go work use .
go work use ../quad/ ou go work use ./quad
```

## Éxplication du code:

### ÉTAPE 1 :

```go
func main() {
	piscine.QuadB(5, 3)
}
```

package `piscine` => la fonction `QuadB`.
`5` et `3` => `x` et `y` (en arguments)

### ETAPE 2 :

```go
func QuadB(x, y int) {
	// print first line
	PrintLine(x, '/', 92, '*')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, '*')
	}
	// print last line
	if y > 1 {
		PrintLine(x, 92, '/', '*')
	}
}
```

`PrintLine` => une fonction qui récupère `x`(int)=5 `firstEdge`(rune) = '/' `lastEdge`(rune) = 92(\) `inter`(rune) = '*'

### ETAPE 3 :

```go
func PrintLine(x int, firstEdge, lastEdge, inter rune) {
	z01.PrintRune(firstEdge)
	if x > 1 {
		for i := 0; i < x-2; i++ {
			z01.PrintRune(inter)
		}
		z01.PrintRune(lastEdge)
	}
	z01.PrintRune('\n')
}
```

`z01.PrintRune(firstEdge)` => coin gauche

si `x` = 1 pas besoin des intermediaires ni du coin droit

boucle qui commence avec `i` = 0 et qui tourne tant que `i` < `x-2`

pourquoi `x-2` ? parcequ'on ne compte pas le coin gauche et le coin droit
`lastedge` => coin droit
`\n` => retour à la LIGNE

### ETAPE 4 :

```go
for i := 0; i < y-2; i++ {
	PrintIntermediateLine(x, '*')
}
```

boucle qui commence avec `i` = 0 et qui tourne tant que `i` < `y-2`
pourquoi `y-2` ? on ne compte pas le coté gauche et coté droit.

si `y` = 3 combien de fois se répète `PrintIntermediateLine` ? pour `y` = 3 une fois.

`x` = 5 `edge` = '*' =====> `PrintIntermediateLine`

### ETAPE 5 :

```go
func PrintIntermediateLine(x int, edge rune) {
	for i := 0; i <= x-1; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(edge)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
```

on récupère `x`=5 et `edge` = '*'

la boucle commence par `i` = 0 et elle continue tant que `i` <= `x-1`
pourquoi `x-1` ? on ne compte pas le côté gauche ni le côté droit

si `i` = 0 ====> imprime le caractere `edge`(*)
quand `i` n'est plus égal à 0 on imprime des espaces
et à un moment donné `i` == `x-1` donc on imprime le côté droit.
puis retour à la LIGNE

### ETAPE 6 :

```go
// print last line
	if y > 1 {
		PrintLine(x, 92, '/', '*')
	}
```

imprime la derniere LIGNE exactement de la même façon que la premiere ligne avec comme arguments
`x`= 5 `firstEdge` = 92(\) `lastEdge` = '/' `inter` = *

la seule chose qui change par rapport à la première ligne c'est une condition supplémentaire

j'imprime la derniere ligne que si `y` > 1
pourquoi ? pourquoi ne pas imprimer si `y` = 1 ?
si `y` = 1 la premiere ligne suffit donc pas de côtés et pas de dernière ligne.