package ffclient

import (
	"fmt"
	"time"

	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/internal/flags"
)

const errorFlagNotAvailable = "flag %v is not present or disabled"
const errorWrongVariation = "wrong variation used for flag %v"

// BoolVariation return the value of the flag in boolean.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func BoolVariation(flagKey string, user ffuser.User, defaultValue bool) (bool, error) {
	return ff.BoolVariation(flagKey, user, defaultValue)
}

// IntVariation return the value of the flag in int.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func IntVariation(flagKey string, user ffuser.User, defaultValue int) (int, error) {
	return ff.IntVariation(flagKey, user, defaultValue)
}

// Float64Variation return the value of the flag in float64.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func Float64Variation(flagKey string, user ffuser.User, defaultValue float64) (float64, error) {
	return ff.Float64Variation(flagKey, user, defaultValue)
}

// StringVariation return the value of the flag in string.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func StringVariation(flagKey string, user ffuser.User, defaultValue string) (string, error) {
	return ff.StringVariation(flagKey, user, defaultValue)
}

// JSONArrayVariation return the value of the flag in []interface{}.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func JSONArrayVariation(flagKey string, user ffuser.User, defaultValue []interface{}) ([]interface{}, error) {
	return ff.JSONArrayVariation(flagKey, user, defaultValue)
}

// JSONVariation return the value of the flag in map[string]interface{}.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
func JSONVariation(
	flagKey string, user ffuser.User, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	return ff.JSONVariation(flagKey, user, defaultValue)
}

// BoolVariation return the value of the flag in boolean.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) BoolVariation(flagKey string, user ffuser.User, defaultValue bool) (bool, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).(bool)
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// IntVariation return the value of the flag in int.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) IntVariation(flagKey string, user ffuser.User, defaultValue int) (int, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).(int)
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// Float64Variation return the value of the flag in float64.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) Float64Variation(flagKey string, user ffuser.User, defaultValue float64) (float64, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).(float64)
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// StringVariation return the value of the flag in string.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) StringVariation(flagKey string, user ffuser.User, defaultValue string) (string, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).(string)
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// JSONArrayVariation return the value of the flag in []interface{}.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) JSONArrayVariation(
	flagKey string, user ffuser.User, defaultValue []interface{}) ([]interface{}, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).([]interface{})
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// JSONVariation return the value of the flag in map[string]interface{}.
// An error is return if you don't have init the library before calling the function.
// If the key does not exist we return the default value.
// Note: Use this function only if you are using multiple go-feature-flag instances.
func (g *GoFeatureFlag) JSONVariation(
	flagKey string, user ffuser.User, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	flag, err := g.getFlagFromCache(flagKey)
	if err != nil {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, err
	}

	res, ok := flag.Value(flagKey, user).(map[string]interface{})
	if !ok {
		g.notifyVariation(flagKey, user.GetKey(), defaultValue)
		return defaultValue, fmt.Errorf(errorWrongVariation, flagKey)
	}
	g.notifyVariation(flagKey, user.GetKey(), res)
	return res, nil
}

// notifyVariation is logging the evaluation result for a flag
// if no logger is provided in the configuration we are not logging anything.
func (g *GoFeatureFlag) notifyVariation(flagKey string, userKey string, value interface{}) {
	if g.config.Logger != nil {
		g.config.Logger.Printf(
			"[%v] user=\"%s\", flag=\"%s\", value=\"%v\"",
			time.Now().Format(time.RFC3339), userKey, flagKey, value)
	}
}

// getFlagFromCache try to get the flag from the cache.
// It returns an error if the cache is not init or if the flag is not present or disabled.
func (g *GoFeatureFlag) getFlagFromCache(flagKey string) (flags.Flag, error) {
	flag, err := g.cache.GetFlag(flagKey)
	if err != nil || flag.Disable {
		return flag, fmt.Errorf(errorFlagNotAvailable, flagKey)
	}
	return flag, nil
}
