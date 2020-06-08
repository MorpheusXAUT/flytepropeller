// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_resyncPeriod", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("resyncPeriod"); err == nil {
				assert.Equal(t, string("30s"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "30s"

			cmdFlags.Set("resyncPeriod", testValue)
			if vString, err := cmdFlags.GetString("resyncPeriod"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ResyncPeriod)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metricsPrefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("metricsPrefix"); err == nil {
				assert.Equal(t, string("flyteK8sSchedulerExtension"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metricsPrefix", testValue)
			if vString, err := cmdFlags.GetString("metricsPrefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MetricsPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_profilerPort", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("profilerPort"); err == nil {
				assert.Equal(t, string("10254"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "10254"

			cmdFlags.Set("profilerPort", testValue)
			if vString, err := cmdFlags.GetString("profilerPort"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ProfilerPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_workers", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				assert.Equal(t, int(4), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("workers", testValue)
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_unrestrictedCpuAutoscalingLimit", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("unrestrictedCpuAutoscalingLimit"); err == nil {
				assert.Equal(t, int(30), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("unrestrictedCpuAutoscalingLimit", testValue)
			if vInt, err := cmdFlags.GetInt("unrestrictedCpuAutoscalingLimit"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.UnrestrictedCPUAutoscalingLimit)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_unrestrictedMemoryAutoscalingLimit", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("unrestrictedMemoryAutoscalingLimit"); err == nil {
				assert.Equal(t, int(30), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("unrestrictedMemoryAutoscalingLimit", testValue)
			if vInt, err := cmdFlags.GetInt("unrestrictedMemoryAutoscalingLimit"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.UnrestrictedMemoryAutoscalingLimit)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}