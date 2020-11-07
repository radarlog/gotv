GoTV
====

GoTV is an open source software for generating m3u playlists from different sources. Currently is supported only sources from http://onelike.tv

CONFIGURATION
-------------

See `meta.yml` for an example:

```yaml
logo_dir: logos
channels:
    mezzo:
        plugin: onelike
        name: Mezzo
        page_url: http://onelike.tv/mezzo.html
        logo_url: http://onelike.tv/images/mezzo.png
```

USAGE
-----

```shell
$ gotv
Usage of gotv:
  -save string
        m3u file to save a new playlist into (default "gotv.m3u")
  -config string
        config file to read configuration from (default "config.yml")
```