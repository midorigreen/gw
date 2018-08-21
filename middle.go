package main

import (
	"fmt"
	"io"
	"time"
)

func Chain(out cmdMiddleware, mid ...cmdMiddleware) cmdMiddleware {
	return func(c Cmder) Cmder {
		fidx := len(mid) - 1
		for i := range mid {
			c = mid[fidx-i](c)
		}
		return out(c)
	}
}

func WrapEndEcho(s string) cmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			defer fmt.Printf("%v\n", s)
			return c.Run(args, stdout, stderr)
		}
		return cmdFunc(fn)
	}
}

func WrapFirstEcho(s string) cmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			fmt.Printf("%v\n", s)
			return c.Run(args, stdout, stderr)
		}
		return cmdFunc(fn)
	}
}

func WrapTime() cmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			s := time.Now()
			defer fmt.Printf("time:%v \n", time.Now().Sub(s))
			return c.Run(args, stdout, stderr)
		}
		return cmdFunc(fn)
	}
}

func WrapSlack(token, title, channel string) cmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			err := c.Run(args, stdout, stderr)
			s := NewSlack(token)
			if err != nil {
				return s.Send(title, channel, fmt.Sprintln(err))
			}
			return s.Send(title, channel, fmt.Sprintln(err))
		}
		return cmdFunc(fn)
	}
}
