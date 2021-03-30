package custom_errors

import "errors"

var InvalidIdError = errors.New("invalid id")
var InvalidAssetOrUserIdError = errors.New("invalid asset or user id")
var InvalidAssetStatusError = errors.New("cannot allocate asset make sure the asset is available")
var InvalidAllocationError = errors.New("cannot deallocate asset make sure asset is allocated")
var InvalidDateFormatError = errors.New("invalid date format")
