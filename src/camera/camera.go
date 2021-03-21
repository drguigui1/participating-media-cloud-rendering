package camera

import (
    "math"

    "volumetric-cloud/vector3"
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

func InitCamera(aspectRatio,
                fieldOfView float64,
                imgWidth,
                imgHeight int,
                origin vector3.Vector3,
                thetaX,
                thetaY,
                thetaZ float64) Camera {

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
