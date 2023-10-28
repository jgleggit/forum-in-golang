#!/bin/bash

# Create the SQLite3 database
sqlite3 forum-db.sqlite3 <<EOF
.quit
EOF
