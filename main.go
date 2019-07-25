package main

import (
    "flag"
    "github.com/veandco/go-sdl2/img"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"
    "log"
    "path/filepath"
)

var WorkDir = flag.String("dir", "-", "Working directory for all medias")

func sample01(wdir string) error {
    var mainImg string = "seagull.jpg"
    var secImg string = "wmark.png"

    mainImgPath := filepath.Join(wdir, "src", mainImg)
    mainSurf, err := img.Load(mainImgPath)
    if err != nil {
        return err
    }

    secImgPath := filepath.Join(wdir, "src", secImg)
    secSurf, err := img.Load(secImgPath)
    if err != nil {
        return err
    }

    rdr, _ := sdl.CreateSoftwareRenderer(mainSurf)
    secTex, _ := rdr.CreateTextureFromSurface(secSurf)

    rdr.Copy(
        secTex,                                // WHO
        &sdl.Rect{0, 0, secSurf.W, secSurf.H}, // SRC
        &sdl.Rect{0, 0, 60, 60},               // DEST
    )

    distPath := filepath.Join(wdir, "dist", "sample01.png")
    err = img.SavePNG(mainSurf, distPath)

    return err
}

func sample02(wdir string) error {
    var mainImg string = "seagull.jpg"
    var msgFont string = "DroidSans.ttf"

    mainImgPath := filepath.Join(wdir, "src", mainImg)
    mainSurf, err := img.Load(mainImgPath)
    if err != nil {
        return err
    }

    fontPath := filepath.Join(wdir, "src", "fonts", msgFont)
    font, err := ttf.OpenFont(fontPath, 32)
    if err != nil {
        return err
    }

    tsurf, _ := font.RenderUTF8Blended("Hello TTF!!!", sdl.Color{255, 255, 0, 255})

    rdr, _ := sdl.CreateSoftwareRenderer(mainSurf)
    ttex, _ := rdr.CreateTextureFromSurface(tsurf)

    //centerX := (mainSurf.W / 2) - (tsurf.W / 2)
    //centerY := (mainSurf.H / 2) - (tsurf.H / 2)

    ntimesY := int(mainSurf.H/tsurf.H) - 1
    ntimesX := int(mainSurf.W/tsurf.W) - 1

    for i := 1; i < ntimesY; i++ {
        for j := 1; j < ntimesX; j++ {
            rdr.Copy(
                ttex,
                &sdl.Rect{0, 0, tsurf.W, tsurf.H},
                &sdl.Rect{int32(j) * tsurf.W, int32(i) * tsurf.H, tsurf.W, tsurf.H},
            )
        }
    }

    distPath := filepath.Join(wdir, "dist", "sample02.png")
    err = img.SavePNG(mainSurf, distPath)

    return err
}

func main() {
    flag.Parse()
    wdir := *WorkDir
    ttf.Init()
    if err := sample01(wdir); err != nil {
        log.Fatalf("Failed on sample01: %+v", err)
    }
    if err := sample02(wdir); err != nil {
        log.Fatalf("Failed on sample02: %+v", err)
    }
}
