package camera

import (
    "math"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
)

type Camera struct {
    AspectRatio float64
    FieldOfView float64
    Origin vector3.Vector3

    ImgWidth int
    ImgHeight int

    RotationX []float64 // size 9 (3x3)
    RotationY []float64 // size 9 (3x3)
    RotationZ []float64 // size 9 (3x3)
}

func InitRota(thetaX, thetaY, thetaZ float64) (rotaX, rotaY, rotaZ []float64){
    rotationX := make([]float64, 9)
    rotationY := make([]float64, 9)
    rotationZ := make([]float64, 9)
    // init rotation matrix
    matXtmp := []float64{
        1.0, 0.0, 0.0,
        0.0, math.Cos(thetaX), -math.Sin(thetaX),
        0.0, math.Sin(thetaX), math.Cos(thetaX),
    }

    matYtmp := []float64{
        math.Cos(thetaY), 0.0, math.Sin(thetaY),
        0.0, 1.0, 0.0,
        -math.Sin(thetaY), 0.0, math.Cos(thetaY),
    }

    matZtmp := []float64{
        math.Cos(thetaZ), -math.Sin(thetaZ), 0.0,
        math.Sin(thetaZ), math.Cos(thetaZ), 0.0,
        0.0, 0.0, 1.0,
    }

    for idx, _ := range matXtmp {
        rotationX[idx] = matXtmp[idx]
        rotationY[idx] = matYtmp[idx]
        rotationZ[idx] = matZtmp[idx]
    }
    return rotationX, rotationY, rotationZ
}


func InitCamera(aspectRatio,
                fieldOfView float64,
                imgWidth,
                imgHeight int,
                origin vector3.Vector3,
                thetaX,
                thetaY,
                thetaZ float64) Camera {

    rotationX, rotationY, rotationZ := InitRota(thetaX, thetaY, thetaZ)



    return Camera{
        AspectRatio: aspectRatio,
        FieldOfView: fieldOfView,
        Origin: origin,
        ImgWidth: imgWidth,
        ImgHeight: imgHeight,
        RotationX: rotationX,
        RotationY: rotationY,
        RotationZ: rotationZ,
    }
}

func (c Camera) CreateRay(i, j float64) ray.Ray {
    // position on the point in the 3D world
    var px float64 = (2.0 * (i + 0.5) / float64(c.ImgWidth) - 1.0) * math.Tan(c.FieldOfView * 0.5)
    var py float64 = (1.0 - 2.0 * (j + 0.5) / float64(c.ImgHeight)) * math.Tan(c.FieldOfView * 0.5) * 1.0 / c.AspectRatio
    var pz float64 = -1.0

    // create the ray direction
    rayDir := vector3.InitVector3(px, py, pz)
    rayDir = vector3.UnitVector(rayDir)

    // apply rotation to the ray
    // simulate the camera to be anywhere

    rayDir = vector3.MultMatVec3(c.RotationX, rayDir)
    rayDir = vector3.MultMatVec3(c.RotationY, rayDir)
    rayDir = vector3.MultMatVec3(c.RotationZ, rayDir)

    return ray.InitRay(c.Origin, rayDir)
}
