package role

import (
	"fmt"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/galaxy/participates"
	"github.com/DSiSc/galaxy/role/common"
	"github.com/DSiSc/galaxy/role/config"
	"github.com/DSiSc/galaxy/role/policy"
	"github.com/DSiSc/txpool/log"
)

type Role interface {
	PolicyName() string
	RoleAssignments() (map[types.NodeAddress]common.Roler, error)
	GetRoles(address types.NodeAddress) common.Roler
}

func NewRole(p participates.Participates, address types.NodeAddress, conf config.RoleConfig) (Role, error) {
	var err error
	var role Role
	rolePolicy := conf.PolicyName
	switch rolePolicy {
	case policy.SOLO_POLICY:
		log.Info("Get role policy is solo.")
		role, err = policy.NewSoloPolicy(p, address)
	default:
		log.Error("Now, we only support solo role policy.")
		err = fmt.Errorf("%s", "Unknown policy type.")
	}
	return role, err
}
