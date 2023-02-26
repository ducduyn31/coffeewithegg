package sunny_stream

import "coffeewithegg/apps/adam/infrastructure/core/sloky"

type SunnyStreamBuildPlan struct {
	sloky.InfrastructureBuildPlan
}

type SunnyStreamBuildStep interface {
	sloky.InfrastructureBuildStep
}

func (plan *SunnyStreamBuildPlan) Build() []*SunnyStreamBuildStep {
	return nil
}
