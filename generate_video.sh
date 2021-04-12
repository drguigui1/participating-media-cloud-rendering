#!/bin/sh

if [ "$#" -ne 1 ]; then
    echo "Wrong number of parameters"
    exit 1
fi

ffmpeg -framerate $1 -i "videos/video_img%d.png" output.mp4
