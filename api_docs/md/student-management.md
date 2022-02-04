# Group Students
# POST /token

+ Request (application/json)
            
+ Response 200 (application/json)

            {
              "token": "your jwt token"
            }

# POST /api/register

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
