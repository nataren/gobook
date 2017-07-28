#! /bin/bash

TZ=US/Eastern ./clock1 -port 8010 &
TZ=europe/London ./clock1 -port 8020 &
TZ=Asia/Tokyo ./clock1 -port 8030 &
