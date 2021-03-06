package setupcode

import (
	"strings"

	"strconv"

	"github.com/sirupsen/logrus"
	qrcode "github.com/skip2/go-qrcode"
)

// GenXhmURI generates the xhm uri for that includes the pin and the setupid
// References are:
// - https://github.com/nfarina/homebridge/issues/1588#issuecomment-341158722
// - https://github.com/KhaosT/HAP-NodeJS/blob/67032e75b9f5f74993ad932c849d5bbb1937a097/lib/Accessory.js#L363
func GenXhmURI(category int64, hapType int64, pin string, setupID string) string {

	prefix := "X-HM://00"
	var payload int64

	payload = 0
	cat := category << 31
	payload |= cat
	logrus.Debugf("payload with cat: %d", payload)

	ip := 1 << 28
	payload |= int64(ip)
	logrus.Debugf("payload with ip: %d", payload)

	u, err := strconv.ParseUint(pin, 10, 64)
	if err != nil {
		logrus.Error(err)
	}

	payload |= int64(u)
	logrus.Debugf("payload with pin: %d", payload)

	// covert to base 36
	s36 := strings.ToUpper(strconv.FormatInt(payload, 36))
	logrus.Debugf("payload base36: %s", s36)

	content := prefix + s36 + setupID
	return content
}

// GenCliQRCode takes a xhm uri string and returns the qr code as string
func GenCliQRCode(xhm string) string {
	logrus.Debugf("X-HM URL: %s", xhm)
	var q *qrcode.QRCode
	q, err := qrcode.New(xhm, qrcode.Highest)
	if err != nil {
		logrus.Fatal(err)
		return ""
	}

	art := q.ToString(false)
	return art
}
