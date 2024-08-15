package retry_function

import (
	"github.com/pkg/errors"
	"reflect"
	"time"
)

// RetryFunc wrapper for executing a function with repeated attempts
func RetryFunc(fn interface{}, args []interface{}, maxAttempts int, delay time.Duration) (interface{}, error) {
	var err error
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}

	fnType := fnValue.Type()
	if fnType.NumOut() == 0 {
		return nil, errors.New("fn must return at least one value")
	}

	inputs := make([]reflect.Value, len(args))
	for i, arg := range args {
		if arg == nil {
			inputs[i] = reflect.Zero(fnType.In(i))
		} else {
			inputs[i] = reflect.ValueOf(arg)
		}
	}

	for attempt := 0; attempt < maxAttempts; attempt++ {
		results := fnValue.Call(inputs)
		if fnType.Out(fnType.NumOut()-1).Kind() == reflect.Interface && fnType.Out(fnType.NumOut()-1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			err, _ = results[fnType.NumOut()-1].Interface().(error)
			if err == nil {
				return results[0].Interface(), nil
			}
		}

		if attempt < maxAttempts-1 {
			time.Sleep(delay)
		}
	}

	return nil, errors.Wrap(err, "max attempts reached")
}
