GoTV
====

GoTV is an open source app for generating m3u playlists from different sources.
Currently only sources from http://only-tv.org are supported.

CONFIGURATION
-------------

See `config.yml` for an example:

```yaml
channels:
    mezzo:
        plugin: onlytv
        name: Mezzo
        page_url: http://only-tv.org/mezzo.html
        logo_url: http://only-tv.org/images/mezzo.png
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