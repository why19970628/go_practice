package reword

// 工厂模式加策略模式
// 返奖规则与设计模式实践

type FactorRewardStrategyFactory struct {
}

type StrategyFactory struct {
	RewardStrategy
}

func NewStrategyFactory(rewardStrategy RewardStrategy) *StrategyFactory {
	return &StrategyFactory{RewardStrategy: rewardStrategy}
}
