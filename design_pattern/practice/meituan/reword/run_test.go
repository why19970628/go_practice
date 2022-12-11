package reword

import "testing"

//
func TestOperator(t *testing.T) {
	var user userType = ""
	var userID int64 = 1

	if user == newUser {
		s := NewStrategyFactory(&NewUserRewardStrategyA{})
		s.reward(userID)
	} else if user == oldUser {
		s := NewStrategyFactory(&NewUserRewardStrategyA{})
		s.reward(userID)
	}

}
