## Packages

- [fyne](https://github.com/fyne-io/fyne)

## Development Process

### 1) Install gcc (Windows version)

- Open web browser
- Go to https://www.msys2.org/
- Download the installer
- Run the installer
- Choose default installation folder, klick next and wait until installation process finish
- Now open the mys2 program

- Updating, creating msys2:

```sh
$ pacman -Syu
    // Proceed with installation? [Y/n] >>> Just hit enter
$ packman -Su
    // wait until done
$ pacman -S --needed base-devel mingw-w64-x86_64-toolchain
$ pacman -S --needed base-devel mingw-w64-x86_64-gcc
$ pacman -S --needed base-devel mingw-w64-x86_64-gdb

```

-

### 2) Introduction & Project Setup

- Pixl
  - About
    - Pixel art editor
    - Utilize many feature of Go (Golang)
      - Pointers
      - Function literals / Closures
      - Module / Packages
      - Interfaces
  - Feature
    - New, load, And save files
    - Color picker
    - Swatches
    - Pan / Zoom image
    - Adaptive layout
    - Brush editor
    - Folder structure breakdown:
      - pixl-project
        - apptype
        - pixl (main folder)
        - pxcanvas
        - swatch
        - ui
        - util

### 3) Canvas Overview & State

- Pxcanvas
  - Overview diagram explained:
    [![](https://github.com/Rianto-RNT/pixl/blob/development/doc-assets/pixl-04-spec_002.jpg)](https://github.com/Rianto-RNT)
- Create new go file in apptype folder
  - apptype.go

### 4) Creating a Swatch

- Swatch.go
- Swatchrenderer.go
- Mouse.go

### 5) Swatch Logout & First Run

- Fyne layout diagram explained:
  [![](https://github.com/Rianto-RNT/pixl/blob/development/doc-assets/pixl-01-fyne_005.jpg)](https://github.com/Rianto-RNT)

- Create new file:
  - Layout.go
  - Swatch.go
  - types.go

### 6) Color Picker & App Layout

- Create picker.go file
- Setup color picker in layout.go file

## Articles

- [Updating MSYS2](https://www.msys2.org/docs/updating/)
