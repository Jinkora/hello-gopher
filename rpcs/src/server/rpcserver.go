package rpcs

import "errors"

type Arith int
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

/*
	提供访问的方法必须遵守以下规则:

	1. 首字母大写
	2. 只有两个参数 ， e.g 有4个参数时提示 method XXX has wrong number of ins: 5
	3. 第二个参数是一个pointer
	4. 必须返回一个error

 */
func (t*Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t*Arith) TwoTimes(a int, reply *int) error {
	*reply = a + a
	return nil
}

func (t*Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
