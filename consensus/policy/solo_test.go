package policy

import (
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/galaxy/consensus/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSoloPolicy(t *testing.T) {
	asserts := assert.New(t)
	sp, err := NewSoloPolicy(nil)
	asserts.Nil(err)
	asserts.NotNil(sp)
	asserts.Equal(SOLO_POLICY, sp.name)
	asserts.Nil(sp.participates)
	asserts.Equal(SOLO_POLICY, sp.PolicyName())
}

func mock_proposal() *common.Proposal {
	var block types.Block
	return &common.Proposal{
		Block: &block,
	}
}

func mock_solo_proposal() *SoloProposal {
	return &SoloProposal{
		propoasl: nil,
		version:  0,
		status:   common.Proposing,
	}
}

func Test_toSoloProposal(t *testing.T) {
	asserts := assert.New(t)
	p := mock_proposal()
	proposal := toSoloProposal(p)
	asserts.NotNil(proposal)
	asserts.Equal(common.Proposing, proposal.status)
	asserts.Equal(common.Version(1), proposal.version)
	asserts.NotNil(proposal.propoasl)
}

func Test_prepareConsensus(t *testing.T) {
	asserts := assert.New(t)
	sp, _ := NewSoloPolicy(nil)
	proposal := mock_solo_proposal()

	err := sp.prepareConsensus(proposal)
	asserts.NotNil(err)

	proposal.version = 1
	err = sp.prepareConsensus(proposal)
	asserts.Nil(err)
	asserts.Equal(common.Propose, proposal.status)
}

func Test_submitConsensus(t *testing.T) {
	asserts := assert.New(t)
	proposal := mock_solo_proposal()
	sp, _ := NewSoloPolicy(nil)
	ok, err := sp.submitConsensus(proposal)
	asserts.False(ok)
	asserts.NotNil(err)

	proposal.status = common.Propose
	ok, err = sp.submitConsensus(proposal)
	asserts.True(ok)
	asserts.Nil(err)
	asserts.Equal(common.Committed, proposal.status)
}

func Test_ToConsensus(t *testing.T) {
	asserts := assert.New(t)
	proposal := mock_proposal()
	sp, _ := NewSoloPolicy(nil)
	ok, err := sp.ToConsensus(proposal)
	asserts.True(ok)
	asserts.Nil(err)
	asserts.Equal(common.Version(1), version)
}
