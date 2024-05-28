## Problem:
Most repos online, mostly Node.js projects, does not contain ".env.example" file
though environment variables are generally required for these projects to run. 

To fix this, you can do it manually by looking for envs, and adding them to your .env file,
but that is a boring and a time-consuming task.

Inorder to automate this task, I decided to create this tool and named it as "envy".
It is written in go, and makes use of go's powerful concurrency for efficiency.
Currently, this tool supports only nodejs projects. 

The goal of this progam/tool is to make setting up the environment variables 
for your project much easier and faster.

## "Envys" is available for installation via npm:
```
npm install envys
```

After installation, please add it to your package.json's scripts section:
```
"scripts": {
  "envys": "envys"
}
```
Once, you have added the script, you can run it by the following command:
```
npm run envys
```

The program will scan all directories and files in your project, extract environment variables, 
and write them to a ".env.example" file. This makes setting up the environment for your project
much easier and faster.


## Build and Run:
If you prefer to build the program directly from the source code, you can do so by 
cloning the repository and building the Go executable. Once built, you can run the
executable in your specific Node.js project directory. It scans all the dirs and files, 
extracts env variables, and then writes it in the ".env.example" file.


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
