#!/bin/bash

# Save private keys from Ganache CLI output to private_keys.txt
# this is only for teasing
ganache-cli | grep "Private Keys" -A 10 > private_keys.txt
