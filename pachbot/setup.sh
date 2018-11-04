# Create the needed repos
echo "Createing Repos...."
pachctl create-repo raw_data
pachctl create-repo parameters

# Add the data
echo "Adding Base Data...."
cd data
pachctl put-file raw_data master iris.csv -f noisy_iris.csv

cd parameters
pachctl put-file parameters master -f c_parameters.txt --split line --target-file-datums 1 
pachctl put-file parameters master -f gamma_parameters.txt --split line --target-file-datums 1


# Create pipelines
echo "Creating Pipelines...."
pachctl create-pipeline -f /pipelines/split.json 
pachctl create-pipeline -f /pipelines/model.json
pachctl create-pipeline -f /pipelines/test.json 
pachctl create-pipeline -f /pipelines/select.json
