package sendbookingconfirmation

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/domain"
	portsendbookingconfirmation "github.com/deday-pool-villa/backend/internal/port/linenotify/sendbookingconfirmation"
)

const dateLayout = "2006-01-02"

var thaiMonths = [...]string{
	"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
	"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
}

type service struct {
	cfg       config.Config
	messenger portsendbookingconfirmation.Messenger
}

// New builds the booking-confirmation notification service.
func New(cfg config.Config, messenger portsendbookingconfirmation.Messenger) Service {
	return &service{cfg: cfg, messenger: messenger}
}

// Execute mirrors sendBookingConfirmation() in line-notify.ts. Failures are
// returned to the caller to log, never surfaced to the customer. Callers run it
// synchronously before responding, since Cloud Run throttles CPU afterwards.
func (s *service) Execute(booking domain.Booking) error {
	if s.cfg.LineChannelToken == "" {
		return nil
	}

	checkinT, err1 := time.Parse(dateLayout, booking.CheckIn)
	checkoutT, err2 := time.Parse(dateLayout, booking.CheckOut)
	nights := 0
	if err1 == nil && err2 == nil {
		nights = int(checkoutT.Sub(checkinT).Hours() / 24)
	}

	message := strings.Join([]string{
		"✅ การจองสำเร็จแล้ว!",
		"",
		fmt.Sprintf("📋 หมายเลขการจอง: %s", booking.ID),
		fmt.Sprintf("👤 ชื่อ: %s", booking.DisplayName),
		fmt.Sprintf("📅 เช็คอิน: %s", thaiLongDate(booking.CheckIn)),
		fmt.Sprintf("📅 เช็คเอาท์: %s", thaiLongDate(booking.CheckOut)),
		fmt.Sprintf("🌙 จำนวนคืน: %d คืน", nights),
		fmt.Sprintf("👥 จำนวนผู้เข้าพัก: %d คน", booking.Guests),
		fmt.Sprintf("💰 ราคารวม: %s บาท", thousandsSep(booking.TotalPrice)),
		"",
		"ทีมงานจะติดต่อกลับเพื่อยืนยันการจองเร็วๆ นี้ 🙏",
	}, "\n")

	return s.messenger.Push(booking.LineUserID, message)
}

// thaiLongDate mirrors `new Date(x).toLocaleDateString('th-TH', { dateStyle: 'long' })`,
// which renders "<day> <month> <Buddhist-Era year>", e.g. "1 สิงหาคม 2569".
func thaiLongDate(dateStr string) string {
	t, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return dateStr
	}
	beYear := t.Year() + 543
	return fmt.Sprintf("%d %s %d", t.Day(), thaiMonths[int(t.Month())-1], beYear)
}

// thousandsSep mirrors `n.toLocaleString()` for a plain number: groups the
// integer part with commas, keeps a trailing fractional part if non-zero.
func thousandsSep(n float64) string {
	whole := int64(n)
	frac := n - float64(whole)

	s := strconv.FormatInt(whole, 10)
	neg := strings.HasPrefix(s, "-")
	if neg {
		s = s[1:]
	}
	var grouped strings.Builder
	for i, c := range s {
		if i != 0 && (len(s)-i)%3 == 0 {
			grouped.WriteByte(',')
		}
		grouped.WriteRune(c)
	}
	out := grouped.String()
	if neg {
		out = "-" + out
	}
	if frac != 0 {
		out += strings.TrimPrefix(strconv.FormatFloat(frac, 'f', 2, 64), "0")
	}
	return out
}
