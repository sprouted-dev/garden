# Sprouted.dev Go Module Configuration

This document outlines the configuration needed for sprouted.dev to serve as a custom Go module domain.

## Required Files on sprouted.dev

### 1. /.well-known/go-import

This file tells Go where to find our modules:

```html
<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="sprouted.dev/weather git https://github.com/sprouted-dev/garden">
    <meta name="go-import" content="sprouted.dev/sprout-cli git https://github.com/sprouted-dev/garden">
    <meta name="go-source" content="sprouted.dev/weather https://github.com/sprouted-dev/garden https://github.com/sprouted-dev/garden/tree/main{/dir} https://github.com/sprouted-dev/garden/blob/main{/dir}/{file}#L{line}">  
    <meta name="go-source" content="sprouted.dev/sprout-cli https://github.com/sprouted-dev/garden https://github.com/sprouted-dev/garden/tree/main{/dir} https://github.com/sprouted-dev/garden/blob/main{/dir}/{file}#L{line}">
</head>
<body>
    Go packages for the Sprouted ecosystem
</body>
</html>
```

### 2. URL Redirects

The web server should respond to these paths:

- `GET sprouted.dev/weather?go-get=1` → serve go-import meta tags for weather package  
- `GET sprouted.dev/sprout-cli?go-get=1` → serve go-import meta tags for sprout-cli package

### 3. Subdirectory Handling

Since our packages live in subdirectories of the monorepo:
- `sprouted.dev/weather` maps to `/libs/weather/` in the GitHub repo
- `sprouted.dev/sprout-cli` maps to `/apps/sprout-cli/` in the GitHub repo

## Web Server Configuration

### Nginx Configuration

```nginx
server {
    server_name sprouted.dev;
    
    # Handle Go module requests
    location ~ ^/(weather|sprout-cli)$ {
        if ($args ~ "go-get=1") {
            return 200 '<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="sprouted.dev/$1 git https://github.com/sprouted-dev/garden">
    <meta name="go-source" content="sprouted.dev/$1 https://github.com/sprouted-dev/garden https://github.com/sprouted-dev/garden/tree/main{/dir} https://github.com/sprouted-dev/garden/blob/main{/dir}/{file}#L{line}">
</head>
<body>Go package: sprouted.dev/$1</body>
</html>';
            add_header Content-Type text/html;
        }
        
        # Regular requests go to package documentation
        try_files $uri $uri/ /pkg/$1/;
    }
    
    # Serve the main website
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

### Apache Configuration

```apache
RewriteEngine On

# Handle Go module requests for weather
RewriteCond %{QUERY_STRING} go-get=1
RewriteRule ^/weather$ /go-import.php?package=weather [L]

# Handle Go module requests for sprout-cli  
RewriteCond %{QUERY_STRING} go-get=1
RewriteRule ^/sprout-cli$ /go-import.php?package=sprout-cli [L]
```

With `go-import.php`:

```php
<?php
$package = $_GET['package'];
$valid_packages = ['weather', 'sprout-cli'];

if (!in_array($package, $valid_packages)) {
    http_response_code(404);
    exit;
}

header('Content-Type: text/html');
?>
<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="sprouted.dev/<?= $package ?> git https://github.com/sprouted-dev/garden">
    <meta name="go-source" content="sprouted.dev/<?= $package ?> https://github.com/sprouted-dev/garden https://github.com/sprouted-dev/garden/tree/main{/dir} https://github.com/sprouted-dev/garden/blob/main{/dir}/{file}#L{line}">
</head>
<body>Go package: sprouted.dev/<?= $package ?></body>
</html>
```

## Testing the Configuration

Once deployed, test with:

```bash
# Test weather package
curl -H "User-Agent: Go-http-client/1.1" "https://sprouted.dev/weather?go-get=1"

# Test sprout-cli package  
curl -H "User-Agent: Go-http-client/1.1" "https://sprouted.dev/sprout-cli?go-get=1"

# Test actual Go get
go get sprouted.dev/weather
go get sprouted.dev/sprout-cli
```

## Package Import Examples

After configuration, developers can import with clean paths:

```go
import (
    "sprouted.dev/weather"
    "sprouted.dev/sprout-cli"
)
```

Instead of the verbose:

```go
import (
    "github.com/sprouted-dev/garden/libs/weather"
    "github.com/sprouted-dev/garden/apps/sprout-cli"
)
```

## Subdirectory Mapping

The go-import system will automatically handle the subdirectory mapping:

- `go get sprouted.dev/weather` → clones GitHub repo → uses `/libs/weather/` subdirectory
- `go get sprouted.dev/sprout-cli` → clones GitHub repo → uses `/apps/sprout-cli/` subdirectory

This is handled by the module paths in our `go.mod` files and the replace directives during development.