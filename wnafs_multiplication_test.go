package ed448

import (
	"encoding/hex"

	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) TestPrepareWNAFTable(c *C) {
	tableSize := uint(4)
	expected := [16]*twPNiels{
		//0
		newTwistedPNiels(
			[56]byte{0x5f, 0xee, 0x87, 0xbc, 0x27, 0x10, 0xa9, 0xd4, 0x39, 0x2d, 0xf7, 0xe0, 0xa6, 0x34, 0x46, 0x07, 0x45, 0x97, 0x84, 0xd0, 0x5f, 0x91, 0xbb, 0xa6, 0x8b, 0x64, 0xe9, 0x76, 0x7a, 0x5d, 0xaa, 0x64, 0x0d, 0x68, 0xd9, 0xc7, 0xc7, 0xfc, 0xc6, 0x32, 0x0d, 0x73, 0xd9, 0x61, 0xd6, 0x9f, 0x8b, 0x04, 0x95, 0x67, 0x66, 0xd1, 0xfd, 0xc7, 0x8d, 0xd9},
			[56]byte{0xba, 0x47, 0x5b, 0x09, 0xe4, 0x4c, 0xd9, 0x91, 0x59, 0xa0, 0x9f, 0x29, 0xfa, 0x45, 0xba, 0xd9, 0x99, 0xeb, 0xa0, 0xbb, 0x5a, 0x51, 0xdf, 0x2a, 0x7e, 0xa3, 0xe6, 0x16, 0xed, 0x3d, 0xc7, 0xca, 0x8c, 0x93, 0xd8, 0x4b, 0x91, 0xa4, 0x00, 0x81, 0xc0, 0x75, 0xb6, 0x67, 0x64, 0x51, 0x32, 0x9f, 0xcc, 0xa6, 0xda, 0xff, 0xb6, 0xb7, 0xa4, 0x74},
			[56]byte{0xff, 0x0c, 0xcc, 0x2d, 0x5e, 0xa9, 0xcc, 0x6e, 0x3a, 0x99, 0x27, 0x45, 0x20, 0xc0, 0x62, 0xd2, 0x0e, 0x7b, 0x96, 0x4b, 0x93, 0xec, 0xb1, 0xe3, 0x68, 0xe4, 0x52, 0xeb, 0x7e, 0x72, 0x49, 0x6f, 0x44, 0x08, 0xf2, 0xa3, 0x0c, 0xd3, 0x92, 0x1a, 0xfa, 0x82, 0x31, 0xb1, 0x58, 0x03, 0x24, 0x78, 0x03, 0x19, 0x95, 0x8a, 0x85, 0x4f, 0x07, 0xfa},
			[56]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		),

		//1
		newTwistedPNiels(
			[56]byte{0x36, 0xaa, 0x95, 0x64, 0xca, 0x47, 0x1c, 0x57, 0x9f, 0xeb, 0x82, 0xa0, 0x26, 0x93, 0xe4, 0x25, 0xd4, 0x24, 0x46, 0x91, 0xb5, 0xf1, 0x36, 0x3f, 0x99, 0x55, 0x33, 0x5d, 0x58, 0x7a, 0xf4, 0x0e, 0x1e, 0xaf, 0x3f, 0x84, 0x2b, 0xba, 0x08, 0xd3, 0x8d, 0x79, 0x24, 0x14, 0xe9, 0x7b, 0xb9, 0x8e, 0x93, 0x33, 0x5f, 0x6b, 0x67, 0xd6, 0xb9, 0xc7},
			[56]byte{0x74, 0x4e, 0x5f, 0x3e, 0x46, 0x88, 0xc1, 0xea, 0xef, 0x52, 0x7d, 0xaf, 0x6d, 0x15, 0x71, 0xd0, 0xd7, 0x90, 0xbf, 0xee, 0xab, 0x0c, 0xbd, 0xd3, 0xbb, 0xb6, 0xf9, 0x2b, 0xa1, 0xdc, 0x31, 0xe7, 0x9e, 0x24, 0xe8, 0x90, 0xb4, 0x8f, 0x7d, 0x01, 0x37, 0x61, 0xac, 0xa2, 0x11, 0xbe, 0x3b, 0x87, 0xdb, 0x0c, 0xcc, 0xd5, 0x3f, 0x77, 0xd0, 0x58},
			[56]byte{0x93, 0xc4, 0xed, 0xca, 0xcd, 0x71, 0xfc, 0x6b, 0xf8, 0x3d, 0x30, 0x01, 0x8b, 0xc8, 0x5a, 0xfa, 0xd7, 0x3e, 0x4f, 0x7e, 0x53, 0x2e, 0x45, 0x3a, 0x5b, 0x08, 0xd7, 0xd6, 0x7f, 0xc2, 0xd6, 0xe3, 0xd1, 0xc5, 0xd5, 0xe0, 0x4d, 0xc8, 0xce, 0x67, 0xaf, 0x8a, 0x6c, 0x7b, 0x28, 0x20, 0x44, 0x4e, 0x81, 0xe6, 0xcb, 0xe9, 0xfc, 0x01, 0xdd, 0x82},
			[56]byte{0x34, 0x57, 0x2a, 0x29, 0xe9, 0xc2, 0xec, 0xba, 0xba, 0x42, 0x85, 0xc3, 0x96, 0xaa, 0x88, 0xf1, 0x8b, 0xbc, 0x50, 0xed, 0x16, 0x60, 0xe5, 0xae, 0xbf, 0xf1, 0xf4, 0x7c, 0x02, 0xc8, 0xd0, 0x2c, 0xc7, 0x82, 0xc0, 0x74, 0x98, 0xa3, 0x98, 0x52, 0x64, 0x73, 0xd5, 0x52, 0x11, 0xa6, 0x53, 0x1f, 0xda, 0x31, 0x25, 0x94, 0x4c, 0x68, 0x1c, 0x6e},
		),

		//2
		newTwistedPNiels(
			[56]byte{0xd9, 0x73, 0x1a, 0x12, 0x0c, 0xde, 0xb1, 0xcb, 0x20, 0x34, 0xbd, 0x0b, 0x3d, 0xf4, 0x34, 0x37, 0x1e, 0x3a, 0x13, 0x55, 0x07, 0x0e, 0xe4, 0x49, 0x80, 0xa4, 0x06, 0xa8, 0x44, 0x28, 0x15, 0xeb, 0xca, 0xd8, 0xbf, 0xf0, 0x1f, 0x4a, 0xaa, 0x65, 0x70, 0x7b, 0x9d, 0x13, 0xad, 0x56, 0xea, 0x33, 0x96, 0x7a, 0x47, 0xf9, 0x9c, 0x54, 0xa1, 0xdb},
			[56]byte{0x4a, 0x1a, 0xa1, 0x44, 0x3e, 0xb2, 0x05, 0x06, 0x61, 0x8f, 0xc4, 0x33, 0x2f, 0x03, 0x27, 0x76, 0x4d, 0x5f, 0x89, 0xac, 0xc9, 0x8a, 0xd7, 0x31, 0x24, 0xb7, 0x01, 0xc6, 0x6e, 0x0e, 0xb3, 0x2a, 0xaa, 0x59, 0x3d, 0x71, 0x04, 0x93, 0x62, 0xf9, 0xaa, 0x32, 0x46, 0x33, 0x36, 0xb6, 0x1a, 0xbb, 0xa9, 0xf1, 0xf5, 0x44, 0x9d, 0xb0, 0x8a, 0xa1},
			[56]byte{0x4e, 0xe7, 0xb1, 0x79, 0xcb, 0x6b, 0x8f, 0x01, 0xad, 0xa3, 0xc7, 0x97, 0xd1, 0x57, 0x0d, 0xa5, 0x1d, 0xbd, 0x9a, 0x37, 0x1a, 0x70, 0x99, 0xb6, 0xe5, 0xc6, 0xe0, 0x25, 0x86, 0x21, 0xc2, 0xec, 0xdf, 0x05, 0x4b, 0x78, 0x71, 0xf6, 0xe7, 0xdb, 0x40, 0xae, 0xb9, 0x89, 0x9a, 0x20, 0x9a, 0x68, 0xd4, 0x29, 0xee, 0x3f, 0x8a, 0xe3, 0xe5, 0x77},
			[56]byte{0xe7, 0xa9, 0x71, 0xf5, 0x59, 0x77, 0x3a, 0xab, 0x5d, 0x5f, 0xb1, 0x22, 0x20, 0xdf, 0x06, 0x71, 0xd5, 0x3e, 0x5a, 0x87, 0x7c, 0xe6, 0xe3, 0x3f, 0x48, 0x60, 0x8c, 0x61, 0x6f, 0x14, 0xa5, 0xd9, 0x33, 0xc8, 0x5a, 0x21, 0xed, 0xf7, 0xa5, 0xbf, 0x95, 0xf9, 0xee, 0x76, 0xfd, 0xb0, 0xdd, 0x74, 0x65, 0x2d, 0xb8, 0x74, 0x9f, 0x91, 0xa1, 0xdb},
		),

		//3
		newTwistedPNiels(
			[56]byte{0x70, 0x0b, 0x26, 0xa9, 0x93, 0xd4, 0xec, 0xee, 0x78, 0x66, 0x90, 0xfb, 0x09, 0xb3, 0x80, 0x1f, 0x47, 0x9a, 0x89, 0x35, 0x46, 0x21, 0x70, 0x9a, 0x48, 0x5b, 0x97, 0x86, 0x12, 0x85, 0x3f, 0x0c, 0x5f, 0x0c, 0x42, 0x74, 0x26, 0x79, 0x22, 0x96, 0xb0, 0xc4, 0x09, 0x54, 0xd8, 0x04, 0x73, 0x6a, 0x30, 0x70, 0xd6, 0x67, 0x76, 0xf9, 0xc6, 0x51},
			[56]byte{0x7f, 0xc4, 0x2a, 0xd0, 0x3d, 0x85, 0xeb, 0xf2, 0x03, 0xeb, 0x29, 0xd9, 0x09, 0x81, 0x9a, 0x27, 0x25, 0x92, 0x9e, 0xc6, 0xdc, 0xb5, 0xe2, 0x99, 0xd6, 0x3b, 0x92, 0xfc, 0x1f, 0x28, 0x90, 0xd6, 0xa8, 0x74, 0xee, 0x73, 0xfb, 0x53, 0x4b, 0x08, 0x50, 0x59, 0xf4, 0x31, 0x74, 0xd4, 0x23, 0xd5, 0x94, 0x7e, 0x3d, 0xa0, 0xe1, 0x0e, 0x06, 0x08},
			[56]byte{0x65, 0x06, 0x32, 0x4d, 0xc5, 0x5f, 0x3a, 0x61, 0xa5, 0xc7, 0x20, 0x09, 0x3f, 0x9b, 0x18, 0x27, 0xe0, 0x9f, 0x5c, 0xda, 0x74, 0x48, 0xb4, 0x5a, 0xdf, 0x1b, 0x07, 0x07, 0x86, 0x07, 0x24, 0xfa, 0x9b, 0x02, 0x41, 0x54, 0x74, 0xab, 0x7b, 0x9c, 0x33, 0x3b, 0xff, 0x22, 0x3c, 0xef, 0x7c, 0x73, 0x31, 0x3e, 0xab, 0x3a, 0x4c, 0x56, 0xdb, 0x39},
			[56]byte{0xe9, 0x49, 0x27, 0xff, 0xf0, 0xbb, 0x4e, 0xea, 0x01, 0xea, 0xa8, 0x54, 0x8a, 0xfe, 0x7f, 0xa7, 0x3c, 0x01, 0xc7, 0x87, 0x5e, 0x80, 0x0c, 0x01, 0x3e, 0xe7, 0x64, 0x24, 0x47, 0x9e, 0x4f, 0xd9, 0x06, 0x42, 0x63, 0x18, 0x1e, 0x05, 0xa8, 0x82, 0xde, 0x8d, 0x8a, 0x68, 0x2e, 0xaa, 0x31, 0x93, 0xee, 0xda, 0xe8, 0x31, 0x72, 0xc8, 0x58, 0x10},
		),

		//4
		newTwistedPNiels(
			[56]byte{0xd4, 0x3a, 0x2d, 0xda, 0x0c, 0x94, 0x10, 0xa0, 0xc7, 0x07, 0xa3, 0x11, 0x5a, 0xdf, 0x81, 0xcb, 0xc2, 0x6b, 0x80, 0x8a, 0x9f, 0xbb, 0xf8, 0xa4, 0xbf, 0x92, 0x8b, 0x63, 0x98, 0x59, 0xd5, 0x77, 0xca, 0x61, 0xd3, 0xb1, 0x60, 0x26, 0x3a, 0xc4, 0x4c, 0x32, 0x53, 0x4d, 0x9d, 0x62, 0x70, 0xc4, 0xd9, 0xfb, 0x5b, 0x68, 0xe6, 0xfe, 0x12, 0x76},
			[56]byte{0x96, 0x8b, 0x16, 0xc2, 0x86, 0x98, 0x1f, 0xef, 0x23, 0x9f, 0x74, 0xa8, 0x63, 0xd0, 0xde, 0x27, 0xa7, 0x84, 0x72, 0x32, 0x98, 0x23, 0x3f, 0x82, 0x22, 0x1a, 0x43, 0xdd, 0xfa, 0x68, 0x62, 0xce, 0x97, 0xf8, 0xdc, 0x96, 0x44, 0xf2, 0xaf, 0x83, 0x28, 0x29, 0x24, 0x13, 0xb7, 0x0b, 0xd7, 0x32, 0x65, 0x0c, 0xf8, 0x68, 0x47, 0x12, 0xc8, 0x62},
			[56]byte{0xee, 0x44, 0xe8, 0x7f, 0xd7, 0x10, 0xe9, 0x29, 0x50, 0xc4, 0x4b, 0x2e, 0xe7, 0x2a, 0x2e, 0xea, 0xc6, 0x7c, 0x18, 0xd1, 0x24, 0xff, 0x95, 0xb7, 0x75, 0x8b, 0x96, 0xe8, 0x96, 0xfd, 0x30, 0x57, 0xc4, 0x4a, 0x7e, 0x08, 0xd6, 0x55, 0x56, 0xb3, 0x42, 0x6d, 0x93, 0xea, 0x9a, 0x54, 0xd3, 0x49, 0xdc, 0xf4, 0x0d, 0xb4, 0x11, 0xf5, 0xf3, 0x1c},
			[56]byte{0x0f, 0xcb, 0xf8, 0xb1, 0xa3, 0x3b, 0xa8, 0x77, 0x36, 0x94, 0xaf, 0x47, 0x62, 0xfb, 0x25, 0xd6, 0x7a, 0x90, 0xf9, 0x77, 0xac, 0xd3, 0x20, 0x7b, 0x1b, 0x25, 0x82, 0x23, 0x83, 0xf9, 0xc7, 0xcd, 0x97, 0xe5, 0x5c, 0x5a, 0xbb, 0x6e, 0xf2, 0x5d, 0x62, 0x7f, 0x34, 0xed, 0x02, 0x35, 0xbb, 0x2d, 0x11, 0x40, 0x3f, 0xd2, 0xa8, 0xd8, 0x0f, 0x4d},
		),

		//5
		newTwistedPNiels(
			[56]byte{0x21, 0x84, 0xf1, 0xf6, 0xe4, 0x21, 0xac, 0x4b, 0xa9, 0x67, 0xe4, 0x91, 0x15, 0xcb, 0xb2, 0x02, 0x6c, 0xf2, 0x8b, 0x52, 0x3d, 0x15, 0x9b, 0x1c, 0xa5, 0x82, 0xf7, 0xea, 0x9b, 0x81, 0xf3, 0xb0, 0xa0, 0xa0, 0x4f, 0x52, 0x65, 0x58, 0x07, 0xf8, 0x94, 0x9c, 0x10, 0xa9, 0xcc, 0x93, 0x0d, 0xe0, 0xd6, 0x46, 0xcb, 0x3c, 0x36, 0xdf, 0x69, 0xf7},
			[56]byte{0x4d, 0xe1, 0xbe, 0x4b, 0x26, 0xe3, 0xad, 0x5e, 0x1b, 0x07, 0x31, 0xa2, 0x1b, 0xa5, 0x4c, 0xb7, 0x1c, 0xe2, 0xfb, 0xd4, 0x24, 0x68, 0xe6, 0x27, 0xbd, 0x07, 0x2e, 0x9e, 0x84, 0xe2, 0x21, 0x36, 0x9e, 0x8d, 0x31, 0x44, 0xf7, 0xee, 0x66, 0x88, 0xf5, 0xa3, 0xad, 0xa3, 0x60, 0x35, 0x32, 0x00, 0x10, 0xe9, 0x4e, 0xb8, 0x0f, 0xcf, 0x04, 0x65},
			[56]byte{0xfc, 0xdf, 0xea, 0x0b, 0x1a, 0x99, 0x16, 0x1f, 0xc0, 0x03, 0x6e, 0x2c, 0xc6, 0xca, 0x44, 0x3b, 0xf2, 0x24, 0x4d, 0x8a, 0x81, 0x99, 0x84, 0xfd, 0x61, 0x77, 0xa2, 0x54, 0x00, 0x86, 0x0b, 0xd6, 0xdb, 0x6f, 0xac, 0x63, 0x85, 0x8f, 0x14, 0x00, 0x9d, 0xb6, 0xa4, 0xda, 0x8f, 0xbd, 0x7b, 0x79, 0xd9, 0x49, 0x52, 0xa0, 0xe0, 0x62, 0x88, 0xf0},
			[56]byte{0x45, 0x88, 0xc8, 0xb2, 0x29, 0x7d, 0x53, 0xdf, 0x48, 0x94, 0xaf, 0x73, 0x19, 0xcf, 0xa4, 0xd1, 0xd9, 0xa9, 0x82, 0x26, 0xfd, 0x9b, 0x36, 0x6d, 0xd4, 0x99, 0x10, 0x6d, 0x32, 0x3e, 0x64, 0x80, 0xe3, 0x9c, 0xc7, 0x18, 0xea, 0x67, 0x90, 0x04, 0x5e, 0x83, 0x8d, 0xa0, 0xc5, 0x99, 0x36, 0xba, 0x09, 0x69, 0x07, 0xb7, 0xcd, 0x7e, 0x71, 0x27},
		),

		//6
		newTwistedPNiels(
			[56]byte{0x2a, 0xd6, 0xfc, 0xfc, 0x10, 0x7d, 0x16, 0xaf, 0x7f, 0x08, 0xca, 0xbb, 0xb8, 0xa6, 0xca, 0xcf, 0xb4, 0x81, 0x46, 0xa0, 0xe3, 0xd0, 0x94, 0xc0, 0xb6, 0xb4, 0x76, 0x78, 0x7d, 0x3f, 0xb3, 0xed, 0xcb, 0x3c, 0x4e, 0x16, 0xb7, 0x1c, 0x17, 0x65, 0x4d, 0x6c, 0xaa, 0x72, 0x6b, 0x7f, 0xdf, 0xea, 0xc4, 0x21, 0xa9, 0xaf, 0xb8, 0x31, 0xc9, 0x12},
			[56]byte{0xc6, 0x07, 0xf8, 0xdb, 0x1b, 0x24, 0x3c, 0x9b, 0xa8, 0xd1, 0x03, 0xb9, 0x10, 0xbb, 0xd2, 0x38, 0xd9, 0xff, 0x53, 0xc8, 0x1e, 0x41, 0x01, 0x4b, 0xfa, 0xcf, 0x56, 0xbb, 0xd8, 0xa3, 0x3c, 0x82, 0x0f, 0x7b, 0xa5, 0x35, 0x2a, 0x56, 0x3e, 0x76, 0x17, 0x23, 0x8b, 0xbf, 0xe9, 0x70, 0x71, 0x56, 0x06, 0xc1, 0x7e, 0x48, 0xb7, 0x60, 0x1e, 0xcf},
			[56]byte{0x68, 0x98, 0x0a, 0x0e, 0x51, 0xe6, 0xcb, 0x23, 0x99, 0x9a, 0x60, 0xa6, 0xac, 0x45, 0xd9, 0x15, 0xdd, 0x22, 0xe8, 0x4d, 0x61, 0x47, 0x56, 0x8d, 0x7a, 0xad, 0x43, 0xba, 0xcc, 0x11, 0xaf, 0x50, 0xd3, 0xd8, 0x90, 0x5f, 0xbd, 0x51, 0x2a, 0x4a, 0xe4, 0x60, 0x40, 0x23, 0x5b, 0x79, 0x2c, 0xf7, 0x5b, 0xdc, 0xc7, 0x6c, 0xea, 0x92, 0xd1, 0x27},
			[56]byte{0x4c, 0x79, 0xe0, 0x2c, 0x72, 0x2d, 0xcc, 0x29, 0x3f, 0xc4, 0x22, 0x9a, 0x5a, 0x98, 0xbf, 0x77, 0x2b, 0x7b, 0x18, 0x7a, 0x4d, 0xa4, 0x4c, 0xe7, 0x47, 0xb9, 0x7a, 0xd2, 0x2d, 0x3f, 0xae, 0x78, 0x7f, 0x56, 0x23, 0x49, 0xcb, 0xc4, 0x98, 0x11, 0xd6, 0x79, 0xc5, 0x72, 0x84, 0x82, 0x56, 0x40, 0xfd, 0xc8, 0x52, 0x51, 0x37, 0x98, 0xea, 0x89},
		),

		//7
		newTwistedPNiels(
			[56]byte{0x27, 0x85, 0xf8, 0xf1, 0x06, 0x11, 0x10, 0xa2, 0xcc, 0x34, 0xb0, 0x42, 0x96, 0xe2, 0xa2, 0x36, 0xbb, 0xa0, 0xda, 0x86, 0xdd, 0x45, 0x78, 0x0c, 0xdb, 0xe3, 0x5c, 0x3d, 0x5d, 0xda, 0x75, 0xca, 0x5f, 0x58, 0xe7, 0x91, 0x75, 0x56, 0xe9, 0xd9, 0x37, 0xd0, 0xd5, 0xf2, 0x75, 0x75, 0xb1, 0xe2, 0x7b, 0x47, 0x94, 0xe0, 0x89, 0xc7, 0x65, 0xf7},
			[56]byte{0xbb, 0x72, 0x9b, 0xcd, 0xc7, 0x95, 0x1d, 0x45, 0x10, 0x17, 0x12, 0xbb, 0x5f, 0x27, 0x2b, 0x18, 0xda, 0x55, 0x09, 0xe2, 0x96, 0x87, 0x86, 0xdc, 0xec, 0x73, 0x8d, 0x63, 0x94, 0x3b, 0xfc, 0x25, 0xd4, 0xbc, 0x1a, 0xf3, 0x73, 0xf7, 0xb0, 0x80, 0xfe, 0xe7, 0x9a, 0xcf, 0x52, 0x9d, 0xc5, 0x5e, 0x36, 0x5f, 0x0b, 0xe7, 0x56, 0x4e, 0x86, 0x05},
			[56]byte{0x17, 0x2a, 0xc7, 0xf3, 0x25, 0xe9, 0x9f, 0x7c, 0xa4, 0x81, 0x4c, 0x60, 0xb1, 0xa0, 0xfd, 0xb6, 0x7a, 0x0b, 0x63, 0x5b, 0xe4, 0x4d, 0x08, 0xa5, 0x12, 0x8a, 0x09, 0x60, 0x53, 0x75, 0xa7, 0x00, 0xbd, 0x1a, 0xd0, 0x23, 0x82, 0x5b, 0x99, 0x7f, 0x45, 0x4a, 0x44, 0x3a, 0xa9, 0xe6, 0x78, 0x2c, 0x4c, 0x10, 0xd8, 0xf2, 0xc0, 0x31, 0xc3, 0xa1},
			[56]byte{0xf1, 0x7e, 0x0d, 0x5f, 0x9c, 0xd2, 0x39, 0x09, 0x4e, 0x07, 0x15, 0x17, 0x93, 0x3a, 0xc5, 0x01, 0x01, 0x2f, 0x29, 0x58, 0xc7, 0x84, 0x9f, 0xa8, 0x29, 0x83, 0xff, 0xf6, 0x24, 0x29, 0x8b, 0x16, 0x60, 0x17, 0xb9, 0xd5, 0x2f, 0xc9, 0xbb, 0x88, 0x5e, 0xd9, 0x17, 0x97, 0x7d, 0xca, 0x15, 0xc7, 0x2b, 0xbc, 0x54, 0x50, 0x0c, 0x6c, 0x0e, 0x73},
		),

		//8
		newTwistedPNiels(
			[56]byte{0xd9, 0xe0, 0xd3, 0xed, 0x61, 0x80, 0xcc, 0x10, 0x35, 0xb4, 0x07, 0x27, 0x44, 0xed, 0xd4, 0x08, 0x63, 0x92, 0x59, 0x2d, 0x8d, 0x54, 0xa8, 0xe0, 0x4a, 0xa1, 0x13, 0x03, 0xf5, 0xb6, 0xfe, 0xf2, 0xe3, 0xe5, 0xaf, 0x89, 0xc6, 0xa6, 0xba, 0x3b, 0xa0, 0xed, 0x1d, 0x3f, 0x07, 0xe6, 0x42, 0x7e, 0x94, 0x07, 0x02, 0x73, 0x52, 0x0e, 0x3e, 0xc8},
			[56]byte{0x7f, 0xe8, 0x10, 0x14, 0xd8, 0x32, 0x36, 0x48, 0x5b, 0x6e, 0x4a, 0x6d, 0xc5, 0xcf, 0x9e, 0xe9, 0x15, 0xc2, 0xcf, 0xe4, 0x36, 0xe1, 0x31, 0x7c, 0x2d, 0x26, 0x82, 0x0e, 0x7e, 0xe9, 0xef, 0x99, 0x11, 0xfb, 0xbd, 0x83, 0x12, 0x5d, 0x45, 0x16, 0xeb, 0xfc, 0x51, 0x40, 0x1a, 0x9e, 0x04, 0xf9, 0x52, 0x35, 0xc0, 0xff, 0x90, 0x6e, 0x96, 0x31},
			[56]byte{0x6f, 0xfd, 0xde, 0x6b, 0x9b, 0x7d, 0x8c, 0xfb, 0xa5, 0x9d, 0x2d, 0x42, 0x31, 0xc7, 0xc5, 0x70, 0xd8, 0xc4, 0x3c, 0x20, 0xfb, 0x7f, 0x8d, 0x59, 0xce, 0xb5, 0x0d, 0x36, 0x04, 0x8d, 0xa2, 0xee, 0xee, 0xe3, 0xb1, 0x4d, 0x7a, 0x33, 0x9a, 0xde, 0xb3, 0x9c, 0xc2, 0x6c, 0x93, 0x4d, 0x17, 0xf3, 0x58, 0x17, 0x75, 0x71, 0x3a, 0x0e, 0xec, 0x8f},
			[56]byte{0x74, 0x35, 0xad, 0xe2, 0xbf, 0x3a, 0x9b, 0x06, 0x6a, 0xab, 0xde, 0x20, 0x5f, 0x69, 0x30, 0x5d, 0xcd, 0xf9, 0xe6, 0x7f, 0xdb, 0x89, 0x2d, 0x7f, 0x48, 0x9f, 0xb2, 0x6e, 0x14, 0x5c, 0xf0, 0xbd, 0xd0, 0xa4, 0xbf, 0x59, 0x57, 0xd0, 0x99, 0xd5, 0x9b, 0x2c, 0x32, 0x72, 0x47, 0xd3, 0x31, 0x54, 0x71, 0x2d, 0x9d, 0x66, 0x48, 0x85, 0xa5, 0xee},
		),

		//9
		newTwistedPNiels(
			[56]byte{0xd6, 0x70, 0x1e, 0xb5, 0xb7, 0xda, 0x6f, 0x3d, 0xc0, 0x12, 0x9e, 0xda, 0xd6, 0x41, 0x7f, 0xe4, 0x1f, 0xcb, 0x6a, 0x68, 0xac, 0x55, 0x3b, 0x15, 0xe9, 0x42, 0x4c, 0xe3, 0x6d, 0x9b, 0x33, 0x61, 0x35, 0xf8, 0x88, 0x42, 0x75, 0x69, 0x35, 0xf8, 0x37, 0x6a, 0xba, 0x72, 0xec, 0xb8, 0xd6, 0x5f, 0x17, 0xf1, 0xcf, 0xc0, 0x34, 0x4f, 0x8d, 0x6c},
			[56]byte{0xbc, 0x0f, 0x44, 0x15, 0x06, 0xc5, 0x35, 0x57, 0x1a, 0x0b, 0x18, 0xed, 0x4d, 0x57, 0x46, 0x8b, 0x25, 0x46, 0x30, 0xee, 0xdb, 0x91, 0x05, 0xb5, 0xf7, 0xb5, 0xf1, 0x9d, 0xdd, 0x1f, 0x89, 0x66, 0x30, 0x94, 0x20, 0xb5, 0x8a, 0xf7, 0x1e, 0x66, 0x70, 0x9f, 0x94, 0xcc, 0xca, 0xc5, 0x12, 0x38, 0x65, 0x10, 0x93, 0x93, 0x72, 0xa2, 0x76, 0xc1},
			[56]byte{0x20, 0xa8, 0xe3, 0xb5, 0x1f, 0x00, 0x47, 0x37, 0xf4, 0xc8, 0x78, 0x31, 0x08, 0x97, 0x66, 0xce, 0xa5, 0x50, 0x03, 0xcb, 0x59, 0x61, 0x58, 0x17, 0xe9, 0xb9, 0xaa, 0x72, 0xe7, 0x10, 0xe0, 0x71, 0x8a, 0x19, 0x93, 0x8f, 0x1f, 0x47, 0x59, 0x24, 0x1f, 0xc4, 0xff, 0x6c, 0x53, 0xa6, 0x2f, 0x58, 0x56, 0xf5, 0x29, 0xd5, 0xce, 0x4b, 0x92, 0x95},
			[56]byte{0x69, 0x4e, 0x97, 0x87, 0x5a, 0x9a, 0x92, 0x33, 0xd0, 0x75, 0x0c, 0x57, 0xed, 0x37, 0x79, 0x9a, 0xd2, 0x7e, 0xbe, 0x6b, 0x1b, 0xba, 0xe1, 0x69, 0x1a, 0x31, 0xea, 0xab, 0xa3, 0x28, 0x7d, 0x0d, 0x07, 0x1f, 0x4e, 0x17, 0xe9, 0x1e, 0x6c, 0x3d, 0x14, 0xea, 0xfa, 0x52, 0x02, 0x29, 0xc6, 0xf1, 0xdf, 0xc7, 0x2f, 0xfa, 0x64, 0x9e, 0xcf, 0xc1},
		),

		//10
		newTwistedPNiels(
			[56]byte{0x5c, 0x93, 0x5e, 0x68, 0x6d, 0x43, 0xd5, 0x9a, 0xdc, 0x5d, 0x9d, 0x82, 0x8c, 0x2a, 0x25, 0x1f, 0xdc, 0x6d, 0x87, 0xbc, 0x0f, 0xe4, 0x2f, 0x86, 0x7b, 0x7b, 0xc1, 0xf0, 0x1c, 0xa4, 0x11, 0x08, 0xa0, 0xc0, 0x39, 0x58, 0x45, 0x65, 0x4e, 0x46, 0x18, 0xd0, 0x23, 0x58, 0xbc, 0xd2, 0x1c, 0xd0, 0xa9, 0xe5, 0x18, 0xe1, 0xe5, 0xeb, 0x30, 0xd0},
			[56]byte{0x45, 0xed, 0x7b, 0xef, 0x3a, 0xb1, 0x04, 0x60, 0x3e, 0xda, 0xf4, 0x3f, 0xe7, 0xcb, 0xa8, 0x21, 0x91, 0xa9, 0x1b, 0xef, 0x40, 0xc3, 0x03, 0x40, 0x58, 0x54, 0xd8, 0x9f, 0xb3, 0xf7, 0x7f, 0x56, 0x15, 0x14, 0xda, 0x80, 0x65, 0x52, 0x05, 0x7d, 0x9b, 0xae, 0xdc, 0x92, 0xc1, 0xfc, 0xdc, 0x01, 0x5d, 0x2c, 0xa8, 0x58, 0x05, 0xce, 0xc6, 0xf6},
			[56]byte{0x85, 0xd8, 0x28, 0x66, 0x15, 0x5d, 0x89, 0x7a, 0x38, 0xa3, 0x81, 0x28, 0xd6, 0x25, 0x11, 0x52, 0x2c, 0x53, 0x7f, 0x00, 0x27, 0x56, 0x4f, 0x11, 0x33, 0x32, 0x74, 0x0c, 0x67, 0xec, 0x3e, 0x9c, 0x21, 0x5e, 0xd4, 0x1c, 0x2a, 0xfe, 0x60, 0x16, 0x5c, 0x2a, 0x6d, 0x93, 0x6b, 0xe9, 0xbc, 0x11, 0x05, 0x11, 0x8e, 0x01, 0x71, 0xd8, 0x01, 0x55},
			[56]byte{0x04, 0x2c, 0x65, 0x63, 0x32, 0x3f, 0xbd, 0x64, 0x91, 0xf2, 0x9a, 0xdf, 0xae, 0xbc, 0xa9, 0x74, 0x51, 0xba, 0xf3, 0xe0, 0xab, 0xd2, 0xe7, 0xc3, 0xc5, 0xda, 0xec, 0x4e, 0xa0, 0xe6, 0x01, 0xb6, 0x9c, 0xb0, 0xd9, 0x85, 0x11, 0x3e, 0xc7, 0x10, 0xd2, 0xc7, 0x68, 0x11, 0xda, 0x70, 0x63, 0x22, 0x44, 0x7b, 0x83, 0x8d, 0xb0, 0x97, 0x59, 0x4f},
		),

		//11
		newTwistedPNiels(
			[56]byte{0x92, 0x91, 0xed, 0x7f, 0xb8, 0x11, 0x83, 0xd6, 0x0a, 0x4f, 0xa7, 0x39, 0xf2, 0x78, 0xc1, 0xf5, 0x46, 0xd1, 0xfa, 0x01, 0x6f, 0xdb, 0xbe, 0xd3, 0xeb, 0xd5, 0xa9, 0xb9, 0x03, 0x13, 0xe9, 0x8f, 0x9f, 0x17, 0x0c, 0x14, 0x12, 0x18, 0x07, 0x0a, 0x95, 0xed, 0xf6, 0xd9, 0x30, 0x8d, 0x12, 0x90, 0x75, 0x50, 0xbb, 0x4b, 0x8f, 0xac, 0x49, 0xad},
			[56]byte{0x2a, 0x7b, 0xdd, 0xf5, 0xdd, 0x1e, 0xad, 0xf4, 0x70, 0x9f, 0xf3, 0x1f, 0x69, 0xa1, 0xca, 0x7f, 0xf9, 0xcb, 0x53, 0x4b, 0x94, 0xc2, 0x3f, 0xa5, 0x0f, 0x7d, 0x1b, 0x6a, 0x6c, 0x21, 0x1a, 0xf7, 0x50, 0x11, 0x76, 0xd9, 0x49, 0xbb, 0xc7, 0x73, 0x5f, 0x48, 0x09, 0x40, 0xe6, 0x37, 0xb6, 0x5c, 0x7e, 0x41, 0x80, 0x2d, 0x63, 0x00, 0x66, 0x82},
			[56]byte{0x8a, 0xd6, 0x1f, 0x0a, 0x1b, 0x91, 0x55, 0xc5, 0x21, 0x8c, 0xa2, 0x03, 0x8d, 0x82, 0x05, 0x37, 0xf1, 0xeb, 0x6a, 0xd9, 0xc0, 0x03, 0x39, 0xda, 0x45, 0x93, 0x3a, 0x3b, 0x10, 0xb8, 0xa8, 0xa9, 0x33, 0xd8, 0x84, 0x70, 0x63, 0x1c, 0x26, 0x12, 0xac, 0x2e, 0x87, 0xbd, 0x3d, 0xfe, 0xae, 0x14, 0xb9, 0x57, 0xf0, 0x31, 0x49, 0xb2, 0x93, 0x3e},
			[56]byte{0x89, 0x1e, 0x0a, 0xbb, 0x40, 0xf9, 0x5e, 0x7c, 0xb3, 0xff, 0x32, 0x43, 0xcd, 0x20, 0x11, 0x26, 0x3e, 0xc1, 0x20, 0x85, 0xc7, 0xdf, 0x18, 0xab, 0xa8, 0x6d, 0xb0, 0xd6, 0xad, 0x41, 0xa7, 0xe1, 0x75, 0xb9, 0x6f, 0x6d, 0x29, 0xdd, 0x82, 0x3d, 0x7c, 0x59, 0xe7, 0x8e, 0x53, 0xad, 0x27, 0x8f, 0x10, 0xcf, 0x36, 0x30, 0x47, 0xf5, 0x02, 0xb3},
		),

		//12
		newTwistedPNiels(
			[56]byte{0xff, 0xca, 0xcf, 0x02, 0x66, 0x70, 0x60, 0x2f, 0x4e, 0xc4, 0x31, 0x4e, 0xb8, 0x41, 0x43, 0x5b, 0x61, 0x51, 0x42, 0xfd, 0x8e, 0x78, 0xf7, 0x00, 0x9d, 0x17, 0x42, 0xda, 0x2f, 0xfb, 0xe2, 0x90, 0x4d, 0x6f, 0x81, 0xb0, 0x24, 0x3e, 0x58, 0xf4, 0xa3, 0xcc, 0x55, 0x5d, 0x42, 0x00, 0xdf, 0x53, 0xfa, 0x88, 0xfc, 0x09, 0x22, 0x8f, 0xa6, 0xde},
			[56]byte{0x36, 0xff, 0xaf, 0x76, 0xf6, 0xbd, 0x72, 0x31, 0x88, 0x95, 0x7c, 0x65, 0xee, 0xde, 0x7c, 0x9a, 0xc0, 0x0b, 0x2f, 0x01, 0x65, 0x8c, 0xbd, 0x1f, 0x55, 0xa2, 0x5c, 0x0d, 0xac, 0x8e, 0x91, 0x06, 0x22, 0xd5, 0xf5, 0x42, 0xae, 0xb9, 0x45, 0xe3, 0xdf, 0xd4, 0xca, 0x87, 0x20, 0x24, 0x0b, 0x69, 0xd4, 0xd9, 0x89, 0xc3, 0x8d, 0x1e, 0x4f, 0xfd},
			[56]byte{0x17, 0x1e, 0x96, 0x11, 0x6e, 0x04, 0x7e, 0x38, 0x59, 0x73, 0x34, 0xcc, 0xd8, 0x6c, 0xec, 0xb2, 0xa1, 0x2f, 0x4c, 0x06, 0x79, 0xbd, 0x12, 0xe8, 0x20, 0x57, 0x21, 0xeb, 0x9f, 0xbe, 0xec, 0x01, 0xde, 0x9e, 0x3c, 0x93, 0xa3, 0x7f, 0xfc, 0xbe, 0x0e, 0x3a, 0xc7, 0xa9, 0xe1, 0x14, 0x3c, 0x72, 0x4c, 0x1d, 0x4b, 0xd6, 0x53, 0x1a, 0xe2, 0x96},
			[56]byte{0x99, 0xe3, 0x76, 0x4f, 0x04, 0x88, 0xcd, 0x7e, 0x4f, 0x76, 0x1f, 0x62, 0x53, 0x74, 0x18, 0xcc, 0x60, 0x27, 0xf9, 0x96, 0x87, 0x72, 0xe1, 0xd4, 0x74, 0x7e, 0x59, 0x48, 0x5a, 0xec, 0x11, 0xda, 0x5e, 0x17, 0x46, 0x05, 0x79, 0xa6, 0x59, 0xc4, 0x5c, 0x2b, 0xf4, 0x75, 0x49, 0xe8, 0xdb, 0xf4, 0xd3, 0xe0, 0x68, 0xd0, 0xf2, 0x89, 0x65, 0x0b},
		),

		//13
		newTwistedPNiels(
			[56]byte{0xbf, 0x0f, 0xcf, 0xaf, 0xcc, 0xdd, 0x65, 0x3c, 0x0c, 0x03, 0xa1, 0x6c, 0x2f, 0x7c, 0x27, 0xde, 0x4e, 0xec, 0x5f, 0x4f, 0x19, 0xaf, 0xf6, 0x31, 0x9b, 0x79, 0xf9, 0x42, 0x6e, 0x98, 0xa2, 0xfb, 0x8a, 0xab, 0x92, 0xb3, 0x57, 0x83, 0x51, 0x87, 0xff, 0xfd, 0x0b, 0xe4, 0xfe, 0xc5, 0x54, 0x9e, 0x4f, 0x57, 0xb5, 0x2b, 0x3b, 0x03, 0x9d, 0xe6},
			[56]byte{0x39, 0x2e, 0x8d, 0xf1, 0x70, 0xcc, 0x11, 0xf9, 0xe2, 0x4b, 0x70, 0x30, 0x9d, 0x2a, 0x8c, 0x60, 0x7a, 0x2e, 0x76, 0xc6, 0x7d, 0xd8, 0xe1, 0x74, 0xe1, 0xe5, 0xc6, 0x0a, 0x81, 0x29, 0xf8, 0x9e, 0xf4, 0x86, 0xab, 0x0c, 0x59, 0xf1, 0x7a, 0x6c, 0xe4, 0xc3, 0xa5, 0x1f, 0x47, 0x7f, 0xf1, 0xec, 0xb8, 0xde, 0xd7, 0xcd, 0x86, 0x4c, 0x9d, 0xf8},
			[56]byte{0x0d, 0xd3, 0xb8, 0xe0, 0x2e, 0xd1, 0x6e, 0xb0, 0x8b, 0x10, 0xf0, 0x39, 0x5c, 0xa8, 0xda, 0xad, 0x99, 0x55, 0xea, 0x53, 0xc7, 0x7d, 0x82, 0x81, 0x32, 0x59, 0xad, 0x20, 0xe0, 0xe9, 0x8a, 0x5a, 0xe6, 0x69, 0x2f, 0x67, 0x31, 0x7d, 0x31, 0x47, 0xa1, 0x0f, 0x28, 0xb1, 0x49, 0x48, 0x6a, 0x70, 0xec, 0x65, 0x26, 0x59, 0x68, 0xa1, 0x61, 0xd9},
			[56]byte{0x8d, 0xaa, 0x22, 0x6e, 0x35, 0xc4, 0x76, 0x57, 0x8b, 0x33, 0x2b, 0xcf, 0x96, 0xc9, 0x3f, 0x11, 0x39, 0xed, 0x41, 0x8e, 0x54, 0x78, 0x44, 0x22, 0xf8, 0xd2, 0xfd, 0xa3, 0x13, 0x73, 0x74, 0xb8, 0x67, 0xaf, 0xdf, 0x91, 0xf7, 0x68, 0x20, 0xe9, 0xe2, 0x26, 0x52, 0x1e, 0x5e, 0xff, 0xfd, 0xbf, 0xea, 0x3e, 0x4c, 0x82, 0xbb, 0x0d, 0x00, 0xb0},
		),

		//14
		newTwistedPNiels(
			[56]byte{0x1c, 0xa9, 0xb7, 0xca, 0x99, 0xb9, 0x5d, 0x40, 0x89, 0xb7, 0x34, 0xd5, 0xa3, 0x7b, 0xd8, 0xd2, 0x5a, 0x88, 0xf9, 0x39, 0x5e, 0xa8, 0x6a, 0x5f, 0x5d, 0xb4, 0x4c, 0xa5, 0xfd, 0xbb, 0xf2, 0x3c, 0x01, 0x2e, 0x15, 0x11, 0x2d, 0xd5, 0xb3, 0xff, 0xb8, 0x87, 0x62, 0xbf, 0x68, 0x66, 0xc5, 0x69, 0xca, 0x48, 0x06, 0xe2, 0x4b, 0x0a, 0xbe, 0xea},
			[56]byte{0x3b, 0x25, 0xe4, 0xe7, 0x1a, 0xfe, 0x2e, 0xf2, 0xf7, 0x5e, 0xf2, 0xd9, 0x5c, 0x86, 0x79, 0xf0, 0x26, 0xea, 0xb1, 0x91, 0xa9, 0x60, 0xa1, 0xda, 0xac, 0x19, 0xa2, 0xc4, 0x83, 0xb1, 0xba, 0x9e, 0xaa, 0xb5, 0x1e, 0xc0, 0xb4, 0x3e, 0x97, 0x64, 0xbe, 0x35, 0x78, 0x87, 0xb8, 0xf2, 0x59, 0x97, 0x52, 0xf1, 0xd4, 0x3b, 0x6b, 0x1b, 0x6c, 0xb0},
			[56]byte{0x9b, 0xdd, 0x6e, 0xc1, 0x13, 0x7e, 0x05, 0x3a, 0x85, 0xe8, 0x9b, 0x54, 0xff, 0xd9, 0x78, 0xbc, 0x3f, 0x61, 0xc4, 0xae, 0x30, 0x35, 0x04, 0x35, 0xf8, 0xca, 0x4e, 0x14, 0xf8, 0xea, 0xba, 0xf7, 0xd4, 0xc9, 0x82, 0x39, 0x9b, 0x58, 0xd0, 0x45, 0x24, 0xaf, 0x26, 0x46, 0x98, 0x91, 0x57, 0xfe, 0x55, 0xa9, 0x32, 0x45, 0x13, 0x75, 0x0a, 0x7d},
			[56]byte{0x32, 0xad, 0xdc, 0xe9, 0xcb, 0x48, 0x9e, 0x53, 0x7e, 0x19, 0x16, 0x7c, 0xf7, 0x2e, 0xcd, 0xd2, 0x8e, 0x2e, 0x12, 0xec, 0xbc, 0x94, 0x89, 0xdc, 0x41, 0x07, 0x94, 0x20, 0xb4, 0x29, 0x6f, 0x16, 0x1d, 0x2c, 0x8a, 0xad, 0xe9, 0xf8, 0xae, 0x01, 0x47, 0x06, 0x31, 0x88, 0xe1, 0xe4, 0xf1, 0x9a, 0x57, 0x53, 0x2e, 0xac, 0xbb, 0x11, 0x9a, 0x13},
		),

		//15
		newTwistedPNiels(
			[56]byte{0xcd, 0xb6, 0x4c, 0x2d, 0x35, 0x80, 0x1d, 0x08, 0xfd, 0xe5, 0xd3, 0x05, 0xfc, 0xe7, 0x02, 0x4f, 0x99, 0x7d, 0x2c, 0x0c, 0xed, 0x8c, 0x3f, 0x30, 0x1b, 0xac, 0xc3, 0x39, 0x87, 0xb3, 0x8a, 0x54, 0xa8, 0x33, 0x37, 0xa7, 0xa9, 0x38, 0x84, 0x8a, 0x27, 0xc1, 0x7a, 0xd0, 0xdd, 0x83, 0x37, 0xc1, 0xf2, 0x27, 0xf2, 0x40, 0xfd, 0xf0, 0x58, 0x3d},
			[56]byte{0xe7, 0x73, 0x8c, 0x78, 0x94, 0x10, 0xa0, 0xac, 0x6c, 0x5f, 0x47, 0xf9, 0xd8, 0xe9, 0xaa, 0xda, 0xa4, 0x67, 0x34, 0x38, 0xe9, 0x2b, 0xfb, 0x4f, 0x5c, 0xe6, 0xbb, 0x92, 0x4f, 0x21, 0x33, 0xc0, 0x28, 0xe3, 0xd0, 0x7c, 0xcb, 0x7d, 0xc6, 0x33, 0x45, 0x99, 0xf1, 0x55, 0x9b, 0x43, 0xcd, 0xec, 0xb1, 0xf0, 0x1f, 0x49, 0x60, 0x0a, 0xb8, 0x3b},
			[56]byte{0x04, 0xf6, 0xfd, 0x1b, 0x2f, 0xc1, 0x3b, 0x10, 0x1e, 0x61, 0x52, 0x66, 0x22, 0xa5, 0x71, 0xe8, 0x93, 0x9d, 0xfe, 0xe5, 0x3d, 0xaa, 0xaf, 0x15, 0x5d, 0x24, 0x66, 0x13, 0xa7, 0x43, 0x6c, 0xdc, 0x74, 0xd9, 0xce, 0xc5, 0xcc, 0x85, 0xca, 0xf7, 0x14, 0x12, 0xe2, 0x63, 0x2d, 0x9d, 0x28, 0x6f, 0x75, 0x2a, 0x4b, 0x0e, 0xfe, 0x3a, 0x83, 0x93},
			[56]byte{0xf1, 0xc9, 0x1f, 0x03, 0x8d, 0x8b, 0x98, 0xf2, 0x32, 0x5f, 0x00, 0xd1, 0xfb, 0xe4, 0x69, 0x09, 0xda, 0x16, 0x2a, 0x59, 0x91, 0xab, 0xe0, 0x3e, 0xc3, 0xa0, 0xac, 0x7c, 0x25, 0xee, 0x86, 0xf8, 0x09, 0x22, 0x2a, 0xb0, 0xa6, 0xf9, 0x1c, 0x43, 0x48, 0xa3, 0xc2, 0x7c, 0x70, 0xec, 0x6b, 0xda, 0xd4, 0x0a, 0x45, 0x51, 0x66, 0xc1, 0x57, 0xac},
		),
	}

	dst := make([]*twPNiels, 1<<tableSize)

	px, _ := hex.DecodeString("4d8b77dc973a1f9bcd5358c702ee8159a71cd3e4c1ff95bfb30e7038cffe9f794211dffd758e2a2a693a08a9a454398fde981e5e2669acad")
	py, _ := hex.DecodeString("27193fda68a08730d1def89d64c7f466d9e3d0ac89d8fdcd17b8cdb446e80404e8cd715d4612c16f70803d50854b66c9b3412e85e2f19b0d")
	pz, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001")
	pt, _ := hex.DecodeString("4d8b77dc973a1f9bcd5358c702ee8159a71cd3e4c1ff95bfb30e7038cffe9f794211dffd758e2a2a693a08a9a454398fde981e5e2669acad")
	pu, _ := hex.DecodeString("27193fda68a08730d1def89d64c7f466d9e3d0ac89d8fdcd17b8cdb446e80404e8cd715d4612c16f70803d50854b66c9b3412e85e2f19b0d")

	p := &twExtensible{
		new(bigNumber).setBytes(px),
		new(bigNumber).setBytes(py),
		new(bigNumber).setBytes(pz),
		new(bigNumber).setBytes(pt),
		new(bigNumber).setBytes(pu),
	}

	prepareWnafTable(dst, p, tableSize)

	for i, di := range dst {
		c.Assert(di.equals(expected[i]), Equals, true)
	}
}

func (s *Ed448Suite) TestWNAFSMultiplication(c *C) {
	c.Skip("not yet, hold your horses!")
	px, _ := hex.DecodeString("4d8b77dc973a1f9bcd5358c702ee8159a71cd3e4c1ff95bfb30e7038cffe9f794211dffd758e2a2a693a08a9a454398fde981e5e2669acad")
	py, _ := hex.DecodeString("27193fda68a08730d1def89d64c7f466d9e3d0ac89d8fdcd17b8cdb446e80404e8cd715d4612c16f70803d50854b66c9b3412e85e2f19b0d")
	pz, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001")
	pt, _ := hex.DecodeString("4d8b77dc973a1f9bcd5358c702ee8159a71cd3e4c1ff95bfb30e7038cffe9f794211dffd758e2a2a693a08a9a454398fde981e5e2669acad")
	pu, _ := hex.DecodeString("27193fda68a08730d1def89d64c7f466d9e3d0ac89d8fdcd17b8cdb446e80404e8cd715d4612c16f70803d50854b66c9b3412e85e2f19b0d")

	p := &twExtensible{
		new(bigNumber).setBytes(px),
		new(bigNumber).setBytes(py),
		new(bigNumber).setBytes(pz),
		new(bigNumber).setBytes(pt),
		new(bigNumber).setBytes(pu),
	}

	x := [fieldWords]word_t{
		0x6c226d73, 0x70edcfc3,
		0x44156c47, 0x084f4695,
		0xe72606ac, 0x9d0ce5e5,
		0xed96d3ba, 0x9ff3fa11,
		0x4a15c383, 0xca38a0af,
		0xead789b3, 0xb96613ba,
		0x48ba4461, 0x34eb2031,
	}

	y := [fieldWords]word_t{
		0x2118b8c6, 0x4356acd5,
		0x26d7e73c, 0x459174b7,
		0xf10bea31, 0x83e528bb,
		0xb960d695, 0xd0da7e28,
		0xbad7f9a1, 0xe9f5ba01,
		0x94ea1518, 0x12c58cca,
		0x302c76eb, 0x3bd0363e,
	}

	px, _ = hex.DecodeString("d902fadbeee8dd1ef391dcce59cc75d286c9efc7229dd919a35236a5447384e84617bf94d4129af02d7667fad1df88985132c1ce1b133428")
	py, _ = hex.DecodeString("ba1d18df944a527ec4ebad9c84cc32643064dcd26bf003a9763dad575104e1a3c9fbb02f971169c2736ed5d8812ad8eeedcfa8226977ddb4")
	pz, _ = hex.DecodeString("2d35e8b251eb6b421291cf3a466597759059e01b7cc89f332f96f801ced244299f4da20b9fcedbaa66c5fd3508dcb61888e2b89bee4fea45")
	pt, _ = hex.DecodeString("8713cc3806a247771ae8567b3b73dd874a8261a610de7c34202fab877f15213120e2fd14e5b191663c1e62d404c54b9f63e1e2e3d98eafb2")
	pu, _ = hex.DecodeString("eafb1cd470e2728ee254c7a312092e820656c14a993f2896479aa211b0a1bb515deee36d06acee20a40a1cad5dc5cc38072cdd63447587e9")
	expectedP := &twExtensible{
		new(bigNumber).setBytes(px),
		new(bigNumber).setBytes(py),
		new(bigNumber).setBytes(pz),
		new(bigNumber).setBytes(pt),
		new(bigNumber).setBytes(pu),
	}

	linear_combo_var_fixed_vt(p, x[:], y[:], wnfsTable[:])

	c.Assert(p.equals(expectedP), Equals, true)
}

func (s *Ed448Suite) TestRecodeWnafForS0(c *C) {
	//nbits_var := 446
	nbits_pre := 446
	table_bits_pre := 5
	//struct smvt_control control_var[nbits_var/(table_bits_var+1)+3];
	controlLen := nbits_pre/(table_bits_pre+1) + 3
	control_pre := make([]smvt_control, controlLen)
	sig := [scalarWords]word_t{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	position := recodeWnaf(control_pre[:], sig[:], nbits_pre, table_bits_pre)

	c.Assert(position, Equals, uint32(0))
	c.Assert(control_pre[position].power, Equals, -1)
	c.Assert(control_pre[position].addend, Equals, 0)
}

func (s *Ed448Suite) TestRecodeWnafForChallenge(c *C) {
	nbits := 446
	table_bits := 4
	controlLen := nbits/(table_bits+1) + 3
	control := make([]smvt_control, controlLen)
	challenge := [scalarWords]word_t{
		0xfd27ffdd, 0xa4a42c92,
		0xd9464f36, 0xac8078dd,
		0x91e922f8, 0x76ebe5e8,
		0x4f1d8f84, 0x968d2c41,
		0x857c5a17, 0x9f74691c,
		0x3595bd83, 0x5b966fb6,
		0xb1428aca, 0x31b43b4d,
	}

	position := recodeWnaf(control[:], challenge[:], nbits, table_bits)

	c.Assert(position, Equals, uint32(67))
	c.Assert(control[position].power, Equals, -1)
	c.Assert(control[position].addend, Equals, 0)
}