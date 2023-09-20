# `heartbeep` - a really, really simple heartbeat

This is a simple Go-based monitoring tool designed to check the availability of a remote machine over TCP by attempting to establish a connection to a specified port. The primary use case for this tool is in isolated, private networks where a lightweight solution is required.

## Features

- Simple command-line interface.
- Customizable target IP and port.
- Efficient and minimal resource usage.
- Easily integrated with `systemd` for periodic checks.

## Requirements

- Go (tested with Go 1.21)
- (optional) A Unix-like system for `systemd` integration (if used)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/YourUsername/heartbeat-monitor.git
cd heartbeat-monitor
```

2. Build the binary:

```bash
go build -o heartbeat heartbeat.go
```

3. Optionally, move the binary to a standard bin directory:

```bash
sudo mv heartbeat /usr/local/bin/
```

## Usage

Run the heartbeat monitor with the `-ip` and `-port` flags to specify the target IP and port, respectively.

```bash
heartbeat -ip=10.0.0.2 -port=22
```

If the remote machine at the specified IP and port is unreachable, the tool will print a log message indicating that the machine is down.

## Systemd Integration

For periodic checks, you can use `systemd` timers:

1. Create a `systemd` service and timer unit as shown in the provided example files.
2. Reload `systemd` and start the timer:

```bash
sudo systemctl daemon-reload
sudo systemctl enable your-timer-name.timer
sudo systemctl start your-timer-name.timer
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

## License

This project is licensed under the MIT License.

---

Make sure to replace `YourUsername` with your actual GitHub username, and adjust the repository URL accordingly. You can also customize other parts of the `README.md` as you see fit. This should provide a good starting point for your GitHub project!
## Features

- Simple command-line interface.
- Customizable target IP and port.
- Efficient and minimal resource usage.
- Easily integrated with `systemd` for periodic checks.

## Requirements

- Go (tested with Go 1.x)
- A Unix-like system for `systemd` integration (if used)

## Installation

```bash
git clone https://github.com/YourUsername/heartbeat-monitor.git
cd heartbeat-monitor

# Build the binary.
go build -o heartbeat heartbeat.go
```

Of course! Here's a comprehensive `README.md` for your project:

---

# Heartbeat Monitor

This is a simple Go-based monitoring tool designed to check the availability of a remote machine over TCP by attempting to establish a connection to a specified port. The primary use case for this tool is in isolated, private networks where a lightweight solution is required.

## Features

- Simple command-line interface.
- Customizable target IP and port.
- Efficient and minimal resource usage.
- Easily integrated with `systemd` for periodic checks.

## Requirements

- Go (tested with Go 1.x)
- A Unix-like system for `systemd` integration (if used)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/YourUsername/heartbeat-monitor.git
cd heartbeat-monitor
```

2. Build the binary:

```bash
go build -o heartbeat heartbeat.go
```

Or, if you want a _totally_ static binary, which doesn't depend on the
underlying Unix system at all:

```
CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o heartbeat heartbeat.go
```

You can then verify the binary is static by running `file heartbeat`.

3. Optionally, move the binary to a standard bin directory:

```bash
sudo mv heartbeat /usr/local/bin/
```

## Usage

Run the heartbeat monitor with the `-ip` and `-port` flags to specify the target IP and port, respectively.

```bash
heartbeat -ip=10.0.0.2 -port=22
```

If the remote machine at the specified IP and port is unreachable, the tool will print a log message indicating that the machine is down. Otherwise it will silently succeed.

## Systemd Integration

For periodic checks, you can use `systemd` timers:

1. Create a `systemd` service and timer unit. For example, say you wanted to
   monitor `10.0.0.2`, every second, on port `22`:
   1. First you would create a `systemd` service, at `/etc/systemd/system/monitor-10-0-0-2.service`:

      ```toml
      [Unit]
      Description=Monitor 10.0.0.2

      [Service]
      Type=oneshot
      ExecStart=/path/to/heartbeat-binary -ip=10.0.0.2 -port=22
      ```

   2. Then you would create `/etc/systemd/system/monitor-10-0-0-2.timer`:

      ```toml
      [Unit]
      Description=Run monitoring every second

      [Timer]
      OnBootSec=10
      OnUnitActiveSec=1
      Unit=monitor-10-0-0-2.service

      [Install]
      WantedBy=timers.target
      ```

2. Reload `systemd` and start the timer:

```bash
sudo systemctl daemon-reload
sudo systemctl enable monitor-10-0-0-2.timer     # or whatever you named it
sudo systemctl start monitor-10-0-0-2.timer      # or whatever you named it
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

## License

This project is licensed under the MIT License.

---

Make sure to replace `YourUsername` with your actual GitHub username, and adjust the repository URL accordingly. You can also customize other parts of the `README.md` as you see fit. This should provide a good starting point for your GitHub project!
