module github.com/kamolhasan/monorepo/multi-module/service-2

go 1.16

require github.com/kamolhasan/monorepo/multi-module/service-1 v0.0.0

replace github.com/kamolhasan/monorepo/multi-module/service-1 => ../service-1
