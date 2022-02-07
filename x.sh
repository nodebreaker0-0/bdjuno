sudo tee /etc/systemd/system/hasura-actions.service > /dev/null <<EOF
[Unit]
Description=BDJuno hasura actions
After=network-online.target

[Service]
User=$USER
ExecStart=$GOPATH/bin/bdjuno hasura-actions
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
