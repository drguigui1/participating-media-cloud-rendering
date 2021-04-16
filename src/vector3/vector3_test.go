package vector3

import (
    "testing"
    "reflect"
)

func TestVector3Length(t *testing.T) {
    v3 := InitVector3(0.0, 0.0, 3.0)
    ref := 3.0

    res := v3.Length()

    if res != ref {
        t.Errorf("Error 'Vector3Length'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3AddVector3(t *testing.T) {
    v1 := InitVector3(2.0, 1.0, 3.0)
    v2 := InitVector3(2.0, 1.0, 3.0)
    ref := InitVector3(4.0, 2.0, 6.0)
    res := AddVector3(v1, v2)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'Vector3 AddVector3'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3AddVector4(t *testing.T) {
    v1 := InitVector3(2.0, 1.0, 3.0)
    v2 := InitVector3(2.0, 1.0, 3.0)
    ref := InitVector3(4.0, 2.0, 6.0)
    v1.AddVector3(v2)

    if !reflect.DeepEqual(v1, ref) {
        t.Errorf("Error 'Vector3 AddVector4'\n")
        t.Errorf("res: %v\n", v1)
        t.Errorf("ref: %v\n", ref)
    }
}


func TestVector3SubVector3(t *testing.T) {
    v1 := InitVector3(2.0, 1.0, 4.0)
    v2 := InitVector3(2.0, 2.0, 3.0)
    ref := InitVector3(0.0, -1.0, 1.0)
    res := SubVector3(v1, v2)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'Vector3 SubVector3'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3DivVector3(t *testing.T) {
    v1 := InitVector3(2.0, 1.0, 4.0)
    ref := InitVector3(1.0, 0.5, 2.0)
    res := DivVector3(v1, 2.0)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'Vector3 DivVector3'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3DotProductVector3(t *testing.T) {
    v1 := InitVector3(1.0, 2.0, 3.0)
    v2 := InitVector3(6.0, 7.0, 8.0)
    ref := 44.0
    res := DotProduct(v1, v2)

    if res != ref {
        t.Errorf("Error 'Vector3 DotProduct'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3UnitVector(t *testing.T) {
    v1 := InitVector3(1.0, 2.0, 3.0)
    res := UnitVector(v1)
    ref := 1.0

    if res.Length() != ref {
        t.Errorf("Error 'Vector3 UnitVector'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestVector3CrossProduct(t *testing.T) {
    v1 := InitVector3(1.0, 1.0, 3.0)
    v2 := InitVector3(1.0, 0.0, 2.0)

    res := CrossProduct(v1, v2)
    ref := InitVector3(2.0, 1.0, -1.0)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'Vector3 CrossProduct'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestMultMatVec3(t *testing.T) {
    v := InitVector3(1.0, 2.0, 3.0)
    m := []float64{1.0, 4.0, 2.0, 1.0, 1.0, 0.0, 0.0, 1.0, 0.0}

    ref := InitVector3(15.0, 3.0, 2.0)
    res := MultMatVec3(m, v)

    if !reflect.DeepEqual(ref, res) {
        t.Errorf("Error 'MultMatVec3'\n")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}
