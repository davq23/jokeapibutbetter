#!/bin/bash
set -x

scp -r /frontend -P $FTP_PORT $FTP_USERNAME:$FTP_PASSWORD@$FTP_HOST$FRONTEND_ASSETS_URL

/main