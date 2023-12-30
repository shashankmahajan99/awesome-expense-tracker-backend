# Create the directory structure for the Google API files

mkdir -p google/api

# Download annotations.proto

curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o google/api/annotations.proto

# Download http.proto

curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o google/api/http.proto
