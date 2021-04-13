package animations

import (
    "fmt"
    "math"
    "strconv"
    "volumetric-cloud/camera"
    "volumetric-cloud/ray"
    "volumetric-cloud/vector3"
    "volumetric-cloud/scene"
)

func AnimRotation(center vector3.Vector3, radius float64,
    imgSizeX, imgSizeY, picNumber, nbRaysPerPixel int, s scene.Scene){

    s.Camera.Origin.X = 0.0
    s.Camera.Origin.Y = 0.0
    s.Camera.Origin.Z = 0.0
    s.Camera.RotationX, s.Camera.RotationY, s.Camera.RotationZ =
        camera.InitRota(0.0, 0.0, 0.0)
    for i := 0; i < picNumber + 1; i+=1 {
        teta := (float64(i) * 2 * math.Pi) / float64(picNumber)
        s.Camera.Origin.X = -radius * math.Sin(teta) + center.X
        s.Camera.Origin.Y = center.Y
        s.Camera.Origin.Z = radius * math.Cos(teta) + center.Z


        s.Camera.RotationX, s.Camera.RotationY, s.Camera.RotationZ =
                    camera.InitRota(0.0, - teta, 0.0)

        image := s.Render(imgSizeX, imgSizeY, nbRaysPerPixel)
        image.SavePNG("videos/video_img" + strconv.Itoa(i) + ".png")
        s.Pixels = s.Pixels[:0]
        fmt.Println("---- img" + strconv.Itoa(i) + "---- done")
    }
}

func AnimTranslate(ray ray.Ray, picNumber, imgSizeX, imgSizeY, nbRaysPerPixel int,
    step float64, s scene.Scene, cam camera.Camera){

    for i:=0; i < picNumber; i++ {
        newPos := ray.RayAt(step * float64(i))
        s.Camera.Origin.X = newPos.X
        s.Camera.Origin.Y = newPos.Y
        s.Camera.Origin.Z = newPos.Z

        image := s.Render(imgSizeX, imgSizeY, nbRaysPerPixel)
        image.SavePNG("videos/video_img" + strconv.Itoa(i) + ".png")
        s.Pixels = s.Pixels[:0]
        fmt.Println("---- img " + strconv.Itoa(i) + "---- done")
    }
}






