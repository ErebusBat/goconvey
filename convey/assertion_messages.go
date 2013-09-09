package convey

import "fmt"

const (
	success             = ""
	needOneValue        = "This expectation requires exactly one comparison value (none provided)."
	onlyAcceptsOneValue = "This expectation only accepts 1 value to be compared (and %v were provided)."
	noValuesAccepted    = "This expectation does not allow for user-supplied comparison values."
)

const ( // equality
	shouldHaveBeenEqual         = "Expected '%v' to equal '%v' (but it didn't)!"
	shouldNotHaveBeenEqual      = "Expected '%v' to NOT equal '%v' (but it did)!"
	shouldHaveResembled         = "Expected '%v' to resemble '%v' (but it didn't)!"
	shouldNotHaveResembled      = "Expected '%v' to NOT resemble '%v' (but it did)!"
	shouldBePointers            = "Both arguments should be pointers "
	shouldHaveBeenNonNilPointer = shouldBePointers + "(the %s was %s)!"
	shouldHavePointedTo         = "Expected '%v' (address: '%v') and '%v' (address: '%v') to be the same address (but their weren't)!"
	shouldNotHavePointedTo      = "Expected '%v' and '%v' to be different references (but they matched: '%v')!"
	shouldHaveBeenNil           = "Expected '%v' to be nil (but it wasn't)!"
	shouldNotHaveBeenNil        = "Expected '%v' to NOT be nil (but it was)!"
	shouldHaveBeenTrue          = "Expected 'true' (not '%v')!"
	shouldHaveBeenFalse         = "Expected 'false' (not '%v')!"
)

const ( // quantity comparisons
	shouldHaveBeenGreater        = "Expected '%v' to ge greater than '%v' (but it wasn't)!"
	shouldHaveBeenGreaterOrEqual = "Expected '%v' to be greater than or equal to '%v' (but it wasn't)!"
	shouldHaveBeenLess           = "Expected '%v' to be less than '%v' (but it wasn't)!"
	shouldHaveBeenLessOrEqual    = "Expected '%v' to be less than or equal to '%v' (but it wasn't)!"
	shouldHaveBeenWithin         = "Expected '%v' to be within '%v' of '%v' (distance was '%v')!"
	shouldNotHaveBeenWithin      = "Expected '%v' NOT to be within '%v' of '%v' (distance was '%v')!"
)

const ( // collections
	shouldHaveContained    = "Expected the collection to contain member: '%v' (but it didn't)!"
	shouldNotHaveContained = "Expected the collection NOT to contain member: '%v' (but it did)!"
	shouldHaveBeenIn       = "Expected '%v' to be in the collection (but it wasn't)!"
	shouldNotHaveBeenIn    = "Expected '%v' NOT to be in the collection (but it was)!"
)

const ( // strings
	shouldHaveStartedWith    = "Expected '%v' to start with: \n         '%v' (but it didn't)!"
	shouldNotHaveStartedWith = "Expected '%v' NOT to start with: \n         '%v' (but it did)!"
	shouldHaveEndedWith      = "Expected '%v' to end with: \n         '%v' (but it didn't)!"
	shouldNotHaveEndedWith   = "Expected '%v' NOT to end with: \n         '%v' (but it didn't)!"
)

const ( // panics
	shouldHavePanickedWith = "Expected func() to panic with '%v' (but it panicked with '%v')!"
	shouldHavePanicked     = "Expected func() to panic with '%v' (but it didn't panic at all)!"
)

const ( // type checking
	shouldHaveBeenA    = "Expected '%v' to be a '%v' (but was a '%v')!"
	shouldNotHaveBeenA = "Expected '%v to NOT be a '%v' (but it was)!"
)

func onlyOne(expected []interface{}) string {
	switch {
	case len(expected) == 0:
		return needOneValue
	case len(expected) > 1:
		return fmt.Sprintf(onlyAcceptsOneValue, len(expected))
	}
	return success
}

func none(expected []interface{}) string {
	if len(expected) > 0 {
		return noValuesAccepted
	}
	return success
}
