package custom_errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var ErrRedisFailedSetValue = func(key string, err error) error {
	return errors.Wrap(err, fmt.Sprintf("failed to set key %s", key))
}

var ErrRedisFailedGetValue = func(key string, err error) error {
	return errors.Wrap(err, fmt.Sprintf("failed to get value for key %s", key))
}

var ErrRedisFailedDeleteValue = func(key string, err error) error {
	return errors.Wrap(err, fmt.Sprintf("failed to delete key %s", key))
}

var ErrRedisKeyNotFound = func(key string, err error) error {
	return errors.Wrap(err, fmt.Sprintf("key %s does not exist", key))
}
