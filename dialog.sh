#!/bin/bash

# dialog.sh
# krusty / Jan 15, 2023

# https://stackoverflow.com/a/73344717

# DIALOGRC overrides:
IFS='' read -r -d '' dialogrc <<'EOF'
use_shadow = OFF
use_colors = ON
screen_color = (GREEN,BLACK,ON)
dialog_color = screen_color
border_color = screen_color
border2_color = screen_color
inputbox_color = screen_color
button_active_color = screen_color
button_inactive_color = screen_color
button_key_active_color = screen_color
button_key_inactive_color = screen_color
button_label_active_color = screen_color
button_label_inactive_color = screen_color
EOF

input_passwd=$(
    3>&1 >/dev/tty \
    DIALOGRC=<(printf '%s' "$dialogrc") \
    dialog \
    --output-fd 3 \
    --no-lines \
    --insecure \
    --passwordbox $' Please enter the passphrase to\n protect your key' 8 60
) || exit 1

printf '\ninput_password=%q\n' "$input_passwd"
