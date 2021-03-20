package logger

import (
	"fmt"

	"github.com/cihub/seelog"
)

const (
	logConfig = `
	<seelog type="asynctimer" asyncinterval="1000000" minlevel="trace" maxlevel="critical">
    <outputs formatid="main">
        <console/>
		<buffered formatid="main" size="10000" flushperiod="1000">      
		    <rollingfile type="date" filename="./logs/log" datepattern="20060102.txt" fullname="true" maxrolls="5" />      
		</buffered>
    </outputs>
    <formats>
        <!-- 设置格式 -->
		<format id="func" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n"/>
		<format id="main" format="%Date %Time [%LEV] [%File:%Line] %Msg%n"/>
    </formats>
</seelog>
	`
)

var Tracef = tracef()

func tracef() func(format string, params ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println(err)
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Tracef
}

var Debugf = debugf()

func debugf() func(format string, params ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Debugf
}

var Infof = infof()

func infof() func(format string, params ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Infof
}

var Warnf = warnf()

func warnf() func(format string, params ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Warnf
}

var Errorf = errorf()

func errorf() func(format string, params ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Errorf
}

var Criticalf = criticalf()

func criticalf() func(format string, params ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Criticalf
}

var Trace = trace()

func trace() func(v ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Trace
}

var Debug = debug()

func debug() func(v ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Debug
}

var Info = info()

func info() func(v ...interface{}) {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Info
}

var Warn = warn()

func warn() func(v ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Warn
}

var Error = errora()

func errora() func(v ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Error
}

var Critical = critical()

func critical() func(v ...interface{}) error {
	defer seelog.Flush() //加载配置文件
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		fmt.Println("parse config error")
	}
	//替换记录器
	seelog.ReplaceLogger(logger)
	return seelog.Critical
}
