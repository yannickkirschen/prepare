# Prepper

If the world comes to an end, at least you've saved your calendar.

## Usage

Before building prepper, you should create the download config in `/etc/prepper/prepper.d/01-download.json`:

```json
{
  "volume": "/path/to/prepper/volume",
  "cron": "0 6 * * *",
  "sources": [
    {
      "file_name": "ical/calendar.ics",
      "url": "https://example.com/calendar.ics"
    }
  ]
}
```

Then you can build and install prepper:

```bash

```bash
go build -ldflags "-linkmode 'external' -extldflags '-static'" -o prepper cmd/main.go
sudo cp prepper /usr/bin/prepper
sudo cp prepper.service /etc/systemd/system/prepper.service
sudo systemctl enable prepper.service
```
