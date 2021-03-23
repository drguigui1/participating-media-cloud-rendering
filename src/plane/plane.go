package plane

import (
	"math"
	"volumetric-cloud/vector3"
	"volumetric-cloud/ray"

)


type Plane struct {
	// Color for testing
	Normal vector3.Vector3
	Color vector3.Vector3
	Corner1 vector3.Vector3
	MinX float64
	MinY float64
	MinZ float64
	MaxX float64
	MaxY float64
	MaxZ float64
}



func InitPlane(normal vector3.Vector3, color vector3.Vector3, corner1 vector3.Vector3,
		corner2 vector3.Vector3, corner3 vector3.Vector3) Plane {
	minX := math.Min(corner1.X, math.Min(corner2.X, corner3.X))
	minY := math.Min(corner1.Y, math.Min(corner2.Y, corner3.Y))
	minZ := math.Min(corner1.Z, math.Min(corner2.Z, corner3.Z))
	maxX := math.Max(corner1.X, math.Max(corner2.X, corner3.X))
	maxY := math.Max(corner1.Y, math.Max(corner2.Y, corner3.Y))
	maxZ := math.Max(corner1.Z, math.Max(corner2.Z, corner3.Z))
	return Plane{Normal: normal,
				Color: color,
				Corner1: corner1,
				MinX: minX,
				MinY: minY,
				MinZ: minZ,
				MaxX: maxX,
				MaxY: maxY,
				MaxZ: maxZ,
		}
}

func (plane Plane) Hit(ray ray.Ray) (float64,bool){
	dotnd := vector3.DotProduct(plane.Normal, ray.Direction)
	if (math.Abs(dotnd) < 0.00001){
		return 0.0, false
	}
	diff := vector3.SubVector3(plane.Corner1, ray.Origin)
	t := vector3.DotProduct(plane.Normal, diff)/dotnd
	if t < 0{
		return 0.0, false
	}
	t -= 0.0001
	intersectionPts := ray.RayAt(t)
	if (plane.MinX != plane.MaxX && (intersectionPts.X < plane.MinX || intersectionPts.X > plane.MaxX)){
		return 0.0, false
	}
	if (plane.MinY != plane.MaxY && (intersectionPts.Y < plane.MinY || intersectionPts.Y > plane.MaxY)){
		return 0.0, false
	}
	if (plane.MinZ != plane.MaxZ && (intersectionPts.Z < plane.MinZ || intersectionPts.Z > plane.MaxZ)){
		return 0.0, false
	}
	return t, true
}