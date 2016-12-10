# Hu[e controller in] go[lang]

# Usage

    $ go run cmd/hugoset/main.go --help
    Usage of hugoctl:
      -h string
           	Home description json file (default "home.json")
      -d   	Disco mode
      -s string
           	Scene description json file (default "scene.json")


# Home description

`home.json` is expected to contain something like:

```json
{
  "bridge": {
    "ip": "192.168.1.412",
    "api-key": "blahblahblah"
  },
  "places": {
    "kitchen": {
      "lights": [
        "Kitchen",
        "Sink"
      ]
    },
    "salon": {
      "lights": [
        "Counter",
        "Salon 1",
        "Salon 2",
        "Salon 3"
      ]
    },
    "desk": {
      "lights": [
        "Lightstrip salon",
        "Bureau",
        "Desk bloom",
        "Spot salon"
      ]
    },
    "under-desk": {
      "lights": [
        "Desk lightstrip"
      ]
    },
    "room": {
      "lights": [
        "Sapin 1",
        "Sapin 2",
        "Sapin 3",
        "Sapin 4",
        "Iris"
      ]
    }
  }
}
```


# Scene description

`scene.json` is expected to contain something like:

```json
[
  {
    "place": "salon",
    "brightness": 100,
    "saturation": 100,
    "hue": 120
  },
  {
    "place": "desk",
    "brightness": 100,
    "saturation": 100,
    "hue": 360
  },
  {
    "place": "under-desk",
    "brightness": 100,
    "saturation": 100,
    "hue": 240
  },
  {
    "place": "kitchen",
    "brightness": 100,
    "saturation": 100,
    "hue": 224
  },
  {
    "place": "room",
    "brightness": 20,
    "saturation": 100,
    "hue": 220
  }
]
```
