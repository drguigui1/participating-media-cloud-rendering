package atmosphere

import (
    "math"

    "volumetric-cloud/sphere"
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/light"
)

type Atmosphere struct {
    Ground sphere.Sphere
    GroundColor vector3.Vector3
    SunImpact vector3.Vector3
    GroundAlbedo float64
    AtmosphereRadius float64 // must be higher than Ground.Radius
    Sun light.Light
    ViewerPoint vector3.Vector3 // Origin of the camera
    NbStep float64 // Step number in the raymarching
    NbStepLight float64

    // scale height values
    ScaleHeightR float64
    ScaleHeightM float64

    // G different from -1 or 1
    G float64

    BetaRayleigh vector3.Vector3
    BetaMie vector3.Vector3
}

func InitAtmosphere(ground sphere.Sphere,
                    groundColor vector3.Vector3,
                    sunImpact vector3.Vector3,
                    betaRayleigh vector3.Vector3,
                    betaMie vector3.Vector3,
                    groundAlbedo,
                    atmosphereRadius float64,
                    sun light.Light,
                    nbStep,
                    nbStepLight float64,
                    scaleHeightR float64,
                    scaleHeightM float64 ) Atmosphere {


    return Atmosphere{
        Ground: ground,
        GroundColor: groundColor,
        SunImpact : sunImpact,
        GroundAlbedo: groundAlbedo,
        AtmosphereRadius: atmosphereRadius,
        Sun: sun,
        NbStep: nbStep,
        NbStepLight: nbStepLight,
        ScaleHeightR: scaleHeightR,
        ScaleHeightM: scaleHeightM,
        G: 0.8,
        BetaRayleigh: betaRayleigh,
        BetaMie: betaMie,
    }
}

// Return color of the sky
func (a Atmosphere) ComputeRayleighMie(r ray.Ray) vector3.Vector3 {
    // make intersection with atmosphere
    atmosphere := sphere.InitSphere(a.Ground.Center, a.AtmosphereRadius)

    t, _,  hasHit := atmosphere.Hit(r)
    if !hasHit {
        // should not be possible if ray origin is inside the atmosphere
        return vector3.InitVector3(0, 0, 0)
    }

    // ray marching from 0 to t
    res := a.RayMarch(r, 0.0, t)
    res.Clamp(0.0, 1.0)
    return res
}

//func (a Atmosphere)

func (a Atmosphere) RayMarch(r ray.Ray, tmin, tmax float64) vector3.Vector3 {
    stepLength := (tmax - tmin) / a.NbStep

    intRayleigh := vector3.InitVector3(0, 0, 0)
    intMie := vector3.InitVector3(0, 0, 0)

    rayleighOD := 0.0
    mieOD := 0.0

    for i := 0; i < int(a.NbStep); i += 1 {
        p := r.RayAt(tmin)

        // compute the height of the light
        currHeight := a.GetCurrentHeight(p)

        // Approximate integral at the current step
        currRayleighApprox := math.Exp(-currHeight / a.ScaleHeightR) * stepLength
        currMieApprox := math.Exp(-currHeight / a.ScaleHeightM) * stepLength

        rayleighOD += currRayleighApprox
        mieOD += currMieApprox

        // Consider sun impact
        sunDir := a.GetSunDir(p)
        rayToLight := ray.InitRay(p, sunDir)
        intRayleigh, intMie = a.RayMarchSunContribution(rayToLight,
                                                      rayleighOD,
                                                      mieOD,
                                                      intRayleigh,
                                                      intMie,
                                                      currRayleighApprox,
                                                      currMieApprox)

        tmin += stepLength
    }

    // compute cosinus of the angle between ray dir and light dir
    theta := vector3.DotProduct(r.Direction, a.GetSunDir(a.ViewerPoint))

    // Phase functions
    rayleighPhase := PhaseFonction(theta, 0.0)
    miePhase := PhaseFonction(theta, a.G)

    // Sum of Rayleigh and Mie and mult by the sun intensity
    skyColorRayleigh := vector3.HadamarProduct(intRayleigh, a.BetaRayleigh)
    skyColorRayleigh.Mul(rayleighPhase)
    skyColorMie := vector3.HadamarProduct(intMie, a.BetaMie)
    skyColorMie.Mul(miePhase)
    return vector3.HadamarProduct(vector3.AddVector3(skyColorRayleigh, skyColorMie), a.Sun.Color)
}

func (a Atmosphere) RayMarchSunContribution(
                                        r ray.Ray,
                                        rayleighOD,
                                        mieOD float64,
                                        intRayleigh,
                                        intMie vector3.Vector3,
                                        currRayleighApprox,
                                        currMieApprox float64) (vector3.Vector3, vector3.Vector3) {
    t := 0.0
    stepLength := vector3.SubVector3(r.Origin, a.Sun.Position).Length() / a.NbStepLight
    rayleighODL := 0.0
    mieODL := 0.0
    for i := 0; i < int(a.NbStepLight); i += 1 {
        p := r.RayAt(t)

        currHeight := a.GetCurrentHeight(p)

        if currHeight < 0 {
            return intRayleigh, intMie
        }

        rayleighODL += math.Exp(-currHeight / a.ScaleHeightR) * stepLength
        mieODL += math.Exp(-currHeight / a.ScaleHeightM) * stepLength

        t += stepLength
    }

    rayleighODSum := rayleighODL + rayleighOD
    mieODSum := mieODL + mieOD
    mieLContribution := vector3.MulVector3Scalar(a.BetaMie, 1.1)
    mieLContribution.Mul(mieODSum)
    rayleightLContribution := vector3.MulVector3Scalar(a.BetaRayleigh, rayleighODSum)
    totalMRLightContribution := vector3.AddVector3(rayleightLContribution, mieLContribution)

    attenuation := vector3.InitVector3(math.Exp(-totalMRLightContribution.X), math.Exp(-totalMRLightContribution.Y), math.Exp(-totalMRLightContribution.Z))
    attenuationR := vector3.MulVector3Scalar(attenuation, currRayleighApprox)
    attenuationM := vector3.MulVector3Scalar(attenuation, currMieApprox)

    intRayleigh.AddVector3(attenuationR)
    intMie.AddVector3(attenuationM)
    return intRayleigh, intMie
}

func (a Atmosphere) GetCurrentHeight(p vector3.Vector3) float64 {
    return vector3.SubVector3(p, a.Ground.Center).Length() - a.Ground.Radius
}

func (a Atmosphere) GetSunDir(p vector3.Vector3) vector3.Vector3 {
    return vector3.UnitVector(vector3.SubVector3(a.Sun.Position, p))
}
