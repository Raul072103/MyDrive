import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Read path from ENV variable
file_path = os.getenv('FILE_SYSTEM_ROOT_FOLDER')

print(file_path)

if not file_path:
    print("FILE_PATH environment variable not set.")
    exit(1)

# Function to create large files
def create_large_file(file_name, size_in_mb):
    with open(file_name, 'wb') as f:
        # Write random data to the file to create the required size
        f.write(os.urandom(size_in_mb * 1024 * 1024))
    print(f"Created file: {file_name} with size: {size_in_mb}MB")

# Example: Create 5 large files (1GB each)
for i in range(1, 6):
    file_name = os.path.join(file_path, f"large_file_{i}.bin")
    create_large_file(file_name, 1024)  # 1024MB = 1GB

print("Large files created successfully.")
