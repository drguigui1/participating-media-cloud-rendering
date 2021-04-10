package record

import (
	"fmt"
	"math"
	"strconv"
	"volumetric-cloud/camera"
	"volumetric-cloud/vector3"
	"volumetric-cloud/scene"


)



func ChangeCam(center vector3.Vector3, radius float64,
	imgSizeX, imgSizeY, picNumber, nbRaysPerPixel int, s scene.Scene){
	// x = rcos(teta)
	// y = rsin(teta)
	for i := 0; i < picNumber + 1; i+=1 {

		s.Camera.Origin.X = (radius * math.Sin(float64(i) * 2 * math.Pi)) / float64(picNumber) + center.X
		s.Camera.Origin.Y = center.Y
		s.Camera.Origin.Z = (radius * math.Cos(float64(i) * 2 * math.Pi)) / float64(picNumber) + center.Z

		s.Camera.RotationX, s.Camera.RotationX, s.Camera.RotationZ =
					camera.InitRota(0.0, (-float64(i) * 2 * math.Pi) / float64(picNumber), 0.0)

		//camera.InitRota(0.0, 0.0, 0.0)
		//s.Camera.RotationX, s.Camera.RotationX, s.Camera.RotationZ =
		image := s.Render(imgSizeX, imgSizeY, nbRaysPerPixel)
		image.SavePPM("video_img" + strconv.Itoa(i) + ".ppm")
		fmt.Println("---- img" + strconv.Itoa(i) + "---- done")
		fmt.Println(s.Camera.Origin)
	}
}
/*
func getCams(center vector3.Vector3, radius, aspectRatio, FOV float64,
	imgW, imgH, number int,) []camera.Camera {


	cams := make([]camera.Camera, number)

	for i := 0; i < number; i+=1 {
		X := radius * math.Cos(float64(i) * math.Pi / float64(number))
		Y := radius * math.Cos(float64(i) * math.Pi / float64(number))
		localVec := vector3.InitVector3(X, Y, center.Z)
		camCenter := vector3.AddVector3(localVec, center)
		camYAxis := float64(i) * math.Pi/float64(number)
		cams[i] = camera.Camera(aspectRatio, FOV, imgW, imgH, camCenter, 0.0, camYAxis, 0.0)

	}
	return cams
}

 */


