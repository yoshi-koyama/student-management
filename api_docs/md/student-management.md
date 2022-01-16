# Group Students
# POST /register

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


# GET /students

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

# GET /students/{id}

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
