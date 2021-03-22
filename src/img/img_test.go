package img

import (
    "testing"
    "reflect"
)

func TestSetPixel1(t *testing.T) {
    img := InitImg(2, 2)
    img.SetPixel(1, 1, 255, 255, 255)

    ref := InitImg(2, 2)
    ref.Pixels[9] = 255
    ref.Pixels[10] = 255
    ref.Pixels[11] = 255

    if !reflect.DeepEqual(ref, img) {
        t.Errorf("Error 'TestSetPixels1'")
        t.Errorf("res: %v\n", img)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestSetPixel2(t *testing.T) {
    img := InitImg(2, 2)
    img.SetPixel(1, 0, 255, 255, 255)

    ref := InitImg(2, 2)
    ref.Pixels[6] = 255
    ref.Pixels[7] = 255
    ref.Pixels[8] = 255

    if !reflect.DeepEqual(ref, img) {
        t.Errorf("Error 'TestSetPixels2'")
        t.Errorf("res: %v\n", img)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestSetPixel3(t *testing.T) {
    img := InitImg(2, 2)
    img.SetPixel(0, 1, 255, 255, 255)

    ref := InitImg(2, 2)
    ref.Pixels[3] = 255
    ref.Pixels[4] = 255
    ref.Pixels[5] = 255

    if !reflect.DeepEqual(ref, img) {
        t.Errorf("Error 'TestSetPixels3'")
        t.Errorf("res: %v\n", img)
        t.Errorf("ref: %v\n", ref)
    }
}

func TestGetPixel1(t *testing.T) {
    img := InitImg(2, 2)
    img.Pixels[3] = 200
    img.Pixels[4] = 100
    img.Pixels[5] = 250

    r, g, b, err := img.GetPixel(0, 1)
    if err != nil {
        t.Errorf("Error 'TestGetPixel'")
    }

    if r != 200 || g != 100 || b != 250 {
        t.Errorf("Error 'TestGetPixel'")
        t.Errorf("ref (rgb): (200, 100, 250)")
        t.Errorf("res (rgb): (%v, %v, %v)", r, g, b)
    }
}

func TestGetPixel2(t *testing.T) {
    img := InitImg(2, 2)
    img.Pixels[6] = 200
    img.Pixels[7] = 100
    img.Pixels[8] = 250

    r, g, b, err := img.GetPixel(1, 0)
    if err != nil {
        t.Errorf("Error 'TestGetPixel'")
    }

    if r != 200 || g != 100 || b != 250 {
        t.Errorf("Error 'TestGetPixel'")
        t.Errorf("ref (rgb): (200, 100, 250)")
        t.Errorf("res (rgb): (%v, %v, %v)", r, g, b)
    }
}
