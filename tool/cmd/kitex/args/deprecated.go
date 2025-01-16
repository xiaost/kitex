package args

import "errors"

type deprecatedFlag struct{}

func (deprecatedFlag) Set(_ string) error {
	return errors.New(`flag is deprecated`)
}

func (deprecatedFlag) String() string { return "" }
