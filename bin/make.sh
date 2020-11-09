first="$1"
second="$2"
command=$first
current_path="${PWD}"
echo $first
if [ $first = "--win" ]; then
    current_path=$(cygpath -w $(pwd))
    command=$second
fi
docker run --rm -v $current_path:/app tooling make $command