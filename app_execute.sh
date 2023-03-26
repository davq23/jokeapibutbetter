#!/bin/bash
set -x

scp -r /frontend $FTP_USERNAME:$FTP_PASSWORD@$FTP_HOST

/main