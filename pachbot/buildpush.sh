
if [ -z $1 ]
then
      echo "USAGE: ./buildpush.sh <image tag>"
else
      docker build -t pskreter/pachbot:$1 .
      docker push pskreter/pachbot:$1
fi

