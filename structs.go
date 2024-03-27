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

type ColorMap struct {
	Colormap_1_12     map[int][]int `json:"colormap-1_12"`
	Colormap_1_8_1    map[int][]int `json:"colormap-1_8_1"`
	Colormap_1_7_2    map[int][]int `json:"colormap-1_7_2"`
	Colormap_Original map[int][]int `json:"colormap-original"`
}
