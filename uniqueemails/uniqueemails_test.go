package uniqueemails

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumUniqueEmails(t *testing.T) {
	assert.Equal(t, 2, numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
