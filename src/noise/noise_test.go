package noise

import (
    "testing"
    "reflect"
)

func TestStack(t *testing.T) {
    a := Arange(3)

    res := Stack(a, a)
    ref := []int{0,1,2,0,1,2}


    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestStack'\n")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}
