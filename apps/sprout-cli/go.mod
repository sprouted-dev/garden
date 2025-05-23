module sprouted.dev/sprout-cli

go 1.21

require sprouted.dev/weather v0.0.0

require github.com/google/uuid v1.6.0 // indirect

replace sprouted.dev/weather => ../../libs/weather
