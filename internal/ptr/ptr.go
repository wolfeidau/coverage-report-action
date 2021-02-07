package ptr

func ToInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func ToInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func ToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func Int(v int) *int {
	return &v
}

func String(v string) *string {
	return &v
}
