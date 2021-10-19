package stringy

import (
	"testing"
)

func TestLeft(t *testing.T) {
	width := 30
	base := "The Fox"
	want := "The Fox-----------------------"
	if got := Left(base, width, "-"); got != want {
		t.Errorf("Left = %q, want %q", got, want)
	}
}

func TestRight(t *testing.T) {
	width := 30
	base := "The Fox"
	want := "-----------------------The Fox"
	if got := Right(base, width, "-"); got != want {
		t.Errorf("Right = %q, want %q", got, want)
	}
}

func TestCenter(t *testing.T) {
	width := 30
	base := "The Fox"

	want := "------------The Fox------------"
	if got := Center(base, width, "-", true); got != want {
		t.Errorf("Center = %q, want %q", got, want)
	}
}

func TestSplitLines(t *testing.T) {
	width := 30
	base := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec tortor neque, mattis eu aliquet quis, ornare in odio. Fusce commodo velit augue, vel iaculis magna ornare eget. Aenean vitae augue quis risus finibus bibendum. Maecenas faucibus ac elit quis ornare. Nunc cursus vehicula sem, congue sodales metus rutrum sed. Sed vulputate nulla nec est vehicula pretium. Mauris pulvinar nibh consequat porttitor lobortis. Maecenas vitae dolor nunc. Donec lacus ligula, imperdiet ut orci id, viverra tempor enim. Duis et laoreet quam, vitae congue neque. Quisque dictum lorem ut purus mollis, at pharetra nibh porta."
	got := SplitLines(base, width)
	want := "Lorem ipsum dolor sit amet,"
	if got[0] != want {
		t.Errorf("Center = %q, want %q", got, want)
	}
}
