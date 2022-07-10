# golang-module/carbon
A simple, semantic and developer-friendly golang package for datetime

## 文档地址
https://github.com/golang-module/carbon

## 安装
```
go get -u github.com/golang-module/carbon/v2
import (
	"github.com/golang-module/carbon/v2"
)
```

## 使用
```
// Return datetime of today
fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15

// Return date of today
carbon.Now().ToDateString() // 2020-08-05
// Return time of today
carbon.Now().ToTimeString() // 13:14:15
// Return datetime of today in a given timezone
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// Return timestamp with second of today
carbon.Now().Timestamp() // 1596604455
// Return timestamp with millisecond of today
carbon.Now().TimestampMilli() // 1596604455000
// Return timestamp with microsecond of today
carbon.Now().TimestampMicro() // 1596604455000000
// Return timestamp with nanosecond of today
carbon.Now().TimestampNano() // 1596604455000000000

// Return datetime of yesterday
fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// Return date of yesterday
carbon.Yesterday().ToDateString() // 2020-08-04
// Return time of yesterday
carbon.Yesterday().ToTimeString() // 13:14:15
// Return datetime of yesterday on a given day
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// Return datetime of yesterday in a given timezone
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// Return timestamp with second of yesterday
carbon.Yesterday().Timestamp() // 1596518055
// Return timestamp with millisecond of yesterday
carbon.Yesterday().TimestampMilli() // 1596518055000
// Return timestamp with microsecond of yesterday
carbon.Yesterday().TimestampMicro() // 1596518055000000
// Return timestamp with nanosecond of yesterday
carbon.Yesterday().TimestampNano() // 1596518055000000000

// Return datetime of tomorrow
fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15 +0800 CST
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// Return date of tomorrow
carbon.Tomorrow().ToDateString() // 2020-08-06
// Return time of tomorrow
carbon.Tomorrow().ToTimeString() // 13:14:15
// Return datetime of tomorrow on a given day
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// Return datetime of tomorrow in a given timezone
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// Return timestamp with second of tomorrow
carbon.Tomorrow().Timestamp() // 1596690855
// Return timestamp with millisecond of tomorrow
carbon.Tomorrow().TimestampMilli() // 1596690855000
// Return timestamp with microsecond of tomorrow
carbon.Tomorrow().TimestampMicro() // 1596690855000000
// Return timestamp with nanosecond of tomorrow
carbon.Tomorrow().TimestampNano() // 1596690855000000000
```