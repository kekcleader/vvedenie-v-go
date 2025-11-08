# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

Project overview
- This repo builds a Russian-language Go primer as a LaTeX book. The main entrypoint is vvedenie_v_go.tex, which includes chapters from text/text.tex. Source code examples are rendered with the minted package (custom gocode environment).
- The exercises/ folder contains standalone Go snippets and one runnable demo (exercises/r.go); it is not wired into the LaTeX build.

Common commands
- Build PDF (XeLaTeX + minted):
  ```sh
  make
  ```
- Direct XeLaTeX invocation (same flags as Makefile):
  ```sh
  xelatex -shell-escape -interaction=nonstopmode vvedenie_v_go.tex
  ```
- Clean build artifacts:
  ```sh
  make clean    # aux/log/toc/minted caches
  make cleanall # also removes vvedenie_v_go.pdf
  ```
- Run the runnable Go demo:
  ```sh
  go run ./exercises/r.go
  ```

Build requirements (from README highlights)
- LaTeX with XeLaTeX and minted (Pygments) enabled; -shell-escape is required by minted.
- Fonts expected in fonts/:
  - Literaturnaja.otf, Literaturnaja-Bold.otf, Literaturnaja-Italic.otf
  - Akademicheskaja.ttf
- On macOS, MacTeX works; on Windows, TeXworks + MiKTeX; on Debian/Ubuntu, a full TeX Live installation is sufficient.

Architecture and structure
- LaTeX book:
  - vvedenie_v_go.tex sets typography (memoir class), fonts, minted style, page layout, and includes text/text.tex.
  - minted style is set to borland; a custom gocode environment formats embedded Go listings with line numbers and frames.
- Go examples:
  - exercises/r.go is a complete program illustrating slice headers (reflect/unsafe) and can be run directly.
  - exercises/slices.go contains illustrative fragments intended for the book; it is not a compilable program.

Linting and tests
- No lint or test tooling is configured in this repo. There are no Go modules or LaTeX lint targets.
