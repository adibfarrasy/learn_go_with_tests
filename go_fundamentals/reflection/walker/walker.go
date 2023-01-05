package walker

import "reflect"

func Walk(x any, fn func(input string)) {
	r := getValue(x)

    walkValue := func(value reflect.Value) {
        Walk(value.Interface(), fn)
    }

	switch r.Kind() {
	case reflect.Struct:
		for i := 0; i < r.NumField(); i++ {
            walkValue(r.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < r.Len(); i++ {
            walkValue(r.Index(i))
		}
    case reflect.Map:
        for _, key := range r.MapKeys() {
            walkValue(r.MapIndex(key))
        }
    case reflect.Chan:
        for v, ok := r.Recv(); ok; v, ok = r.Recv() {
            walkValue(v)
        }
    case reflect.Func:
        rCallResult := r.Call(nil)
        for _, res := range rCallResult {
            walkValue(res)
        }
	case reflect.String:
		fn(r.String())
	}
}

func getValue(x any) reflect.Value {
	r := reflect.ValueOf(x)

	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}

	return r
}
