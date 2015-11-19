package converter

import "strconv"

func ConvertBytesToString(in interface{}, out *interface{}) error {
	*out = string(in.([]byte))
	return nil
}

func ConvertStringToBytes(in interface{}, out *interface{}) error {
	*out = []byte(in.(string))
	return nil
}

func ConvertIntegerToString(in interface{}, out *interface{}) error {
	*out = strconv.Itoa(in.(int))
	return nil
}

func ConvertStringToInteger(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.Atoi(in.(string))
	return err
}

func ConvertStringToFloat(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.ParseFloat(in.(string), 64)
	return err
}

func ConvertFloatToString(in interface{}, out *interface{}) (err error) {
	*out = strconv.FormatFloat(in.(float64), 'f', -1, 64)
	return err
}
