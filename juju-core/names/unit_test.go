// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "launchpad.net/gocheck"

	"launchpad.net/juju-core/names"
)

type unitSuite struct{}

var _ = gc.Suite(&unitSuite{})

func (s *unitSuite) TestUnitTag(c *gc.C) {
	c.Assert(names.UnitTag("wordpress/2"), gc.Equals, "unit-wordpress-2")
}

var unitNameTests = []struct {
	pattern string
	valid   bool
}{
	{pattern: "wordpress/42", valid: true},
	{pattern: "rabbitmq-server/123", valid: true},
	{pattern: "foo", valid: false},
	{pattern: "foo/", valid: false},
	{pattern: "bar/foo", valid: false},
	{pattern: "20/20", valid: false},
	{pattern: "foo-55", valid: false},
	{pattern: "foo-bar/123", valid: true},
	{pattern: "foo-bar/123/", valid: false},
	{pattern: "foo-bar/123-not", valid: false},
}

func (s *unitSuite) TestUnitNameFormats(c *gc.C) {
	for i, test := range unitNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		c.Assert(names.IsUnit(test.pattern), gc.Equals, test.valid)
	}
}

func (s *unitSuite) TestInvalidUnitTagFormats(c *gc.C) {
	for i, test := range unitNameTests {
		if !test.valid {
			c.Logf("test %d: %q", i, test.pattern)
			expect := fmt.Sprintf("%q is not a valid unit name", test.pattern)
			testUnitTag := func() { names.UnitTag(test.pattern) }
			c.Assert(testUnitTag, gc.PanicMatches, expect)
		}
	}
}
