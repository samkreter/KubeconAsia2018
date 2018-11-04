
# Update all pipelines
pachctl update-pipeline -f /pipelines/split.json
pachctl update-pipeline -f /pipelines/model.json
pachctl update-pipeline -f /pipelines/test.json
pachctl update-pipeline -f /pipelines/select.json

# for pipeline in /pipelines/*.json
# do 
#     echo "INFO: Updating pipeline $pipeline"
#     pachctl update-pipeline -f $pipeline
# done