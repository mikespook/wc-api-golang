# WooCommerce API - Golang Client

A Golang wrapper for the WooCommerce REST API. Easily interact with the WooCommerce REST API using this library.

[![build status](https://secure.travis-ci.org/mikespook/wc-api-golang.svg)](http://travis-ci.org/mikespook/wc-api-golang)

## Installation

```bash
$ go get github.com/mikespook/wc-api-golang
```

## Getting started

Generate API credentials (Consumer Key & Consumer Secret) following this instructions <http://docs.woocommerce.com/document/woocommerce-rest-api/>
.

Check out the WooCommerce API endpoints and data that can be manipulated in <https://woocommerce.github.io/woocommerce-rest-api-docs/>.

## Setup

Setup for the new WP REST API integration (WooCommerce 2.6 or later):

```golang
import (
  "github.com/mikespook/wc-api-golang"  
)

var woocommerce = wc.NewClient(
    "http://example.com", 
    "ck_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", 
    "cs_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    wc.Options {
        "wp_api": true,
        "version": "wc/v1",
    }
);
```

### Options

|       Option      |   Type   |                Description                 |
| ----------------- | -------- | ------------------------------------------ |
| `url`             | `string` | Your Store URL, example: http://woo.dev/   |
| `consumer_key`    | `string` | Your API consumer key                      |
| `consumer_secret` | `string` | Your API consumer secret                   |
| `options`         | `wc.Options`  | Extra arguments (see client options table) |

#### Client options

|        Option       |   Type   | Required |                                                      Description                                                       |
|---------------------|----------|----------|------------------------------------------------------------------------------------------------------------------------|
| `wp_api`            | `bool`   | no       | Allow make requests to the new WP REST API integration (WooCommerce 2.6 or later)                                      |
| `wp_api_prefix`     | `string` | no       | Custom WP REST API URL prefix, used to support custom prefixes created with the `rest_url_prefix` filter               |
| `version`           | `string` | no       | API version, default is `v3`                                                                                           |
| `timeout`           | `int`    | no       | Request timeout, default is `15`                                                                                       |
| `verify_ssl`        | `bool`   | no       | Verify SSL when connect, use this option as `false` when need to test with self-signed certificates, default is `true` |
| `query_string_auth` | `bool`   | no       | Force Basic Authentication as query string when `true` and using under HTTPS, default is `false`                       |
| `oauth_timestamp`   | `string` | no       | Custom oAuth timestamp, default is `time()`                                                                            |

## Methods

|    Params    |   Type   |                         Description                          |
| ------------ | -------- | ------------------------------------------------------------ |
| `endpoint`   | `string` | WooCommerce API endpoint, example: `customers` or `order/12` |
| `data`       | `interface{}`  | Only for POST and PUT, data that will be converted to JSON   |
| `parameters` | `url.Values`  | Only for GET and DELETE, request query string                |

### GET

```golang
woocommerce.Get(endpoint, parameters)
```

### POST

```golang
woocommerce.Post(endpoint, data)
```

### PUT

```golang
woocommerce.Put(endpoint, data)
```

### DELETE

```golang
woocommerce.Delete(endpoint, parameters)
```

### OPTIONS

```golang
woocommerce.Options(endpoint)
```

#### Response

All methods will return nil on success or an error on failure.

## Release History

