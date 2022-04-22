package rescueutil

// Recover from panic
// Usage:
//	func main() {
//		defer rescueutil.Recover(func(err any) {
//	        fmt.Printf("Recover error: %v \n", err)
//		},func() {
//			...
//		})
// 		...
// }
func Recover(errHandler func(err any), cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if r := recover(); r != nil {
		errHandler(r)
	}
}
