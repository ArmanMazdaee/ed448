package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_ScalarAddition(c *C) {
	s1 := &scalar32{
		0x529eec33, 0x721cf5b5,
		0xc8e9c2ab, 0x7a4cf635,
		0x44a725bf, 0xeec492d9,
		0x0cd77058, 0x00000002,
	}
	s2 := &scalar32{0x00000001}
	expected := scalar32{
		0x529eec34, 0x721cf5b5,
		0xc8e9c2ab, 0x7a4cf635,
		0x44a725bf, 0xeec492d9,
		0x0cd77058, 0x00000002,
	}
	out := scalar32{}
	out.scalarAdd(s1, s2)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_ScalarHalve(c *C) {
	expected := scalar32{6}
	s1 := &scalar32{12}
	s2 := &scalar32{4}
	out := scalar32{}
	out.scalarHalve(s1, s2)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_littleScalarMul_Identity(c *C) {
	x := &scalar32{
		0xd013f18b, 0xa03bc31f,
		0xa5586c00, 0x5269ccea,
		0x80becb3f, 0x38058556,
		0x736c3c5b, 0x07909887,
		0x87190ede, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &scalar32{0x00000001}

	expected := &scalar32{
		0xf19fb32f, 0x62bc6ae6,
		0xed626086, 0x0e2d81d7,
		0x7a83d54b, 0x38e73799,
		0x485ad3d6, 0x45399c9e,
		0x824b12d9, 0x5ae842c9,
		0x5ca5b606, 0x3c0978b3,
		0x893b4262, 0x22c93812,
	}

	out := &scalar32{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, expected)
	out.montgomeryMultiply(out, scalarR2)
	c.Assert(out, DeepEquals, x)
}

func (s *Ed448Suite) Test_littleScalarMul_Zero(c *C) {
	x := &scalar32{
		0xd013f18b, 0xa03bc31f,
		0xa5586c00, 0x5269ccea,
		0x80becb3f, 0x38058556,
		0x736c3c5b, 0x07909887,
		0x87190ede, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &scalar32{}

	out := &scalar32{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, y)
}

func (s *Ed448Suite) Test_littleScalarMul_fullMultiplication(c *C) {
	x := &scalar32{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &scalar32{
		0xd8bedc42, 0x686eb329,
		0xe416b899, 0x17aa6d9b,
		0x1e30b38b, 0x188c6b1a,
		0xd099595b, 0xbc343bcb,
		0x1adaa0e7, 0x24e8d499,
		0x8e59b308, 0x0a92de2d,
		0xcae1cb68, 0x16c5450a,
	}

	expected := scalar32{
		0x14aec10b, 0x426d3399,
		0x3f79af9e, 0xb1f67159,
		0x6aa5e214, 0x33819c2b,
		0x19c30a89, 0x480bdc8b,
		0x7b3e1c0f, 0x5e01dfc8,
		0x9414037f, 0x345954ce,
		0x611e7191, 0x19381160,
	}

	out := scalar32{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_Add(c *C) {
	one := &scalar32{0x1}
	two := &scalar32{0x2}
	three := &scalar32{0x3}

	result := &scalar32{}
	result.Add(one, two)

	c.Assert(result, DeepEquals, three)
}

func (s *Ed448Suite) Test_Sub(c *C) {
	twelve := &scalar32{0xc}
	thirteen := &scalar32{0xd}
	one := &scalar32{0x1}

	result := &scalar32{}
	result.Sub(thirteen, twelve)

	c.Assert(result, DeepEquals, one)
}

func (s *Ed448Suite) Test_Mul(c *C) {
	x := &scalar32{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}

	y := &scalar32{
		0xd8bedc42, 0x686eb329,
		0xe416b899, 0x17aa6d9b,
		0x1e30b38b, 0x188c6b1a,
		0xd099595b, 0xbc343bcb,
		0x1adaa0e7, 0x24e8d499,
		0x8e59b308, 0x0a92de2d,
		0xcae1cb68, 0x16c5450a,
	}

	expected := &scalar32{
		0xa18d010a, 0x1f5b3197,
		0x994c9c2b, 0x6abd26f5,
		0x08a3a0e4, 0x36a14920,
		0x74e9335f, 0x07bcd931,
		0xf2d89c1e, 0xb9036ff6,
		0x203d424b, 0xfccd61b3,
		0x4ca389ed, 0x31e055c1,
	}
	x.Mul(x, y)
	c.Assert(x, DeepEquals, expected)
}

func (s *Ed448Suite) Test_Copy(c *C) {
	expected := &scalar32{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	x := expected.Copy()
	c.Assert(x, DeepEquals, expected)
}
