package reword

import "fmt"

type RewardStrategy interface {
	reward(userId int64)
}

type userType string

var (
	newUser userType = "NewUser"
	oldUser userType = "OldUser"
)

var m = map[string]string{
	"a": "1",
}

//var M = map[userType]RewardStrategy{
//	newUser: NewUserRewardStrategyA{},
//	oldUser: OldUserRewardStrategyA{},
//}

type NewUserRewardStrategyA struct {
}

func (u *NewUserRewardStrategyA) reward(userId int64) {
	fmt.Println("new user")
}

type OldUserRewardStrategyA struct {
}

func (u *OldUserRewardStrategyA) reward(userId int64) {
	fmt.Println("old user")
}
