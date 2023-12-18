package prompts

type Input struct {
	icon         string
	placeholder  string
	value        string
	defaultValue string
}

func (i Input) Value() string {
	if i.value == "" {
		return i.defaultValue
	}
	return i.value
}

func NewInput(icon, placeholder, defaultValue string) Input {
	return Input{
		icon:         icon,
		placeholder:  placeholder,
		value:        "",
		defaultValue: defaultValue,
	}
}
