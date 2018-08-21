package participates

import (
	"github.com/DSiSc/galaxy/config"
	"github.com/DSiSc/galaxy/participates/policy"
	"github.com/DSiSc/txpool/common"
	"github.com/DSiSc/txpool/common/log"
)

type Participates interface {
	PolicyName() string
	GetParticipates() ([]common.Address, error)
	ChangeParticipates() error
}

const (
	PARTICIPATES_SOLO = "solo"
	// Structure must matching with defination of config/config.json
	Symbol = "participates"
	Policy = "participates.policy"
)

func NewParticipatePolicy() (Participates, error) {
	var err error
	var participates Participates
	conf := config.New(config.ConfigAbsPath())
	participatesPolicy := conf.GetConfigItem(Policy).(string)
	switch participatesPolicy {
	case PARTICIPATES_SOLO:
		log.Info("Get participates policy is solo.")
		participates, err = policy.NewSoloPolicy()
	default:
		log.Error("Now, we only support solo policy participates.")
	}
	return participates, err
}