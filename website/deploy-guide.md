# Sprouted.dev Deployment Guide

## Files Ready for Deployment

- `index.html` - Main website with Go module support
- `install.sh` - Installation script for Sprout CLI

## Deployment Checklist

### 1. Upload Files
Upload these files to your sprouted.dev web server:
- `index.html` → `/index.html` (root)
- `install.sh` → `/install.sh`

### 2. Web Server Configuration

#### For Nginx:
```nginx
server {
    server_name sprouted.dev;
    root /var/www/sprouted.dev;
    index index.html;
    
    # Handle Go module requests
    location ~ ^/(weather|sprout-cli)$ {
        if ($args ~ "go-get=1") {
            try_files /index.html =404;
        }
        # Regular requests can go to documentation or redirect
        return 301 https://pkg.go.dev/sprouted.dev$uri;
    }
    
    # Installation script
    location /install.sh {
        add_header Content-Type text/plain;
        try_files /install.sh =404;
    }
    
    # Static files
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    # Security headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header Referrer-Policy "no-referrer-when-downgrade" always;
}
```

#### For Apache:
```apache
<VirtualHost *:80>
    ServerName sprouted.dev
    DocumentRoot /var/www/sprouted.dev
    DirectoryIndex index.html
    
    # Handle Go module requests
    RewriteEngine On
    RewriteCond %{QUERY_STRING} go-get=1
    RewriteRule ^/(weather|sprout-cli)$ /index.html [L]
    
    # Regular package requests redirect to pkg.go.dev
    RewriteCond %{QUERY_STRING} !go-get=1
    RewriteRule ^/(weather|sprout-cli)$ https://pkg.go.dev/sprouted.dev$1 [R=301,L]
    
    # Installation script with correct MIME type
    <Files "install.sh">
        ForceType text/plain
    </Files>
</VirtualHost>
```

### 3. SSL Certificate
Ensure HTTPS is enabled (required for Go modules):
```bash
# Using Certbot/Let's Encrypt
certbot --nginx -d sprouted.dev
# or
certbot --apache -d sprouted.dev
```

### 4. Test Go Module Configuration

After deployment, test the Go module redirects:

```bash
# Test weather package
curl -H "User-Agent: Go-http-client/1.1" "https://sprouted.dev/weather?go-get=1"

# Test sprout-cli package
curl -H "User-Agent: Go-http-client/1.1" "https://sprouted.dev/sprout-cli?go-get=1"

# Test installation script
curl -fsSL https://sprouted.dev/install.sh | head -10
```

### 5. Verify Website Functionality

Visit https://sprouted.dev and check:
- [ ] Website loads correctly
- [ ] Dark/light mode toggle works
- [ ] Rotating headlines animate
- [ ] Installation commands are correct
- [ ] GitHub links point to correct repository
- [ ] Email signup form works (if connected to backend)

### 6. DNS Configuration

Ensure your DNS is pointing to the web server:
```
sprouted.dev    A       YOUR_SERVER_IP
www.sprouted.dev CNAME  sprouted.dev
```

## Go Module Testing

Once deployed, test the full Go module workflow:

```bash
# Test Go module imports work
cd /tmp
mkdir test-sprouted && cd test-sprouted
go mod init test

# This should work once your repo is public
go get sprouted.dev/weather
go get sprouted.dev/sprout-cli
```

## Post-Deployment Tasks

1. **Make Repository Public**: The Go modules won't work until the GitHub repository is public
2. **Create GitHub Release**: For the installation script to work, you'll need to create releases with binaries
3. **Update pkg.go.dev**: Go modules will automatically appear on pkg.go.dev once they're public
4. **Test Installation Script**: Once you have releases, test the full installation flow

## Analytics & Monitoring

Consider adding:
- Google Analytics or privacy-friendly analytics
- Error monitoring for the installation script
- Usage tracking for Go package downloads

## SEO & Performance

The website is already optimized with:
- Semantic HTML structure
- Meta descriptions
- Responsive design
- Fast loading (no external dependencies except Tailwind CDN)
- Proper heading hierarchy

## Email Collection

The email signup form is ready but needs backend integration. Consider:
- Netlify Forms
- ConvertKit
- Mailchimp API
- Simple email forwarding service

## Maintenance

Regular tasks:
- Monitor installation script success rates
- Update Go module documentation
- Keep website content fresh
- Monitor GitHub repository activity

The website is production-ready and includes all the features needed for a successful open source launch! 🚀