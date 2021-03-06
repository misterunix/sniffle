package hashing

import "testing"

func TestFileHash(t *testing.T) {
	fn := "../hashtest.txt"
	want := "019879f8ba05404753d76b016c646154"
	if got, _ := FileHash(MD5, fn); got != want {
		t.Errorf("FileHash(MD5) = %q, want %q", got, want)
	}

	want = "0a601bf91f2cf535107fc739366eb8d30d671f20eaa53e421d48acab0b6fb4cf"
	if got, _ := FileHash(SHA256, fn); got != want {
		t.Errorf("FileHash(SHA256) = %q, want %q", got, want)
	}

	want = "3110705d9b47092004a7340298b271185db66002e988a616aa3de442b521c7869851b5537e91e8d17f8c763c34491e162adbcebc5874718a8eb2eb33c7608da7"
	if got, _ := FileHash(SHA512, fn); got != want {
		t.Errorf("FileHash(SHA512) = %q, want %q", got, want)
	}

	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec tortor neque, mattis eu aliquet quis, ornare in odio. Fusce commodo velit augue, vel iaculis magna ornare eget. Aenean vitae augue quis risus finibus bibendum. Maecenas faucibus ac elit quis ornare. Nunc cursus vehicula sem, congue sodales metus rutrum sed. Sed vulputate nulla nec est vehicula pretium. Mauris pulvinar nibh consequat porttitor lobortis. Maecenas vitae dolor nunc. Donec lacus ligula, imperdiet ut orci id, viverra tempor enim. Duis et laoreet quam, vitae congue neque. Quisque dictum lorem ut purus mollis, at pharetra nibh porta."
	want = "019879f8ba05404753d76b016c646154"
	if got := StringHash(MD5, s); got != want {
		t.Errorf("StringHash(MD5) = %q, want %q", got, want)
	}

	want = "0a601bf91f2cf535107fc739366eb8d30d671f20eaa53e421d48acab0b6fb4cf"
	if got := StringHash(SHA256, s); got != want {
		t.Errorf("StringHash(SHA256) = %q, want %q", got, want)
	}

	want = "3110705d9b47092004a7340298b271185db66002e988a616aa3de442b521c7869851b5537e91e8d17f8c763c34491e162adbcebc5874718a8eb2eb33c7608da7"
	if got := StringHash(SHA512, s); got != want {
		t.Errorf("StringHash(SHA512) = %q, want %q", got, want)
	}

}
