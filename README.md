# `heartbeep` - minimum viable heartbeat

Minimalistic Go tool to periodically check the availability of a remote machine over TCP.

Intended use is for remote machines on a local network, and with `systemd`'s timers to do this loop on repeat.

## Usage

Run `heartbeep` using the `-ip` and `-port` flags to specify the target IP and port:

```bash
./heartbeep -ip=8.8.8.8 -port=53    # returns silently, because it succeeds.
./heartbeep -ip=10.0.0.2 -port=22
```

To integrate `heartbeep` with `systemd` for consistent periodic checks (this example uses `10.0.0.2` on port `22`):

1. Create a `systemd` service file `/etc/systemd/system/heartbeep-monitor.service`:

```ini
[Unit]
Description=Heartbeep Monitor for 10.0.0.2

[Service]
Type=oneshot
ExecStart=/path/to/heartbeep -ip=10.0.0.2 -port=22
```

2. Create a timer file `/etc/systemd/system/heartbeep-monitor.timer`:

```ini
[Unit]
Description=Run Heartbeep monitoring every second

[Timer]
OnBootSec=10
OnUnitActiveSec=1
Unit=heartbeep-monitor.service

[Install]
WantedBy=timers.target
```

3. Activate and start the timer:

```bash
sudo systemctl daemon-reload
sudo systemctl enable heartbeep-monitor.timer
sudo systemctl start heartbeep-monitor.timer
```

## How to build it

_Requires Go._

1. Clone the repository:

```bash
git clone https://github.com/hiAndrewQuinn/heartbeep.git
cd heartbeep
```

2. Build the binary.

   - To build the `heartbeep` binary normally:

     ```bash
     go build -o heartbeep heartbeep.go
     ```

   - If you want a fully static binary, aka something you can `sftp` up to a Unix server and just _go_:

     ```bash
     CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o heartbeep heartbeep.go
     ```

     This ensures that the binary doesn't depend on any shared libraries on the system, making it highly portable.

## License

This project is released under the [Creative Commons Zero v1.0 Universal](https://creativecommons.org/publicdomain/zero/1.0/), allowing for maximum flexibility for end users.
