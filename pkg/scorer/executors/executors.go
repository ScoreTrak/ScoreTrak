package executors

import "github.com/go-playground/validator/v10"

var v = func() *validator.Validate {
	v := validator.New()

	mustRegisterValidation(v, "smb_operation", validateSmbOperation)
	//mustRegisterValidation(v, "prom_expr", validatePromExpression)
	//mustRegisterValidation(v, "prom_label_key", validatePromLabelKey)
	//mustRegisterValidation(v, "prom_label_value", validatePromLabelValue)
	//mustRegisterValidation(v, "prom_annot_key", validatePromAnnotKey)
	//mustRegisterValidation(v, "name", validateName)
	//mustRegisterValidation(v, "required_if_enabled", validateRequiredEnabledAlertName)
	//mustRegisterValidation(v, "template_vars", validateTemplateVars)
	//v.RegisterStructValidation(validateOneSLI, SLI{})
	//v.RegisterStructValidation(validateSLOGroup, SLOGroup{})
	//v.RegisterStructValidation(validateSLIEvents, SLIEvents{})
	return v
}()

// mustRegisterValidation is a helper so we panic on start if we can't register a validator.
func mustRegisterValidation(v *validator.Validate, tag string, fn validator.Func) {
	err := v.RegisterValidation(tag, fn)
	if err != nil {
		panic(err)
	}
}
