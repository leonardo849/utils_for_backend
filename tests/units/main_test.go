package units_test

import (
	"os"
	"reflect"
	"testing"
	"time"
	"unicode"
	"github.com/leonardo849/utils_for_backend/pkg/date"
	"github.com/leonardo849/utils_for_backend/pkg/hash"
	"github.com/leonardo849/utils_for_backend/pkg/random"

)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}





func TestHash(t *testing.T) {
	password := "batman123"
	hashP, err := hash.StringToHash(password)
	if err != nil {
		t.Error(err.Error())
	}
	if !hash.CompareHash(password, hashP) {
		t.Error("compare hash error")
	} 
	if hash.CompareHash(password+"32", hashP) {
		t.Error("compare hash error")
	}
}


func isPointer(i any) bool {
    return reflect.TypeOf(i).Kind() == reflect.Ptr
}




func TestGetPtrDate(t *testing.T) {
	now := time.Now()
	ptr := date.PtrTime(now)
	if !isPointer(ptr) {
		t.Error("it is not a pointer")
	}
}

func isNumeric(s string) bool {
    if s == "" {
        return false 
    }
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}

func TestRandomCode(t  *testing.T) {
	for i := 0; i<4; i++ {
		code := random.EncodeToString(6)
		if len(code) != 6 || !isNumeric(code) {
			t.Error("it doesn't have 6 characterics or it isn't numeric")
		}
	}
}