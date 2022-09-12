# LCMS plugin for Steampipe

This read the latest release of [Pix Editor](https://github.com/1024pix/pix-editor) and expose it as tables.

## Usage

    steampipe plugin install ghcr.io/1024pix/learningcontent

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/learningcontent.spc ~/.steampipe/config/learningcontent.spc

## License

Apache 2

[steampipe]: https://steampipe.io

