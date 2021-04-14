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
    GroundAlbedo float64
    AtmosphereRadius float64 // must be higher than Ground.Radius
    Sun light.Light
    ViewerPoint vector3.Vector3 // Origin of the camera
    NbStep float64 // Step number in the raymarching
    NbStepLight float64

    // scale height values
    HR float64
    HM float64

    G float64

    // Beta coef for Rayleigh and Mie
    // Can be computed and not harcoded
    BetaRayleigh vector3.Vector3
    BetaMie vector3.Vector3
}

func InitAtmosphere(ground sphere.Sphere,
                    groundColor vector3.Vector3,
                    groundAlbedo,
                    atmosphereRadius float64,
                    sun light.Light,
                    nbStep,
                    nbStepLight float64) Atmosphere {
    HR := 7994.0
    HM := 1200.0
    g := 0.76
    betaRayleigh := vector3.InitVector3(0.0000038, 0.0000135, 0.0000331)
    vMie := 0.000021
    betaMie := vector3.InitVector3(vMie, vMie, vMie)

    return Atmosphere{
        Ground: ground,
        GroundColor: groundColor,
        GroundAlbedo: groundAlbedo,
        AtmosphereRadius: atmosphereRadius,
        Sun: sun,
        NbStep: nbStep,
        NbStepLight: nbStepLight,
        HR: HR,
        HM: HM,
        G: g,
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
    return a.RayMarch(r, 0.0, t)
}

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
        currRayleighApprox := math.Exp(-currHeight / a.HR) * stepLength
        currMieApprox := math.Exp(-currHeight / a.HM) * stepLength

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
    rayleighPhase := RayleighPhase(theta)
    miePhase := MiePhase(theta, a.G)

    // Sum of Rayleigh and Mie and mult by the sun intensity
    skyColorRayleigh := vector3.HadamarProduct(intRayleigh, a.BetaRayleigh)
    skyColorRayleigh.Mul(rayleighPhase)
    skyColorRayleigh.Mul(20.0)
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

        rayleighODL += math.Exp(-currHeight / a.HR) * stepLength
        mieODL += math.Exp(-currHeight / a.HM) * stepLength

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

func (a Atmosphere) GetSunHeight() float64 {
    return vector3.SubVector3(a.Sun.Position, a.Ground.Center).Length() - a.Ground.Radius
}
