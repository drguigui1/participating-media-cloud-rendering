package render

import (
    "volumetric-cloud/vector3"
)

// data to pass through the chan
type  PixelInfo struct {
    I int
    J int
    Color vector3.Vector3
}

// Return the color of the pixel to render
func RenderPixel(i, j int, c chan PixelInfo, wg *sync.WaitGroup) {
    data := PixelInfo{
        I: i,
        J: j,
        Color: vector3.InitVector3(1.0, 0.2, 0.2),
    }
    c <- data
    wg.Done()
}
