GoTV
====

GoTV is an open source app for generating m3u playlists from different sources.
Currently only sources from https://telik.live/ are supported.

CONFIGURATION
-------------

See `config.yml` for an example:

```yaml
channels:
    mezzo:
      plugin: teliklive
      title: Mezzo Live HD
      page_url: https://telik.live/mezzo-live-hd.html
      logo_url: https://telik.live/images/mezzo-live-hd.png
```

USAGE
-----

```shell
$ gotv
Usage of gotv:
  -config string
        config file to read configuration from (default "config.yml")
  -m3u string
        m3u file to save a new playlist into (default "gotv.m3u")
```
