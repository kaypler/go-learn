// *celsiusFlag 满足 flag.Value 接口
package tempconv

import (
	"fmt"
	"flag"

	"go-learn/ch2/tempconv/Celsius"
	"go-learn/ch2/tempconv/Fahrenheit"
	"go-learn/ch2/tempconv/FToC"
)

type celsiusFLag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFLag(value)
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

