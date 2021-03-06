package versions

import (
	"github.com/openshift/odo/pkg/devfile/versions/common"
)

// DevfileData is an interface that defines functions for Devfile data operations
type DevfileData interface {
	GetComponents() []common.DevfileComponent
	GetAliasedComponents() []common.DevfileComponent
	GetProjects() []common.DevfileProject
	GetCommands() []common.DevfileCommand
}
