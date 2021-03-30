package voxel_grid

import (
    "math"
  //  "math/rand"

    "volumetric-cloud/light"
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/background"
)

/*
** One voxel == one cube
** 'Density': density in the voxel
** 'Transparency': how the voxel transmit the light
** 'Color': color inside the voxel
*/
type Voxel struct {
    Density float64
    Transparency float64
    Color vector3.Vector3
}

/*
** 3D grid of voxels
** 'VoxelSize': size of one voxel (one cube)
** 'NbVoxelX': Number of lattice for the x axis: (NbVoxelsX - 1) = NbCubesX
** 'ShiftedPoint': Shift the voxel grid in the space
*/
type VoxelGrid struct {
    VoxelSize float64
    NbVerticeX int
    NbVerticeY int
    NbVerticeZ int
    Step float64

    Shift vector3.Vector3
    OppositeCorner vector3.Vector3
    Voxels []Voxel
}

func InitVoxel(density, transparency float64, color vector3.Vector3) Voxel {
    return Voxel{
        Density: density,
        Transparency: transparency,
        Color: color,
    }
}

func InitVoxelGrid(voxelSize float64,
                   shift vector3.Vector3,
                   oppositeCorner vector3.Vector3,
                   step float64) VoxelGrid {

    distX := math.Abs(shift.X - oppositeCorner.X)
    distY := math.Abs(shift.Y - oppositeCorner.Y)
    distZ := math.Abs(shift.Z - oppositeCorner.Z)

    // compute number of VerticeX / VerticeY / VerticeZ
    var nbVerticeX int = int(distX / voxelSize) + 1
    var nbVerticeY int = int(distY / voxelSize) + 1
    var nbVerticeZ int = int(distZ / voxelSize) + 1


    nbVertices := nbVerticeX * nbVerticeY * nbVerticeZ
    voxels := make([]Voxel, nbVertices)
    maxDist := math.Sqrt(math.Pow(float64(nbVerticeX/2), 2.0) + math.Pow(float64(nbVerticeY/2), 2.0) + math.Pow(float64(nbVerticeZ/2), 2.0))
    for z := 0; z < nbVerticeZ; z += 1 {
        for y := 0; y < nbVerticeY; y += 1 {
            for x := 0; x < nbVerticeX; x += 1 {
                // compute distance between (x, y, z) and center and make ratio with maxdistance which is distance from center to corner
                dist := math.Sqrt(math.Pow(float64(x - nbVerticeX/2), 2.0) + math.Pow(float64(y-nbVerticeY/2), 2.0) + math.Pow(float64(z-nbVerticeZ/2), 2.0))
                // result between 0 and 1
                density := 1-dist/maxDist
                voxels[x + y * nbVerticeX + z * nbVerticeX * nbVerticeY] = InitVoxel(density, 0.0, vector3.InitVector3(100.0 / 255.0,
                    100.0 / 255.0,
                    100.0 / 255.0))
            }
        }

    }
    /* Init voxels with random floats between 0 and 1
    for i := 0; i < nbVertices; i += 1 {
        // TODO (change with perlin noise for density)
        density := rand.Float64()
        // transmitivity will be set latter
        voxels[i] = InitVoxel(density, 0.0, vector3.InitVector3(100.0 / 255.0,
                                                                100.0 / 255.0,
                                                                100.0 / 255.0))
    }
     */


    return VoxelGrid{
        VoxelSize: voxelSize, // size of one voxel
        NbVerticeX: nbVerticeX,
        NbVerticeY: nbVerticeY,
        NbVerticeZ: nbVerticeZ,
        Shift: shift,
        OppositeCorner: oppositeCorner,
        Voxels: voxels,
        Step: step,
    }
}

// Shift the voxel grid from its position to (0,0,0)
// so, we need to shift the input point to get its position in the voxel grid (with (0,0,0) position)
func (vGrid VoxelGrid) ShiftToVoxelCoordinates(p vector3.Vector3) vector3.Vector3 {
    return vector3.SubVector3(p, vGrid.Shift)
}

// Opposite as the previous one
func (vGrid VoxelGrid) ShiftToWorldCoordinates(voxelCoordinatePoint vector3.Vector3) vector3.Vector3 {
    return vector3.AddVector3(voxelCoordinatePoint, vGrid.Shift)
}

// Get the position in the world coordinate system of the input coordinates of the voxel grid
// ex:
// input 0,0,0 (first vertice of the voxelGrid)
// 1) (0,0,0) * voxelSize -> (0,0,0)
// 2) (0,0,0) + shift
// -> the output will be the shift
func (vGrid VoxelGrid) GetWorldPosition(v vector3.Vector3) vector3.Vector3 {
    res := vector3.MulVector3Scalar(v, vGrid.VoxelSize)
    return vGrid.ShiftToWorldCoordinates(res) // shift the points to real world coordinates
}

// From world position to voxel position
// ex: 
// input = (0,0,-4) (shift == 0,0,-4)
// output -> (0, 0, 0)
func (vGrid VoxelGrid) GetVoxelIndex(v vector3.Vector3) vector3.Vector3 {
    // shift the point into the voxel coordinates (voxel start at (0,0,0) in the world coordinates)
    res := vGrid.ShiftToVoxelCoordinates(v)
    res.Div(vGrid.VoxelSize)
    res.Floor()
    return res
}

func (vGrid VoxelGrid) GetDensity(i, j, k int) float64 {
    return vGrid.Voxels[i + j * vGrid.NbVerticeX + k * vGrid.NbVerticeX * vGrid.NbVerticeY].Density
}

func (vGrid VoxelGrid) GetTransparency(i, j, k int) float64 {
    return vGrid.Voxels[i + j * vGrid.NbVerticeX + k * vGrid.NbVerticeX * vGrid.NbVerticeY].Transparency
}

func (vGrid VoxelGrid) GetColor(i, j, k int) vector3.Vector3 {
    return vGrid.Voxels[i + j * vGrid.NbVerticeX + k * vGrid.NbVerticeX * vGrid.NbVerticeY].Color
}

func (vGrid *VoxelGrid) SetTransparency(i, j, k int, value float64) {
    vGrid.Voxels[i + j * vGrid.NbVerticeX + k * vGrid.NbVerticeX * vGrid.NbVerticeY].Transparency = value
}

func (vGrid VoxelGrid) IsInsideVoxelGrid(p vector3.Vector3) bool {
    pVoxel := vGrid.ShiftToVoxelCoordinates(p)

    if pVoxel.X < 0 || pVoxel.X > (vGrid.VoxelSize * float64(vGrid.NbVerticeX - 1)) ||
       pVoxel.Y < 0 || pVoxel.Y > (vGrid.VoxelSize * float64(vGrid.NbVerticeY - 1)) ||
       pVoxel.Z < 0 || pVoxel.Z > (vGrid.VoxelSize * float64(vGrid.NbVerticeZ - 1)) {
        return false
    }
    return true
}

// AABB cube intersection
func (vGrid VoxelGrid) Hit(ray ray.Ray) (float64, bool, vector3.Vector3) {
    // compute tmin and tmax for x component
    // check where the plane intersect planes made by x component of both points
    // tmin: intersection with the clothest x plane (between both point)

    tmin := (vGrid.Shift.X - ray.Origin.X) / ray.Direction.X
    tmax := (vGrid.OppositeCorner.X - ray.Origin.X) / ray.Direction.X

    if tmin > tmax {
        tmin, tmax = tmax, tmin // swap values
    }

    tymin := (vGrid.Shift.Y - ray.Origin.Y) / ray.Direction.Y
    tymax := (vGrid.OppositeCorner.Y - ray.Origin.Y) / ray.Direction.Y

    if tymin > tymax {
        tymin, tymax = tymax, tymin
    }

    // the ray does not hit the cube
    if tmin > tymax || tymin > tmax {
        return 0.0, false, vector3.Vector3{}
    }

    if tymin > tmin {
        tmin = tymin
    }

    if tymax < tmax {
        tmax = tymax
    }

    tzmin := (vGrid.Shift.Z - ray.Origin.Z) / ray.Direction.Z
    tzmax := (vGrid.OppositeCorner.Z - ray.Origin.Z) / ray.Direction.Z

    if tzmin > tzmax {
        tzmin, tzmax = tzmax, tzmin
    }

    if tmin > tzmax || tzmin > tmax {
        return 0.0, false, vector3.Vector3{}
    }

    if tzmin > tmin {
        tmin = tzmin
    }

    if tzmax < tmax {
        tmax = tzmax
    }

    if tmin < 0 {
        return 0.0, false, vector3.Vector3{}
    }

    tmin += 0.0001
    p := ray.RayAt(tmin)

    var color vector3.Vector3
    if Round4(p.X) == vGrid.Shift.X {
        color = vector3.InitVector3(255, 0, 0)
    } else if Round4(p.X) == vGrid.OppositeCorner.X {
        color = vector3.InitVector3(255, 255, 0)
    } else if Round4(p.Y) == vGrid.Shift.Y {
        color = vector3.InitVector3(255, 0, 255)
    } else if Round4(p.Y) == vGrid.OppositeCorner.Y {
        color = vector3.InitVector3(0, 0, 255)
    } else if Round4(p.Z) == vGrid.Shift.Z {
        color = vector3.InitVector3(0, 255, 0)
    } else if Round4(p.Z) == vGrid.OppositeCorner.Z {
        color = vector3.InitVector3(100, 100, 100)
    }

    return tmin, true, color
}

func (voxelGrid VoxelGrid) RayMarch(ray ray.Ray) ([]vector3.Vector3, bool) {
    // Check if already inside
    var t float64
    var hasHit bool
    if voxelGrid.IsInsideVoxelGrid(ray.Origin) {
        t = 0
    } else {
        // Get first point (origin)
        t, hasHit, _ = voxelGrid.Hit(ray)

        if !hasHit {
            return []vector3.Vector3{}, false
        }

    }

    // allocate slice
    points := make([]vector3.Vector3, 0)

    // compute origin
    o := ray.RayAt(t + 0.001)

    for voxelGrid.IsInsideVoxelGrid(o) {
        points = append(points, o)
        // TODO: add random value to step
        o = vector3.AddVector3(o, vector3.MulVector3Scalar(ray.Direction, voxelGrid.Step))
    }

    return points, true
}

func (voxelGrid *VoxelGrid) ComputeInsideLightTransparency(light light.Light) {
    for i := 0; i < voxelGrid.NbVerticeX; i += 1 {
        for j := 0; j < voxelGrid.NbVerticeY; j += 1 {
            for k := 0; k < voxelGrid.NbVerticeZ; k += 1 {
                // get the position of the voxel point
                pWorld := voxelGrid.GetWorldPosition(vector3.InitVector3(float64(i), float64(j), float64(k)))
                lDir := vector3.UnitVector(vector3.SubVector3(light.Position, pWorld))

                // build the ray from pWorld to the light
                ray := ray.InitRay(pWorld, lDir)

                // launch the raymarching from this point to the light
                pts, _ := voxelGrid.RayMarch(ray)

                insideTransparency := 1.0
                for _, p := range pts {
                    indexGrid := voxelGrid.GetVoxelIndex(p) // get the proper position in the grid

                    // TODO maybe interpolate density (make function 'GetDensityInterp')
                    density := voxelGrid.GetDensity(int(indexGrid.X), int(indexGrid.Y), int(indexGrid.Z))
                    insideTransparency *= math.Exp(-voxelGrid.Step * density)
                }

                // set the transmittance in the voxel grid (position i,j,k)
                voxelGrid.SetTransparency(i, j, k, insideTransparency)
            }
        }
    }
}

// return the proper color
func (vGrid VoxelGrid) ComputePixelColor(ray ray.Ray, lightColor vector3.Vector3) (vector3.Vector3, bool) {
    pts, hasHit := vGrid.RayMarch(ray)
    if !hasHit {
        return vector3.Vector3{}, false
    }

    var accTransparency float64 = 1.0;
    color := vector3.InitVector3(0.0, 0.0, 0.0)

    for _, p := range pts {
        var voxelLight vector3.Vector3

        // get the index in the voxelGrid of the points 'p'
        // TODO: use interpolation
        vGridCoord := vGrid.GetVoxelIndex(p)
        density := vGrid.GetDensity(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))

        // get transparency / transmittance
        insideTransparency := vGrid.GetTransparency(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))

        // get the color at specific position of the voxel
        //voxelColor := vGrid.GetColor(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))

        //voxelLight = vector3.HadamarProduct(voxelColor, lightColor)
        voxelLight = lightColor
        voxelLight.Mul(insideTransparency)
        voxelLight.Mul(density)

        accTransparency *= math.Exp(-vGrid.Step * density)

        voxelLight.Mul(accTransparency * vGrid.Step)
        color.AddVector3(voxelLight)
    }

    // compute background color
    backgroundColor := background.RenderGradient(ray)

    // background contribution
    backgroundColor.Mul(accTransparency)
    color.AddVector3(backgroundColor)

    color.Clamp(0.0, 1.0)
    return color, true
}
