# Dungeons of Go
A roguelike for the 2025 RoguelikeDev Does The Complete Roguelike Tutorial.

This roguelike is built as an exercise in a couple of areas:
* Building a full-fledge roguelike using the Go programming language
* Leveraging Raylib (v5.5) instead of libtcod
* Providing support from the start for both desktop and web builds


## Why Go?
Recently, I've taken an interest in Go for a few reasons - some professional, some personal. For the purposes of this tutorial session, however, I wanted to build the tutorial using a new language that typically doesn't see a lot of traction. Looking at previous years entries (https://old.reddit.com/r/roguelikedev/wiki/python_tutorial_series), you see a lot of Python, Rust, C, C#, etc... But there are very few Go entries.

While Go isn't a particularly popular language for game development (the biggest knocks against it -- as far as I know -- being garbage collection and poor library support), I find Go to be a good blend of easy-to-understand and performant. For something as 'trivial' as this roguelike, I'm not too concerned about garbage collection. Not only are there popular games utilizing other garbage-collected languages like C# and Java, but the overhead of garbage collection on a smaller 2D game seems minimal. As for the library support ...

## Why Raylib?
Raylib is a library I've been very interested in for a while. Similar to my attraction to Go as a simpler language to wrap one's head around, Raylib offers a nice blend of lower-level control (as opposed to using an engine like Godot, Unity, or Unreal) to build something tailor-made to my vision, while also not being _as_ in the weeds as something like SDL. The fact that the library has numerous bindings (https://github.com/raysan5/raylib/blob/master/BINDINGS.md) makes it easier to pick up the library in one language (such as Python) and apply many of the same approaches in another language. It also offers 2D and 3D support; this isn't necessarily relevant for this project, but I'd like to apply the work done here to other projects in the future -- not having to learn a whole new library to move from 2D to 3D is an intriguing (if ultimately impractical) idea for me. Also, as of a couple months ago, some incredible WASM bindings were released for raylib go, which dovetails perfectly into my final point.

## Why Desktop AND Web?
Having started participating in game jams this year, I quickly realized that a web build made "getting your name out there" a lot easier -- sometimes, folks just want to press play. Especially when you're an unknown quantity, having an environment where one can test a game without downloading or compiling anything lowers a critical barrier to entry. I've also found myself enjoying web-based roguelikes on itch.io from time to time. While I don't think web is necessarily the platform to target for something particularly deep or complex (this is just a gut feeling, not something based in facts or research), providing folks following along with this tutorial a way to play the game at various stages is something that interests me as well. With Go and Raylib, I have the tools available to make minimal changes to my project and build for multiple deployments, which is incredibly intriguing to me.

## Things I'll be using
I'd like to give credit where credit is due, and provide attribution for everything I'll be using to build out this game. In no particular order:

### Libraries:
* Raylib-Go: https://github.com/gen2brain/raylib-go - Go bindings for the Raylib library, pnned (at time of writing) to version 5.5
* Raylib-Go-Wasm: https://github.com/BrownNPC/Raylib-Go-Wasm - Bindings that allow Raylib-Go to work with the web, utilizing WASM and very few constraints and changes to the code
* RogueYun Agm Edit Dwarf Fortress character set: https://dwarffortresswiki.org/Tileset_repository#16x16-RogueYun-AgmEdit.png - In keeping with the ASCII asthetic, I'll be using this spritesheet as a template for the tutorial
* The libtcod tutorial: https://rogueliketutorials.com/tutorials/tcod/v2/ - the "source material" for this project, I'll be following along with the libtcod tutorial and adapting as appropriate (given the library _and_ the language will be changing)

## Following along
My goal is to create a distinct branch for each chapter of the tutorial, so that those who are following along can see each final checkpoint along the way. 

I also plan to document my journey at https://github.com/old-pops-mcgee/dungeons-of-go/wiki - with a wiki page for each chapter, so that those following along don't need to do the same mental calculus to translate the original wiki into Go

## Installation instructions
I've included a Makefile to easily build both the desktop and web versions of the game

### For Desktop
From the root directory, run `make build`. This creates a desktop binary (by default, for linux/amd64 - you can edit the Makefile to target different architectures) called `app`. Run `app` and you have the game!

### For Web
For this, you must first ensure you've installed Raylib-Go-Wasm from https://github.com/BrownNPC/Raylib-Go-Wasm . I have not included my local clone of the package, but the makefile assumes the clone is present at the root. Once the clone is present, run the following _once_:
`cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./Raylib-Go-Wasm/index/wasm_exec.js`
`go build ./Raylib-Go-Wasm/server/server.go`

Next, run `make build-web`. This will set up the web application as a WASM binary, and put it in the right location within `Raylib-Go-Wasm`. With that done, run `./server` - your game should be installed and running off a local web server on localhost:8080.
