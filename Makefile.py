import os
import shutil
import subprocess
import sys


"""
 

go test ./pkg/...
go test ./pkg/toolfile/...

"""
def test():
    print("Testing...")
    env = os.environ.copy()
    #env['CGO_ENABLED'] = '1' #for -race flag #cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
    command = ['go', 'test'
               #, '-race'
               , '-timeout=60s', '-count=1', './...']
    subprocess.run(command, env=env) #, "-v"

def help():
    print("Usage:")
    print("  python build.py test     - Run test")
    print("  python build.py help     - Display this help message")
    

if len(sys.argv) > 1:
    command = sys.argv[1]
    if command == "test":
        test() 
    elif command == "help":
        help() 
    else:
        help()
        exit(1)
else:
    help()



"""
git add .
git commit -m "Initial commit"
git tag v1.0.0-beta

"""


#BINARY_NAME = "app.exe" if os.name == "nt" else "app"

# def build():
#     print("Building the binary...")
#     subprocess.run(["go", "build", "-C", "cmd/app", "-o",f"./../../dist/{BINARY_NAME}" ])

# def run():
#     build()
#     print("Running the application...")
#     subprocess.run([f"./dist/{BINARY_NAME}"])

# def clean():
#     print("Cleaning up...")
#     if os.path.isdir("./dist"):# os.path.exists("./dist"):
#         shutil.rmtree("./dist") #os.remove("./dist")
#     os.makedirs("./dist")
# def rebuild():
#     clean()
#     build()

