log: centralized & structured logging package
=============================================

Example
-------

Note: *it is highly recommended to set the release tag to a real Git release
tag across repositoires!* (with `log.SetRelease` method)

```go
	if env == "production" {
		log.DialLogstash("tcp", "192.168.199.11:9876")
		log.SetRelease("1.2.3")
	}

	log.SetLevel(log.DebugLevel)

	log.Fields{
		"key1": "value1",
		"key2":   12,
	}.Info("info message")

	log.Fields{
		"key1": "value1",
		"key2":   12,
	}.Infof("info message %d", 1234)

	log.Debug("debug message")

	log.Errorf("an error: %d", 1234)
```

Default Log Items
-----------------

There are two modes: Development & Production. Some information are omitted in
Development mode for clarity.

* @timestamp: time when the entry is logged
* level: panic, fatal, error, warn, info, debug
* pos
    - pkg: package path
    - file: source code file
    - func: function name
    - line: line number
* process: (Production mode only) path of the current executable file
* release: (Production mode only) release version tag.
* message: a text log message

Customized structured object can be logged as long as the object can be
marshaled as JSON.

If there is a `message` field in the log.Fields map, it will be marshaled as
`fields.message` in Logstash JSON format.

Log Daemon
----------
Log Daemon is used to change the log level on the fly. It is automatically
started when the program starts, and listens to TCP address 127.0.0.1 on a 
random port. The port information is logged into an "info" level log entry:

* level: info
* service: Log Daemon
* msg: Log Daemon is up and listening
* port: xxxxx

You can use telnet to connect to the Log Daemon and use command
`level <level name>` to change the log level dynamically:

```
telnet 127.0.0.1 <port>

scape character is '^]'.

Log Daemon Usage:

quit
    quit the session.
show level
    print the current level.
level (panic|fatal|error|warn|info|debug)
    change the current level.
help
    print the help message.
```

log-forwarder
-------------

```yml
input:
    type: redis
    redis:
        addr:     127.0.0.1:6379
        listkey:  log-forwarder:utrack-log
output:
    type: redis
    redis:
        addr:     sh.appcoachs.com:6389
        password: N^9j430i5poiasdff72GjSW90^fsdf13adMU3sodfq10764snbh^f32fdsYiDFAjd32sasfasdfdsfasfdadfafNJs34Us*Y4FdDE5Wlnb94301GWZasadfsdfDAP35145310^T
        listkey:  logstash-utrack
```
