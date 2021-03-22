package render

import (
    //"volumetric-cloud/vector3"
    "volumetric-cloud/img"

    "sync"
)

func Render(imgSizeY, imgSizeX int) img.Img {
    image := img.InitImg(imgSizeY, imgSizeX)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY * imgSizeX)

    for i := 0; i < imgSizeY; i += 1 {
        for j := 0; j < imgSizeX; j += 1 {
            go renderPixel(image, i, j, &wg)
            //image.SetPixel(i, j, 200, 100, 20)
        }
    }

    return image
}

func renderPixel(image img.Img, i, j int, wg *sync.WaitGroup) {
    image.SetPixel(i, j, 200, 100, 20)
    wg.Done()
}
