// Package api is an implementation of dynamic service calls.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"
)

var serviceRegistry = map[string]map[string]MethodMeta{} // serviceID and its methods

type MethodMeta struct {
	Method       reflect.Method
	ParamTypes   []reflect.Type // List of parameter types
	ParamNames   []string
	ReceiverType reflect.Type // To create instances
}

type ServiceInitializer interface {
	SetDB(db *pgds.PgProvider)
	SetSession(sess session.Session)
	SetQueryID(queryID string)
}

// RegisterMethods registers all exported methods of the given type under typeName
func RegisterMethods(typeName string, t ServiceInitializer) {
	tType := reflect.TypeOf(t)

	methods := map[string]MethodMeta{}
	for i := range tType.NumMethod() {
		m := tType.Method(i)
		if !m.IsExported() {
			continue
		}
		numIn := m.Type.NumIn()
		paramTypes := []reflect.Type{}
		for j := 1; j < numIn; j++ {
			paramTypes = append(paramTypes, m.Type.In(j))
		}
		methods[m.Name] = MethodMeta{
			Method:       m,
			ParamTypes:   paramTypes,
			ReceiverType: m.Type.In(0), // receiver is always param 0
		}
	}
	serviceRegistry[typeName] = methods
}

func ConvertParamToType(param string, t reflect.Type) (reflect.Value, error) {
	isPtr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}

	    // Handle interface{} type
	    if t.Kind() == reflect.Interface {
		// For interface{}, we'll unmarshal JSON into a generic interface{}
		var v interface{}
		if err := json.Unmarshal([]byte(param), &v); err != nil {
		    return reflect.Value{}, err
		}
		if isPtr {
		    // Can't have pointer to interface, so just return the value
		    return reflect.ValueOf(&v), nil
		}
		return reflect.ValueOf(v), nil
	    }
	    
	switch t.Kind() {
	case reflect.String:
		v := reflect.ValueOf(param)
		if isPtr {
			ptr := reflect.New(t)
			ptr.Elem().Set(v)
			return ptr, nil
		}
		return v, nil

	case reflect.Struct:
		ptr := reflect.New(t)
		if err := json.Unmarshal([]byte(param), ptr.Interface()); err != nil {
			return reflect.Value{}, err
		}
		if isPtr {
			return ptr, nil
		}
		return ptr.Elem(), nil

	case reflect.Slice:
		slicePtr := reflect.New(t)
		if err := json.Unmarshal([]byte(param), slicePtr.Interface()); err != nil {
			return reflect.Value{}, err
		}
		if isPtr {
			return slicePtr, nil
		}
		return slicePtr.Elem(), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		v := reflect.ValueOf(i).Convert(t)
		if isPtr {
			ptr := reflect.New(t)
			ptr.Elem().Set(v)
			return ptr, nil
		}
		return v, nil

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		v := reflect.ValueOf(f).Convert(t)
		if isPtr {
			ptr := reflect.New(t)
			ptr.Elem().Set(v)
			return ptr, nil
		}
		return v, nil

	case reflect.Bool:
		b, err := strconv.ParseBool(param)
		if err != nil {
			return reflect.Value{}, err
		}
		v := reflect.ValueOf(b)
		if isPtr {
			ptr := reflect.New(t)
			ptr.Elem().Set(v)
			return ptr, nil
		}
		return v, nil

	default:
		return reflect.Value{}, fmt.Errorf("unsupported kind: %v", t.Kind())
	}
}

type ServiceContext struct {
	DB      *pgds.PgProvider
	Session session.Session
	QueryID string
}

func CallMethod(ctx context.Context, typeName, methodName string, paramStrs []string, svc *ServiceContext) ([]reflect.Value, error) {
	service, ok := serviceRegistry[typeName]
	if !ok {
		return nil, fmt.Errorf("type not registered: %s", typeName)
	}
	meta, ok := service[methodName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}

	// Context is the firs so add one.
	if len(paramStrs)+1 != len(meta.ParamTypes) {
		return nil, fmt.Errorf("expected %d params, got %d", len(meta.ParamTypes)-1, len(paramStrs))
	}

	// Convert params. The first param is always context.
	args := []reflect.Value{reflect.ValueOf(ctx)}
	for i, s := range paramStrs {
		v, err := ConvertParamToType(s, meta.ParamTypes[i+1])
		if err != nil {
			return nil, fmt.Errorf("param %d: %w", i+1, err)
		}
		args = append(args, v)
	}

	// Create instance of receiver
	receiver := CreateInstance(meta.ReceiverType)
	if s, ok := receiver.Interface().(ServiceInitializer); ok {
		s.SetDB(svc.DB)
		s.SetSession(svc.Session)
		s.SetQueryID(svc.QueryID)
	}
	// Call
	return meta.Method.Func.Call(append([]reflect.Value{receiver}, args...)), nil
}

func CreateInstance(receiverType reflect.Type) reflect.Value {
	if receiverType.Kind() == reflect.Ptr {
		return reflect.New(receiverType.Elem())
	}
	return reflect.New(receiverType).Elem()
}
