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

### 7) Pixel Canvas Structure

- Create pxcanvase.go

### 8) Pixel Canvas Renderer

- Create pxcanvasrenderer.go
- Implementing pxcanvasrenderer in pxcanvas.go

### 9) Pixel Canvas Layout

- Implementing layout in types.go, pixl.go and layout.go

### 10) Panning & Zooming

- Create ops.go file to operating mouse operation
- Create mouse.go file to move or scroll the mouse
- Implementing package mouse movement in pxcanvas.go

### 11) Painting Pixels

- Create interface Brushable in apptype to be able set the color
- Create new function and implementing interface brushable in pxcanvas
- Create brush.go file
- Setup mouse movement in mouse.go

### 12) Cursor Display

- Create Cursor function in brush.go file
- Implementing canvas cursor function in pxcanvasrenderer.go
- Setup cursor in mouse.go file

### 13) Creating New Images

- New function for new file or new drawing
- Create Menus ui (menus.go)
- Implementing menus.go in layout.go

### 14) Saving Images

- Create save as images

### 15) Loading Images

- Create Util.go to to get image color function
- Implementing Load image in menus.go

## Articles

- [Updating MSYS2](https://www.msys2.org/docs/updating/)
