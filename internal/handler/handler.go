package handler

import (
	bookingcreate "github.com/deday-pool-villa/backend/internal/service/booking/create"
	bookinglistforuser "github.com/deday-pool-villa/backend/internal/service/booking/listforuser"

	availabilitycheck "github.com/deday-pool-villa/backend/internal/service/availability/check"
	availabilitygetunavailabledates "github.com/deday-pool-villa/backend/internal/service/availability/getunavailabledates"

	pricingcalculate "github.com/deday-pool-villa/backend/internal/service/pricing/calculate"

	adminauthlogin "github.com/deday-pool-villa/backend/internal/service/adminauth/login"
	adminauthparsetoken "github.com/deday-pool-villa/backend/internal/service/adminauth/parsetoken"

	lineauthverifyidtoken "github.com/deday-pool-villa/backend/internal/service/lineauth/verifyidtoken"

	linenotifysendbookingconfirmation "github.com/deday-pool-villa/backend/internal/service/linenotify/sendbookingconfirmation"

	adminblockeddatescreate "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/create"
	adminblockeddatesdelete "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/delete"
	adminblockeddateslist "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/list"

	admincustomperiodscreate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/create"
	admincustomperiodsdelete "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/delete"
	admincustomperiodslist "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/list"
	admincustomperiodsupdate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/update"

	admindayrateslist "github.com/deday-pool-villa/backend/internal/service/admin/dayrates/list"
	admindayratesupsert "github.com/deday-pool-villa/backend/internal/service/admin/dayrates/upsert"

	adminbookingslist "github.com/deday-pool-villa/backend/internal/service/admin/bookings/list"
	adminbookingsupdatestatus "github.com/deday-pool-villa/backend/internal/service/admin/bookings/updatestatus"

	"github.com/deday-pool-villa/backend/internal/config"
)

// Handler holds every application service the HTTP layer (the driving
// adaptor) calls into. It has no business logic of its own — only request
// parsing, response shaping, and delegation to a service.
type Handler struct {
	Cfg config.Config

	BookingCreate      bookingcreate.Service
	BookingListForUser bookinglistforuser.Service

	AvailabilityCheck               availabilitycheck.Service
	AvailabilityGetUnavailableDates availabilitygetunavailabledates.Service

	PricingCalculate pricingcalculate.Service

	AdminAuthLogin      adminauthlogin.Service
	AdminAuthParseToken adminauthparsetoken.Service

	LineAuthVerifyIDToken lineauthverifyidtoken.Service

	LineNotifySendBookingConfirmation linenotifysendbookingconfirmation.Service

	AdminBlockedDatesList   adminblockeddateslist.Service
	AdminBlockedDatesCreate adminblockeddatescreate.Service
	AdminBlockedDatesDelete adminblockeddatesdelete.Service

	AdminCustomPeriodsList   admincustomperiodslist.Service
	AdminCustomPeriodsCreate admincustomperiodscreate.Service
	AdminCustomPeriodsUpdate admincustomperiodsupdate.Service
	AdminCustomPeriodsDelete admincustomperiodsdelete.Service

	AdminDayRatesList   admindayrateslist.Service
	AdminDayRatesUpsert admindayratesupsert.Service

	AdminBookingsList         adminbookingslist.Service
	AdminBookingsUpdateStatus adminbookingsupdatestatus.Service
}
