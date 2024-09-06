package problem2660

func isWinner(player1 []int, player2 []int) int {
    computeScore:=func(player []int)int{
		doubleRound:=0
		sum:=0
		for i:=0;i<len(player);i++{
			if doubleRound>0{
				sum+=2*player[i]
				doubleRound--
			}else{
				sum+=player[i]
			}

			if player[i]==10{
				doubleRound=2
			}

		}
		return sum
	}
	player1Score:=computeScore(player1)
	player2Score:=computeScore(player2)
	if player1Score==player2Score{
		return 0
	}else if player1Score<player2Score{
		return 2
	}else {
		return 1
	}
}