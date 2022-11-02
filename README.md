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

### Introduction & Project Setup

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

## Articles

- [Updating MSYS2](https://www.msys2.org/docs/updating/)
