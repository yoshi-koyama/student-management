# Group Students
# POST /token

+ Request (application/json)
            
+ Response 200 (application/json)

            {
              "token": "your jwt token"
            }

# POST /api/students

+ Request Sample
```
curl -XPOST -H "Content-type: application/json" -d '{
  "name":"yoshi",
  "nrc":"1234",
  "age":20,
  "gender":"male",
  "password":"password"
}' 'http://localhost:8080/api/students'
```

+ Request (application/json)

            {
              "name": "yoshihito koyama",
              "nrc": "abcde",
              "gender": "male",
              "password": "abcdefghijklmn"
            }

+ Response 200 (application/json)

            {
              "message": "create student success"
            }

# GET /api/students

+ Request Sample
```
curl http://localhost:8080/api/students
```

+ Parameters

  + gender (string, optional) - Specify to filter students.

+ Response 200 (application/json)

            [
              {
                "id": 1,
                "name": "yoshihito koyama",
                "nrc": "abcde",
                "gender": "male",
                "password": "abcdefghijklmn"
              }
            ]

# GET /api/students/{id}

+ Request Sample
```
curl http://localhost:8080/api/students/1
```

+ Parameters

  + id: 1 (number) - An unique identifier of the student.


+ Response 200 (application/json)

            {
              "id": 1,
              "name": "yoshihito koyama",
              "nrc": "abcde",
              "gender": "male",
              "password": "abcdefghijklmn"
            }
