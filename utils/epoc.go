package utils

import (
	"strconv"
	"time"
)




func ConvertEpocToHuman(convert_epoch string) string  {
	epoch, err:= strconv.ParseInt(convert_epoch, 10, 64)

	if err != nil{
		return "Error to convert epoch"
	}
	tempo:= time.Unix(epoch, 0)

	return tempo.Format("02-Jan-2006 15:04:05")
}