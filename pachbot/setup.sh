# Create the needed repos
pachctl create-repo raw_data
pachctl create-repo parameters

# Add the data
pachctl put-file raw_data master iris.csv -f data/noisy_iris.csv 
pachctl put-file parameters master -f data/parameters/c_parameters.txt --split line --target-file-datums 1 
pachctl put-file parameters master -f data/parameters/gamma_parameters.txt --split line --target-file-datums 1

# Create pipelines
pachctl create-pipeline -f /pipelines/split.json 
pachctl create-pipeline -f /pipelines/model.json
pachctl create-pipeline -f /pipelines/test.json 
pachctl create-pipeline -f /pipelines/select.json
