#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/letzy" -importPath github.com/jen6/letzy -srcPath "$SCRIPTPATH/src" -runMode dev
