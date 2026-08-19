package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/handler"
	"github.com/deday-pool-villa/backend/internal/middleware"

	bookingcreate "github.com/deday-pool-villa/backend/internal/service/booking/create"
	bookinglistforuser "github.com/deday-pool-villa/backend/internal/service/booking/listforuser"

	portbookingcreate "github.com/deday-pool-villa/backend/internal/port/booking/create"
	portbookinglistforuser "github.com/deday-pool-villa/backend/internal/port/booking/listforuser"

	availabilitycheck "github.com/deday-pool-villa/backend/internal/service/availability/check"
	availabilitygetunavailabledates "github.com/deday-pool-villa/backend/internal/service/availability/getunavailabledates"

	portavailabilitycheck "github.com/deday-pool-villa/backend/internal/port/availability/check"
	portavailabilitygetunavailabledates "github.com/deday-pool-villa/backend/internal/port/availability/getunavailabledates"

	portpricingcalculate "github.com/deday-pool-villa/backend/internal/port/pricing/calculate"
	pricingcalculate "github.com/deday-pool-villa/backend/internal/service/pricing/calculate"

	portadminauthlogin "github.com/deday-pool-villa/backend/internal/port/adminauth/login"
	adminauthlogin "github.com/deday-pool-villa/backend/internal/service/adminauth/login"
	adminauthparsetoken "github.com/deday-pool-villa/backend/internal/service/adminauth/parsetoken"

	portlineauthverifyidtoken "github.com/deday-pool-villa/backend/internal/port/lineauth/verifyidtoken"
	lineauthverifyidtoken "github.com/deday-pool-villa/backend/internal/service/lineauth/verifyidtoken"

	portlinenotifysendbookingconfirmation "github.com/deday-pool-villa/backend/internal/port/linenotify/sendbookingconfirmation"
	linenotifysendbookingconfirmation "github.com/deday-pool-villa/backend/internal/service/linenotify/sendbookingconfirmation"

	portadminblockeddatescreate "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/create"
	portadminblockeddatesdelete "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/delete"
	portadminblockeddateslist "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/list"
	adminblockeddatescreate "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/create"
	adminblockeddatesdelete "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/delete"
	adminblockeddateslist "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/list"

	portadmincustomperiodscreate "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/create"
	portadmincustomperiodsdelete "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/delete"
	portadmincustomperiodslist "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/list"
	portadmincustomperiodsupdate "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/update"
	admincustomperiodscreate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/create"
	admincustomperiodsdelete "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/delete"
	admincustomperiodslist "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/list"
	admincustomperiodsupdate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/update"

	portadmindayrateslist "github.com/deday-pool-villa/backend/internal/port/admin/dayrates/list"
	portadmindayratesupsert "github.com/deday-pool-villa/backend/internal/port/admin/dayrates/upsert"
	admindayrateslist "github.com/deday-pool-villa/backend/internal/service/admin/dayrates/list"
	admindayratesupsert "github.com/deday-pool-villa/backend/internal/service/admin/dayrates/upsert"

	portadminbookingslist "github.com/deday-pool-villa/backend/internal/port/admin/bookings/list"
	portadminbookingsupdatestatus "github.com/deday-pool-villa/backend/internal/port/admin/bookings/updatestatus"
	adminbookingslist "github.com/deday-pool-villa/backend/internal/service/admin/bookings/list"
	adminbookingsupdatestatus "github.com/deday-pool-villa/backend/internal/service/admin/bookings/updatestatus"
)

func New(db *gorm.DB, cfg config.Config) *gin.Engine {
	h := buildHandler(db, cfg)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS(cfg))

	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	api := r.Group("/api")
	{
		api.GET("/pricing-data", h.GetPricingData)
		api.GET("/availability", h.GetAvailability)
		api.POST("/webhook/line", h.LineWebhook)

		customer := api.Group("")
		customer.Use(middleware.LineAuth(cfg, h.LineAuthVerifyIDToken))
		{
			customer.POST("/bookings", h.CreateBooking)
			customer.GET("/user/bookings", h.GetUserBookings)
		}

		api.POST("/admin/login", h.AdminLogin)

		admin := api.Group("/admin")
		admin.Use(middleware.AdminAuth(h.AdminAuthParseToken))
		{
			admin.GET("/bookings", h.AdminListBookings)
			admin.PATCH("/bookings/:bookingId/status", h.AdminUpdateBookingStatus)

			admin.GET("/blocked-dates", h.AdminListBlockedDates)
			admin.POST("/blocked-dates", h.AdminCreateBlockedDate)
			admin.DELETE("/blocked-dates/:id", h.AdminDeleteBlockedDate)

			admin.GET("/day-rates", h.AdminListDayRates)
			admin.PUT("/day-rates", h.AdminUpdateDayRates)

			admin.GET("/custom-periods", h.AdminListCustomPeriods)
			admin.POST("/custom-periods", h.AdminCreateCustomPeriod)
			admin.PUT("/custom-periods/:id", h.AdminUpdateCustomPeriod)
			admin.DELETE("/custom-periods/:id", h.AdminDeleteCustomPeriod)
		}
	}

	return r
}

// buildHandler wires every port adaptor into its service, and every service
// into the Handler — the composition root for the whole hexagon.
func buildHandler(db *gorm.DB, cfg config.Config) *handler.Handler {
	availabilityCheckSvc := availabilitycheck.New(portavailabilitycheck.NewPostgresRepository(db))
	availabilityGetUnavailableSvc := availabilitygetunavailabledates.New(portavailabilitygetunavailabledates.NewPostgresRepository(db))
	pricingCalculateSvc := pricingcalculate.New(portpricingcalculate.NewPostgresRepository(db))

	bookingCreateSvc := bookingcreate.New(portbookingcreate.NewPostgresRepository(db), availabilityCheckSvc, pricingCalculateSvc)
	bookingListForUserSvc := bookinglistforuser.New(portbookinglistforuser.NewPostgresRepository(db))

	adminAuthLoginSvc := adminauthlogin.New(portadminauthlogin.NewPostgresRepository(db), cfg)
	adminAuthParseTokenSvc := adminauthparsetoken.New(cfg)

	lineAuthVerifySvc := lineauthverifyidtoken.New(cfg, portlineauthverifyidtoken.NewLineAPIVerifier())
	lineNotifySvc := linenotifysendbookingconfirmation.New(cfg, portlinenotifysendbookingconfirmation.NewLineAPIMessenger(cfg.LineChannelToken))

	blockedDatesListSvc := adminblockeddateslist.New(portadminblockeddateslist.NewPostgresRepository(db))
	blockedDatesCreateSvc := adminblockeddatescreate.New(portadminblockeddatescreate.NewPostgresRepository(db))
	blockedDatesDeleteSvc := adminblockeddatesdelete.New(portadminblockeddatesdelete.NewPostgresRepository(db))

	customPeriodsListSvc := admincustomperiodslist.New(portadmincustomperiodslist.NewPostgresRepository(db))
	customPeriodsCreateSvc := admincustomperiodscreate.New(portadmincustomperiodscreate.NewPostgresRepository(db))
	customPeriodsUpdateSvc := admincustomperiodsupdate.New(portadmincustomperiodsupdate.NewPostgresRepository(db))
	customPeriodsDeleteSvc := admincustomperiodsdelete.New(portadmincustomperiodsdelete.NewPostgresRepository(db))

	dayRatesListSvc := admindayrateslist.New(portadmindayrateslist.NewPostgresRepository(db))
	dayRatesUpsertSvc := admindayratesupsert.New(portadmindayratesupsert.NewPostgresRepository(db))

	bookingsListSvc := adminbookingslist.New(portadminbookingslist.NewPostgresRepository(db))
	bookingsUpdateStatusSvc := adminbookingsupdatestatus.New(portadminbookingsupdatestatus.NewPostgresRepository(db))

	return &handler.Handler{
		Cfg: cfg,

		BookingCreate:      bookingCreateSvc,
		BookingListForUser: bookingListForUserSvc,

		AvailabilityCheck:               availabilityCheckSvc,
		AvailabilityGetUnavailableDates: availabilityGetUnavailableSvc,

		PricingCalculate: pricingCalculateSvc,

		AdminAuthLogin:      adminAuthLoginSvc,
		AdminAuthParseToken: adminAuthParseTokenSvc,

		LineAuthVerifyIDToken: lineAuthVerifySvc,

		LineNotifySendBookingConfirmation: lineNotifySvc,

		AdminBlockedDatesList:   blockedDatesListSvc,
		AdminBlockedDatesCreate: blockedDatesCreateSvc,
		AdminBlockedDatesDelete: blockedDatesDeleteSvc,

		AdminCustomPeriodsList:   customPeriodsListSvc,
		AdminCustomPeriodsCreate: customPeriodsCreateSvc,
		AdminCustomPeriodsUpdate: customPeriodsUpdateSvc,
		AdminCustomPeriodsDelete: customPeriodsDeleteSvc,

		AdminDayRatesList:   dayRatesListSvc,
		AdminDayRatesUpsert: dayRatesUpsertSvc,

		AdminBookingsList:         bookingsListSvc,
		AdminBookingsUpdateStatus: bookingsUpdateStatusSvc,
	}
}
