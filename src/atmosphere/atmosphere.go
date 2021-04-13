package atmosphere

import (
    "volumetric-cloud/sphere"
    "volumetric-cloud/vector3"
)

type Atmosphere struct {
    Ground sphere.Sphere
    GroundColor vector3.Vector3
    GroundAlbedo float64
    AtmosphereRadius float64 // must be higher than Ground.Radius
}
