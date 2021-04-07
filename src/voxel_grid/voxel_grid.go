package voxel_grid

import (
    "math"
    "sync"

    "volumetric-cloud/light"
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/noise"
    "volumetric-cloud/interpolation"
    "volumetric-cloud/height_distribution"
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

    // noise
    Noise noise.PerlinNoise

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
                   step float64,
                   noise noise.PerlinNoise) VoxelGrid {

    distX := math.Abs(shift.X - oppositeCorner.X)
    distY := math.Abs(shift.Y - oppositeCorner.Y)
    distZ := math.Abs(shift.Z - oppositeCorner.Z)

    // compute number of VerticeX / VerticeY / VerticeZ
    var nbVerticeX int = int(distX / voxelSize) + 1
    var nbVerticeY int = int(distY / voxelSize) + 1
    var nbVerticeZ int = int(distZ / voxelSize) + 1

    voxelGrid := VoxelGrid{
        VoxelSize: voxelSize, // size of one voxel
        NbVerticeX: nbVerticeX,
        NbVerticeY: nbVerticeY,
        NbVerticeZ: nbVerticeZ,
        Shift: shift,
        OppositeCorner: oppositeCorner,
        Step: step,
        Noise: noise,
    }


    nbVertices := nbVerticeX * nbVerticeY * nbVerticeZ
    voxels := make([]Voxel, nbVertices)

    center := voxelGrid.GetWorldPosition(vector3.InitVector3(float64(nbVerticeX / 2.0), float64(nbVerticeY / 2.0), float64(nbVerticeZ / 2.0)))
    maxDist := vector3.SubVector3(shift, center).Length()

    for z := 0; z < nbVerticeZ; z += 1 {
        for y := 0; y < nbVerticeY; y += 1 {
            for x := 0; x < nbVerticeX; x += 1 {
                // put x,y,z in world coordinate
                worldVec := voxelGrid.GetWorldPosition(vector3.InitVector3(float64(x), float64(y), float64(z)))
                // compute distance between (x, y, z) and center and make ratio with maxdistance which is distance from center to corner
                dist := vector3.SubVector3(worldVec, center).Length()
                height := height_distribution.HeightDistribution(float64(y) /  (float64(nbVerticeY)), 10, 0.2)

                //h, _ := GaussianPdf(mu, mat_cov, seed, vector3.InitVector3(float64(x), float64(y), float64(z)))
                //height += h

                noiseValue := voxelGrid.Noise.GeneratePerlinNoise(worldVec.X, worldVec.Y, worldVec.Z)
                noiseValue *= height

                dist = dist / maxDist
                sharpness := 0.8
                d := 2.0
                dist -= 0.3
                noiseValue -= dist
                if noiseValue < 0 {
                    noiseValue = 0
                }
                noiseValue *= d
                density := 1.0 - math.Pow(sharpness, noiseValue)

                if density < 0 {
                    density = 0
                }

                voxels[x + y * nbVerticeX + z * nbVerticeX * nbVerticeY] = InitVoxel(density, 0.0, vector3.InitVector3(100.0 / 255.0,
                    100.0 / 255.0,
                    100.0 / 255.0))
            }
        }
    }

    voxelGrid.Voxels = voxels
    return voxelGrid
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
    if !vGrid.IsInsideVoxelGrid(vGrid.GetWorldPosition(vector3.InitVector3(float64(i), float64(j), float64(k)))) {
        return 0.0
    }
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

// (x,y,z) point in the real world
func (vGrid *VoxelGrid) LinearInterpolateDensity(x, y, z float64) float64 {
    // Get the index of the current position in the voxel
    p1 := vGrid.GetVoxelIndex(vector3.InitVector3(x, y, z))
    p2 := vector3.InitVector3(p1.X, p1.Y + 1, p1.Z)
    p3 := vector3.InitVector3(p1.X, p1.Y, p1.Z + 1)
    p4 := vector3.InitVector3(p1.X, p1.Y + 1, p1.Z + 1)
    p5 := vector3.InitVector3(p1.X + 1, p1.Y, p1.Z)
    p6 := vector3.InitVector3(p1.X + 1, p1.Y + 1, p1.Z)
    p7 := vector3.InitVector3(p1.X + 1, p1.Y, p1.Z + 1)
    p8 := vector3.InitVector3(p1.X + 1, p1.Y + 1, p1.Z + 1)

    d1 := vGrid.GetDensity(int(p1.X), int(p1.Y), int(p1.Z))
    d2 := vGrid.GetDensity(int(p2.X), int(p2.Y), int(p2.Z))
    d3 := vGrid.GetDensity(int(p3.X), int(p3.Y), int(p3.Z))
    d4 := vGrid.GetDensity(int(p4.X), int(p4.Y), int(p4.Z))
    d5 := vGrid.GetDensity(int(p5.X), int(p5.Y), int(p5.Z))
    d6 := vGrid.GetDensity(int(p6.X), int(p6.Y), int(p6.Z))
    d7 := vGrid.GetDensity(int(p7.X), int(p7.Y), int(p7.Z))
    d8 := vGrid.GetDensity(int(p8.X), int(p8.Y), int(p8.Z))

    // Get t value for interpolation
    t := vector3.SubVector3(vector3.InitVector3(x, y, z), vGrid.GetWorldPosition(p1))
    t.Div(vGrid.VoxelSize)

    // interpolation with x axis
    interpP1P5 := interpolation.Lerp(d1, d5, t.X)
    interpP2P6 := interpolation.Lerp(d2, d6, t.X)
    interpP3P7 := interpolation.Lerp(d3, d7, t.X)
    interpP4P8 := interpolation.Lerp(d4, d8, t.X)

    // y axis
    interpY1 := interpolation.Lerp(interpP1P5, interpP2P6, t.Y)
    interpY2 := interpolation.Lerp(interpP3P7, interpP4P8, t.Y)

    // z axis
    return interpolation.Lerp(interpY1, interpY2, t.Z)
}

// (x,y,z) point in the real world
func (vGrid *VoxelGrid) LinearInterpolateTransparency(x, y, z float64) float64 {
    // Get the index of the current position in the voxel
    p1 := vGrid.GetVoxelIndex(vector3.InitVector3(x, y, z))
    p2 := vector3.InitVector3(p1.X, p1.Y + 1, p1.Z)
    p3 := vector3.InitVector3(p1.X, p1.Y, p1.Z + 1)
    p4 := vector3.InitVector3(p1.X, p1.Y + 1, p1.Z + 1)
    p5 := vector3.InitVector3(p1.X + 1, p1.Y, p1.Z)
    p6 := vector3.InitVector3(p1.X + 1, p1.Y + 1, p1.Z)
    p7 := vector3.InitVector3(p1.X + 1, p1.Y, p1.Z + 1)
    p8 := vector3.InitVector3(p1.X + 1, p1.Y + 1, p1.Z + 1)

    t1 := vGrid.GetTransparency(int(p1.X), int(p1.Y), int(p1.Z))
    t2 := vGrid.GetTransparency(int(p2.X), int(p2.Y), int(p2.Z))
    t3 := vGrid.GetTransparency(int(p3.X), int(p3.Y), int(p3.Z))
    t4 := vGrid.GetTransparency(int(p4.X), int(p4.Y), int(p4.Z))
    t5 := vGrid.GetTransparency(int(p5.X), int(p5.Y), int(p5.Z))
    t6 := vGrid.GetTransparency(int(p6.X), int(p6.Y), int(p6.Z))
    t7 := vGrid.GetTransparency(int(p7.X), int(p7.Y), int(p7.Z))
    t8 := vGrid.GetTransparency(int(p8.X), int(p8.Y), int(p8.Z))

    // Get t value for interpolation
    t := vector3.SubVector3(vector3.InitVector3(x, y, z), vGrid.GetWorldPosition(p1))
    t.Div(vGrid.VoxelSize)

    // interpolation with x axis
    interpP1P5 := interpolation.Lerp(t1, t5, t.X)
    interpP2P6 := interpolation.Lerp(t2, t6, t.X)
    interpP3P7 := interpolation.Lerp(t3, t7, t.X)
    interpP4P8 := interpolation.Lerp(t4, t8, t.X)

    // y axis
    interpY1 := interpolation.Lerp(interpP1P5, interpP2P6, t.Y)
    interpY2 := interpolation.Lerp(interpP3P7, interpP4P8, t.Y)

    // z axis
    return interpolation.Lerp(interpY1, interpY2, t.Z)
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

func (voxelGrid *VoxelGrid) ComputeInsideLightTransparency(lights []light.Light) {
    wg := sync.WaitGroup{}
    wg.Add(voxelGrid.NbVerticeX)

    for i := 0; i < voxelGrid.NbVerticeX; i += 1 {
        go voxelGrid.ComputeInsideLightTransparencyYZ(lights, i, &wg)
    }
    wg.Wait()
}

func (voxelGrid *VoxelGrid) ComputeInsideLightTransparencyYZ(lights []light.Light, i int, wg *sync.WaitGroup) {
    for j := 0; j < voxelGrid.NbVerticeY; j += 1 {
        for k := 0; k < voxelGrid.NbVerticeZ; k += 1 {
            // get the position of the voxel point
            pWorld := voxelGrid.GetWorldPosition(vector3.InitVector3(float64(i), float64(j), float64(k)))

            // iterate on all lights
            insideTransparencySum := 0.0
            for _, light := range lights {
                lDir := vector3.UnitVector(vector3.SubVector3(light.Position, pWorld))

                // build the ray from pWorld to the light
                ray := ray.InitRay(pWorld, lDir)

                // launch the raymarching from this point to the light
                pts, _ := voxelGrid.RayMarch(ray)

                insideTransparency := 1.0
                for _, p := range pts {
                    // indexGrid := voxelGrid.GetVoxelIndex(p) // get the proper position in the grid

                    // density := voxelGrid.GetDensity(int(indexGrid.X), int(indexGrid.Y), int(indexGrid.Z))
                    density := voxelGrid.LinearInterpolateDensity(p.X, p.Y, p.Z)
                    insideTransparency *= math.Exp(-voxelGrid.Step * density)
                }
                insideTransparencySum += insideTransparency
            }

            // set the transmittance in the voxel grid (position i,j,k)
            voxelGrid.SetTransparency(i, j, k, insideTransparencySum / float64(len(lights)))
        }
    }
    wg.Done()
}

// return the proper color
func (vGrid VoxelGrid) ComputePixelColor(ray ray.Ray, lightColor vector3.Vector3) (vector3.Vector3, float64, bool) {
    pts, hasHit := vGrid.RayMarch(ray)
    if !hasHit {
        return vector3.Vector3{}, 0.0, false
    }

    var accTransparency float64 = 1.0;
    color := vector3.InitVector3(0.0, 0.0, 0.0)

    for _, p := range pts {
        var voxelLight vector3.Vector3

        // get the index in the voxelGrid of the points 'p'
        //vGridCoord := vGrid.GetVoxelIndex(p)
        //density := vGrid.GetDensity(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))
        density := vGrid.LinearInterpolateDensity(p.X, p.Y, p.Z)

        if density < 0.001 {
            continue
        }

        // get transparency / transmittance
        //insideTransparency := vGrid.GetTransparency(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))
        insideTransparency := vGrid.LinearInterpolateTransparency(p.X, p.Y, p.Z)

        // get the color at specific position of the voxel
        //voxelColor := vGrid.GetColor(int(vGridCoord.X), int(vGridCoord.Y), int(vGridCoord.Z))

        voxelLight = lightColor
        voxelLight.Mul(insideTransparency)
        voxelLight.Mul(density)

        beerlambert := math.Exp(-vGrid.Step * density)
        accTransparency *= beerlambert

        //voxelLight.Mul(accTransparency * vGrid.Step)
        voxelLight.Mul((1 - beerlambert) / density)
        color.AddVector3(voxelLight)
    }

    // compute background color
    //backgroundColor := background.RenderGradient(ray)

    // background contribution
    //backgroundColor.Mul(accTransparency)
    //color.AddVector3(backgroundColor)

    //color.Clamp(0.0, 1.0)
    return color, accTransparency, true
}
