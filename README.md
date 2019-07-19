# mosby
> Mosby is an executable that validates a yaml-config.

[![Build Status][travis-image]][travis-url]

_Mosby_ checks if a yaml-file contains one or more of the following-configs:
```yaml
- name: service1
  context:
    - service1
  backend_configuration:
    backends:
      - host: service1.example.com
        port: 443
    connect_timeout: 2
    first_byte_timeout: 2
    between_bytes_timeout: 2
  probe:
    url: "/internal/health"
```

## Installation

```bash
git clone https://github.com/fr3dch3n/mosby
cd mosby
make test

# for ubuntu
make mosby_amd64

# for alpine
make mosby_alpine

# for mac
make mosby_darwin

# for windows
mosby.exe
```

## Usage example\

`./mosby --path "test-resources/valid.config.yaml"`

Run `./mosby help` to see all possibilities.

## Release History

* 0.0.1
    * initial release

## Meta

[@fr3dch3n](https://twitter.com/fr3dch3n)

Distributed under the Apache 2.0 license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/fr3dch3n/mosby/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[travis-image]: https://img.shields.io/travis/fr3dch3n/mosby/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/fr3dch3n/mosby
