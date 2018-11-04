
# Update all pipelines
pachctl update-pipeline -f /pipelines/model.json




# for pipeline in /pipelines/*.json
# do 
#     echo "INFO: Updating pipeline $pipeline"
#     pachctl update-pipeline -f $pipeline
# done