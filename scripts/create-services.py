#!/usr/bin/env python3
import argparse
import os
from colorama import Fore, Back, Style

supported_services = [ 'auth', 'users', 'vms' ]

parser = argparse.ArgumentParser(description='Create GO API services from OpenAPI definitions.')
parser.add_argument('--api-dir', help='Diretory of the API specifications', required=True, dest='apiDir', action='store')
parser.add_argument('--sdk-dir', help='Output Diretory', required=True, dest='sdkDir', action='store')
parser.add_argument('--service', help='The service name to process or "all" to process all supported services', required=True, dest='serviceName', action='store')
parser.add_argument('--version', help='The API version', dest='version', default='v1.0.0-beta', action='store')
args = parser.parse_args()

args.apiDir = os.path.abspath(args.apiDir)
if not os.path.exists(args.apiDir):
    msg = f"API Directory [{args.apiDir}] does not exist"
    raise Exception(msg)

if args.serviceName != 'all' and args.serviceName not in supported_services:
    msg = f"Unsupported service name [{args.serviceName}]"
    raise Exception(msg)

args.sdkDir = os.path.abspath(args.sdkDir)

print(Fore.MAGENTA + f'Creating API defintions from [{args.apiDir}] in [{args.sdkDir}]\n')

for service in [service for service in supported_services if args.serviceName == "all" or service == args.serviceName]:
    print(Fore.MAGENTA + f"Processing service {service}\n")

    # autorest <apiDir>/specification/<serviceName>/readme.md --latest --go --go-sdk-folder=<sdkDir> --package-version=<version> --user-agent='QNAP-QVS-SDK-For-Go/<version> services' --version=2
    cmd = f"autorest {args.apiDir}/specification/{service}/readme.md --latest --go --go-sdk-folder={args.sdkDir} --package-version={args.version} --user-agent='QNAP-QVS-SDK-For-Go/{args.version} services' --version=2"
    print(cmd)
    os.system(cmd)

    print('')

cmd = f"gofmt -w {args.sdkDir}/services"
print(Fore.MAGENTA + cmd + "\n")
os.system(cmd)

if not os.path.exists(f"{args.sdkDir}/go.mod"):
    print(Fore.MAGENTA + f"Initializing go.mod at {args.sdkDir}")

    os.chdir(f"{args.sdkDir}")
    os.system("go mod init github.com/qnap/qvs-sdk-for-go")
    os.system("go mod tidy")
