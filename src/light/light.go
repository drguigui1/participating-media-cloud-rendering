package light

import (
    "volumetric-cloud/vector3"
)

type Light struct {
    Position vector3.Vector3
    Color vector3.Vector3
}

func InitLight(position, color vector3.Vector3) Light {
    return Light{
        Position: position,
        Color: color,
    }
}
