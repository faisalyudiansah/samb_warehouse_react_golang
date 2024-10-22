package servers

import (
	"server/helpers/validationutils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(validationutils.TagNameFormatter)
		v.RegisterCustomTypeFunc(validationutils.DecimalType, decimal.Decimal{})

		v.RegisterValidation("dgt", validationutils.DecimalGT)
		v.RegisterValidation("dlt", validationutils.DecimalLT)
		v.RegisterValidation("dgte", validationutils.DecimalGTE)
		v.RegisterValidation("dlte", validationutils.DecimalLTE)
		v.RegisterValidation("date", validationutils.ValidateDate)
		v.RegisterValidation("role", validationutils.RoleValidator)
		v.RegisterValidation("numeric", validationutils.NumericValidator)
		v.RegisterValidation("password", validationutils.PasswordValidator)
		v.RegisterValidation("clean_input", validationutils.CleanInputValidator)
		v.RegisterValidation("phone_number", validationutils.PhoneNumberValidator)
		v.RegisterValidation("time_format", validationutils.TimeFormatValidator)
		v.RegisterValidation("day_of_weeks", validationutils.DayOfWeekValidator)
		v.RegisterValidation("no_duplicates", validationutils.NoDuplicatesValidator)
	}
}
