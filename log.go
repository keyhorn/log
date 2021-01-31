package log

// var std = NewStdoutLogger()

// // Fatal is equivalent to Print() followed by a call to os.Exit(1).
// func Fatal(v ...interface{}) {
// 	log.Fatal(v...)
// }

// // Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
// func Fatalf(format string, v ...interface{}) {
// 	log.Fatalf(format, v...)
// }

// // Fatalln is equivalent to Println() followed by a call to os.Exit(1).
// func Fatalln(v ...interface{}) {
// 	log.Fatalln(v...)
// }

// // Flags returns the output flags for the standard logger. The flag bits are Ldate, Ltime, and so on.
// func Flags() int {
// 	return log.Flags()
// }

// // Output writes the output for a logging event. The string s contains the text to print after the prefix specified by the flags of the Logger. A newline is appended if the last character of s is not already a newline. Calldepth is the count of the number of frames to skip when computing the file name and line number if Llongfile or Lshortfile is set; a value of 1 will print the details for the caller of Output.
// func Output(calldepth int, s string) error {
// 	return log.Output(calldepth, s)
// }

// // Panic is equivalent to Print() followed by a call to panic().
// func Panic(v ...interface{}) {
// 	log.Panic(v...)
// }

// // Panicf is equivalent to Printf() followed by a call to panic().
// func Panicf(format string, v ...interface{}) {
// 	log.Panicf(format, v...)
// }

// // Panicln is equivalent to Println() followed by a call to panic().
// func Panicln(v ...interface{}) {
// 	log.Panicln(v...)
// }

// // Prefix returns the output prefix for the standard logger.
// func Prefix() string {
// 	return log.Prefix()
// }

// // Print calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
// func Print(v ...interface{}) {
// 	log.Print(v...)
// }

// // Printf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
// func Printf(format string, v ...interface{}) {
// 	log.Printf(format, v...)
// }

// // Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
// func Println(v ...interface{}) {
// 	log.Println(v...)
// }

// // SetFlags sets the output flags for the standard logger. The flag bits are Ldate, Ltime, and so on.
// func SetFlags(flag int) {
// 	log.SetFlags(flag)
// }

// // SetOutput sets the output destination for the standard logger.
// func SetOutput(w io.Writer) {
// 	log.SetOutput(w)
// }

// // SetPrefix sets the output prefix for the standard logger.
// func SetPrefix(prefix string) {
// 	log.SetPrefix(prefix)
// }

// // Writer returns the output destination for the standard logger.
// func Writer() io.Writer {
// 	return log.Writer()
// }
