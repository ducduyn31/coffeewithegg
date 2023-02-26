package sloky

type InfrastructureBuildPlan struct {
}

type InfrastructureBuildStep interface {
	Do()
}
