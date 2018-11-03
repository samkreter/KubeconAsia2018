
# Update all pipelines
for pipeline in /pipelines/*.json
do 
    echo "INFO: Updating pipeline $pipeline"
    pachctl update-pipeline -f $pipeline
done