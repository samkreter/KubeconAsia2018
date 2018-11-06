
# Update all pipelines
echo "Updating all pipelines...."
pachctl update-pipeline -f ../pipelines/split.json
pachctl update-pipeline -f ../pipelines/model.json
pachctl update-pipeline -f ../pipelines/test.json
pachctl update-pipeline -f ../pipelines/select.json
