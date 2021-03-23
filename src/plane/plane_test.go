package plane

import (
	"math"
	"testing"
	"volumetric-cloud/ray"
	"volumetric-cloud/vector3"
)

func TestPlaneHit(t *testing.T) {
	normal := vector3.InitVector3(0, 0, -1)
	color := vector3.InitVector3(1, 0, 0)
	corner1 := vector3.InitVector3(0, 0, -1)
	corner2 := vector3.InitVector3(1, 0, -1)
	corner3 := vector3.InitVector3(0, 1, -1)
	plane := InitPlane(normal, color, corner1, corner2, corner3)
	raycenter := vector3.InitVector3(0.5, 0.5 , 2)
	raydirection := vector3.InitVector3(0, 0, -1)
	ray := ray.InitRay(raycenter, raydirection)
	res, hasHit := plane.Hit(ray)
	ref := 3.0
	if (math.Round(res) != ref || hasHit != true){
		t.Errorf("Error 'TestPlaneHit'")
		t.Errorf("Res: %v, %v \n", res, hasHit)
		t.Errorf("Ref: %v, %v \n", ref, true)
	}
}