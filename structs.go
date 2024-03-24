package main

type MapData struct {
	Data struct {
		ZCenter           int      `nbt:"zCenter"`
		UnlimitedTracking byte     `nbt:"unlimitedTracking"`
		TrackingPosition  byte     `nbt:"trackingPosition"`
		Frames            []int    `nbt:"frames"`
		Scale             byte     `nbt:"scale"`
		Locked            byte     `nbt:"locked"`
		Dimension         string   `nbt:"dimension"`
		Banners           []string `nbt:"banners"`
		XCenter           int      `nbt:"xCenter"`
		Colors            []byte   `nbt:"colors"`
	}
	DataVersion int `nbt:"DataVersion"`
}
