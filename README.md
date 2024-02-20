
## Run Locally

Clone the project

```bash
  git clone https://github.com/Rpqshka/user-app
```
```
  cd user-app
```

Build

```bash
  docker-compose up --build
```


## Usage/Examples

1. Create User
```
localhost:8000/user/create
```
TEST INPUT:
```
{
    "id":"TestID",
    "first_name": "TestFN",
    "last_name": "TestLN",
    "age": 22
}
```

2. Get All Users
```
localhost:8000/user/all
```
3. Get Users By Date And Age
```
localhost:8000/user/search
```

TEST INPUT:
```
{
    "min_age": 20,
    "max_age": 23,
    "start_date": 1508420406,
    "end_date": 1708420406
}
```

TEST OUTPUT:
```
{
    "count": 1,
    "users": [
        {
            "Id": "TestID",
            "FirstName": "TestFN",
            "LastName": "TestLN",
            "Age": 22,
            "RecordingDate": 1708420406
        }
    ]
}
```