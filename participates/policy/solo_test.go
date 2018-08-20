package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mock_NewSoloPolicy() *SoloPolicy {
	policy, _ := NewSoloPolicy()
	return policy
}

func Test_NewSoloPolicy(t *testing.T) {
	assert := assert.New(t)
	policy, err := NewSoloPolicy()
	assert.NotNil(policy)
	assert.Nil(err)
	assert.Equal(POLICY_NAME, policy.name, "they should not be equal")
}

func Test_PolicyName(t *testing.T) {
	assert := assert.New(t)
	policy := mock_NewSoloPolicy()
	policyName := policy.PolicyName()
	assert.Equal(POLICY_NAME, policyName, "they should not be equal")
}

func Test_GetParticipates(t *testing.T) {
	assert := assert.New(t)
	policy := mock_NewSoloPolicy()
	address := policy.GetParticipates()
	assert.NotNil(address)
	assert.Equal(0, len(address), "they should not be equal")
}
