package main

import (
	"context"
	"fmt"
)

type Hello struct {
}

func (g *Hello) Info(ctx context.Context, req *pd.I) {
	fmt.Println()

}
