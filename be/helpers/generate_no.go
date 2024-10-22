package helpers

import (
	"fmt"
	"time"
)

func GenerateTrxInNo(lastTrxInPK int64, desc string) string {
	currentYear := time.Now().Year()
	if lastTrxInPK == 0 {
		return fmt.Sprintf("%v-%03d-%d", desc, 1, currentYear)
	}
	return fmt.Sprintf("%v-%03d-%d", desc, lastTrxInPK+1, currentYear)
}
