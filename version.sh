#!/usr/bin/env bash
###############################################################
# version.sh :                                                #
#             This module is used by deployment to generate   #
#             version information for application to be used  #
#             in production.                                  #
# warning :                                                   #
#           This module can not ouput any error or debug msg  #
#           As the output of this module is used by the build #
#           pipeline to name the final docker image.          #
#           This module should recover from all errors and    #
#           generate a unique image name each time            #
###############################################################

# version module configuration
SUFFIX_SEP=${SUFFIX_SEP:--}
VERSION_SUFFIX=${VERSION_SUFFIX}

# Get current commit id
GIT_COMMIT=$(git rev-parse HEAD)
GIT_COMMIT_SHORT_SHA=$(git rev-parse --short HEAD)

# Buil dtimestamp
BUILD_TIMESTAMP=$(date +%FT%T%z)

# Get current git tag
GIT_TAG=$(git describe --abbrev=0 --tags 2>/dev/null)

# GIT TAGs
if [ -z "${GIT_TAG}" ]; then
    GIT_TAG="0.0"
    COMMIT_COUNT=$(git rev-list HEAD --count 2>/dev/null)
    if [ -z "${COMMIT_COUNT}" ]; then
        COMMIT_COUNT=0
    fi
else
    # Get commit count since last tag
    COMMIT_COUNT=$(git rev-list ${GIT_TAG}.. --count 2>/dev/null)
    if [ -z "${COMMIT_COUNT}" ]; then
        COMMIT_COUNT=0
    fi
fi

VERSION="${GIT_TAG}.${COMMIT_COUNT}"

VERSION_SUFFIX="${VERSION_SUFFIX}${SUFFIX_SEP}${GIT_COMMIT_SHORT_SHA}"

# Current Branch
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Check current branch
# Master and release branch is not included in the prefix
if [ "${GIT_BRANCH}" != "master" ] && [ "${GIT_BRANCH}" != "release" ] && [ "${GIT_BRANCH}" != "main" ] && [ -n "${GIT_BRANCH}" ]; then
    VERSION_SUFFIX="${VERSION_SUFFIX}${SUFFIX_SEP}${GIT_BRANCH}"
fi

# Full version specifier inluding branch information
VERSION_FULL=${VERSION}${VERSION_SUFFIX}

# Print current version to console. This will be used by build pipeline to decide name the final image
echo "${VERSION_FULL}"
