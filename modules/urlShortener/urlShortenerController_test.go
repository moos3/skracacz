package urlShortener

import (
	"testing"

	"github.com/moos3/skracacz/utils"
)

func TestGenerateShortUrl(t *testing.T) {
	utils.AssertStringEquals(t, "a", generateShortUrl(0))
	utils.AssertStringEquals(t, "b", generateShortUrl(1))
	utils.AssertStringEquals(t, "ba", generateShortUrl(62))
	utils.AssertStringEquals(t, "ca", generateShortUrl(124))
	utils.AssertStringEquals(t, "baa", generateShortUrl(3844))
	utils.AssertStringEquals(t, "bbb", generateShortUrl(3907))
}

func TestRevertShortUrl(t *testing.T) {
	utils.AssertIntEquals(t, 0, revertShortUrl("a"))
	utils.AssertIntEquals(t, 1, revertShortUrl("b"))
	utils.AssertIntEquals(t, 62, revertShortUrl("ba"))
	utils.AssertIntEquals(t, 124, revertShortUrl("ca"))
	utils.AssertIntEquals(t, 3844, revertShortUrl("baa"))
	utils.AssertIntEquals(t, 3907, revertShortUrl("bbb"))
}
