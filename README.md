## Problem:
Most repos online, mostly Node.js projects, does not contain ".env.example" file
though environment environment are generally required for these projects to run. 

To fix this, you can do it manually by looking for envs, and adding them to your .env file,
but that is a boring and a time-consuming task.

Inorder to automate this task, I decided to create this tool and named it as "envy".
It is written in go, and makes use of go's powerful concurrency for efficiency.
Currently, this tool supports only nodejs projects. 

Once you build the program, you can run it in your specific nodejs project directory.
It scans all the dirs and files, extracts env variables, and then writes it in the 
".env.example" file.


## Build and Run:
1. Clone the Repo:
```
git clone https://github.com/CoderParth/envy.git
```

2. Build the tool:
```
go build 
```

3. Copy the executable file "envy" and then run the tool in your specified directory:
```
./envy
```

After a successful run, ".env.example" will be created with all the envs 
written in the file. 


## Tests
After cloning the repo, you can run a test in this same cloned dir.

```
go test
```

This dir contains a nested dir called "test_dir", which contains some env variables
inside some files. After execution, something like below will be printed. 

```
DB_URL=
AWS_ACCESS_KEY_ID=
AWS_ACCESS_KEY_SECRET=
SENDGRID_API_KEY=
FROM_EMAIL=
API_KEY=
API_SECRET=
JWT_SECRET=
REDIS_URL=
COOKIE_SECRET=

PASS
ok      github.com/CoderParth/envy      0.559s
```
