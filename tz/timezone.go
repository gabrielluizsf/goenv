package tz

type TZ interface {
	Name() string
}

var Timezone = timezone{}
