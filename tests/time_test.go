package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	format := "2006-01-02T15:04:05Z"
	now := time.Now()
	fmt.Println("now:", now.Format(format))

	a, _ := time.Parse(format, "2020-01-14T02:51:17Z")

	fmt.Println("now.Format(format)", now.Format(format), "a.Format(format)", a.Format(format))
	fmt.Println("now.After(a)", now.After(a))
	fmt.Println("now.Before(a)", now.Before(a))
	fmt.Println("now.Unix()", now.Unix(), "a.Unix()", a.Unix())
}
