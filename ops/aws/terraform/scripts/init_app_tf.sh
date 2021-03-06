#!/bin/sh

# Generate Terraform configuration for the combination of a Soapbox
# application, environment, and type of infrastructure to manage
# (network or deployment).
#
# The only output to stdout is the name of a directory containing the
# generated config files, making it suitable for use in scripts. Any
# output from Terraform itself is sent to stderr.

set -e

STATE_BUCKET="${S3_BUCKET-soapbox-app-tf-state}"
REGION="${REGION-us-east-1}"
DYNAMO_TABLE="${DYNAMO_TABLE-soapbox-app-state-locking}"
PWD=$(pwd)

usage() {
    echo "Usage: $(basename $0) [-a application] [-e environment] [-t <network|deployment>]" 1>&2
    exit 1
}

while getopts ":a:e:t:" opt; do
    case "$opt" in
	a)
	    APP="$OPTARG"
	    ;;
	e)
	    ENV="$OPTARG"
	    ;;
        t)
            TYPE="$OPTARG"
            ;;
        *)
            usage
            ;;
    esac
done

if [ -z "$APP" ] || [ -z "$ENV" ] || [ -z "$TYPE" ]; then
    usage
fi

# Get a temp dir, copy the required tf files
APP_TMP=$(mktemp -d)
cp -R $PWD/../$TYPE/* $APP_TMP/

# Echoing temp dir path to stdout, so controlling application can set
# it as the working directory for running Terraform.
echo $APP_TMP

# Change to the temp dir
if [ "$TYPE" = "deployment" ]; then
  cd $APP_TMP/asg
else
  cd $APP_TMP
fi

# Replaces a variable name with a value in a given file, editing it
# in-place. The variable name should be surrounded by '#'s in the
# file, for example: #APP_NAME#
replace_in_place () {
    local name="$1"
    local value="$2"
    local filename="$3"
    sed -e "s/#${name}#/$value/" "$filename" > "$filename.new"
    mv -- "$filename.new" "$filename"
}

# Populate backend.tfvar.sample with appropriate valuess
replace_in_place APP $APP backend.tfvars.sample
replace_in_place ENV $ENV backend.tfvars.sample
replace_in_place STATE_BUCKET $STATE_BUCKET backend.tfvars.sample
replace_in_place REGION $REGION backend.tfvars.sample
replace_in_place DYNAMO_TABLE $DYNAMO_TABLE backend.tfvars.sample

# Ceremonial promotion of .sample to bonafide .tfvars file
mv backend.tfvars.sample backend.tfvars

# Redirect stdout to stderr from here one, so a calling script doesn't
# accidentally capture terraform's output when trying to get the value
# of the temp dir
exec 1>&2

# Get any required terraform modules
terraform get -no-color

# Initialize the backend
terraform init -backend-config=backend.tfvars -no-color
