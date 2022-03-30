package flyweight享元

//复用课程中的 ，
//如果我们现在正在做一个棋牌类的游戏，例如象棋，无论是什么对局，棋子的基本属性其实是固定的，并不会因为随着下棋的过程变化。
//那我们就可以把棋子变为享元，让所有的对局都共享这些对象，以此达到节省内存的目的。

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board1 := NewChessBoard()
	board1.Move(1, 1, 2)
	board2 := NewChessBoard()
	board2.Move(2, 2, 3)

	assert.Equal(t, board1.chessPieces[1].Unit, board2.chessPieces[1].Unit)
	assert.Equal(t, board1.chessPieces[2].Unit, board2.chessPieces[2].Unit)
}
