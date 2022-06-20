# Présentation
https://golang.org


## Origine du langage

.play welcome/hello.go

Rob Pike compile une application C++ : 1h d'attente

Comment réduire les temps de compilations ?

- Eliminer le code compilé inutilement
  - Les imports inutiles
  - Les variables inutilisées

- Simplifier le langage
  - Réduction du nombre d'instructions
  - Une seule manière de faire, explicite
  - Pas d'héritage (mais composition et interface)

- Utiliser plusieurs coeurs. Les procs ne progressent plus beaucoup en
  revanche on en a plusieurs en parallèle.

- Gain en rapidité mais également gain en simplicité, en particulier pour
  les nouveaux programmeurs, nombreux chez Google. Facilité de développement
  d'outils autour du langage. Philosophie Unix.
- Simplicité qui doit permettre de monter à l'échelle sur app gigantesques
  (Google)

Ce ne sont pas les performances pures qui sont recherchées mais
plutôt les performances globales dans le contexte présent.

Langage dit "système" mais vu d'aujourd'hui : réseau, cloud...


## Mise au point

.play welcome/hello2.go
.play welcome/hello2_ex.go

- Rob Pike avec Ken Thomson (Unix) et Robert Griesemer (V8).
- Ils se donnent comme consigne de n'intégrer au langage que ce qui fait
  l'unanimité, le minimum vital.
- Démarrage de manière très concise, langage système pour remplacer le C++
- Ajout de quelques facilités (maps, slice...)
- Garbage collector

- Finalement c'est le côté simple d'utilisation qui prédomine, le langage
  séduit autant les dev Python que C (ce qui n'était pas prévu au départ).
- En 2012 version 1.0 garantie de compatibilité ascendante, même en 2.0 (précisions
  plus tard)
- Réécriture du compilateur en Go lui-même
- Performance pas au détriment de la simplicité (pas de flags au compilateur)
  mais recherche permanente (gains importants sur le GC)

## Particularités

.play welcome/hello2b.go
.play welcome/hellogo.go

- Simplicité.
- Multitache / multicoeur, communication par channel.
- Typage statique.
- Compilé (mais très rapidement).
- Déploiement (linkage statique, cross-compilation).
- Bibliothèque standard très complète, adaptée au web et réseau.

## Installation

- Sans installation : [go.dev/play](https://go.dev/play)
- Installation facile, pas de dépendances [golang.org/doc/install](https://golang.org/doc/install)

Attention : redirection en cours vers le site [go.dev](https://go.dev)

- Nombreux outils en ligne de commande
  - `go run` compile et lance le programme
  - `go build` compile et crée un exécutable autonome
  - `GOOS=windows GOARCH=386 go build` cross compilation
  - `go get` installation de dépendances
  - `go fmt` format esthétique standard
  - etc.

## Editeurs

.play welcome/hello.go

[golang.org/doc/editors.html](https://golang.org/doc/editors.html)

- Vim (avec plugin vim-go et lsp) [doc gopls](https://github.com/golang/tools/blob/master/gopls/doc/vim.md)
- VSCode (avec extension Go)
[vscode](https://code.visualstudio.com/docs/setup/linux)

- GoLand

> `gofmt` permet de formater le code manière standard. A lancer systématiquement lors de l'enregistrement du code.
> Encore mieux `goimports` permet à la fois de formater le code mais aussi d'ajouter ou enlever les importations en
> début de code.

Essayez de formatter le code n'importe comment (est-ce qu'il marche ?), de retirer `import "fmt"` puis utilisez le bouton __Format__

Raccourcis intéressants : `:GoDef` sur Vim `F12` sur VSCode
