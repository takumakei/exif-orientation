package orientation

// Orientation
//
//            1:TopLeft                   2:TopRight
//             FFFFFFFF                    FFFFFFFF
//             FF                                FF
//             FFFFFF                        FFFFFF
//             FF                                FF
//             FF                                FF
//             FF                                FF
//
//                 8:LeftBottom    5:LeftTop
//     FF          FFFFFFFFFFFF    FFFFFFFFFFFF          FF
//     FF  FF            FF  FF    FF  FF            FF  FF
//     FF  FF            FF  FF    FF  FF            FF  FF
//     FFFFFFFFFFFF          FF    FF          FFFFFFFFFFFF
//     6:RightTop                             7:RightBottom
//
//                   FF                    FF
//                   FF                    FF
//                   FF                    FF
//               FFFFFF                    FFFFFF
//                   FF                    FF
//             FFFFFFFF                    FFFFFFFF
//           3:BottomRight               4:BottomLeft
//
type Orientation int

const (
	TopLeft Orientation = 1 + iota
	TopRight
	BottomRight
	BottomLeft
	LeftTop
	RightTop
	RightBottom
	LeftBottom
)

//go:generate stringer -type Orientation
