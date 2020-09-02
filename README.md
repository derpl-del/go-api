# go-api

========================================================================

API generatesender POST
Endpoint :
localhost:9000/api/v1/generatesender

Request :
{"sender":"youremail","password":"yourpassword"}

Desc :
API for set your email sender while sending invoice and verifyuser

========================================================================

API createuser POST
Endpoint :
localhost:9000/api/v1/createuser

Request :
{"user_name":"banana1","wallet":10,"mail":"youremail@gmail.com"}

Response :
{"error_code":"0000","error_msg":"success"}

Desc :
API for creating username with validation, if validation equal true verification email will be send

========================================================================

API Verify User POST
Endpoint :
localhost:9000/api/v1/verifyuser

Request :
localhost:9000/api/v1/verifyuser?req=Vm3RV/7lfQ==

Desc :
API for verify your email, while trigger data will be insert to database, if user already trigger more than 1 the request will be ignored

========================================================================

API createproduct POST
Endpoint :
localhost:9000/api/v1/createproduct

Request :
{"name":"c3","amount":10,"price":1}

Response :
{"error_code":"0000","error_msg":"success"}

Desc :
API for creating product with validation, if validation equal true product will be created in file (.JSON)

========================================================================

API buyproduct POST
Endpoint :
localhost:9000/api/v1/buyproduct

Request :
{"user_name":"banana1","product_name":"c3","amount":1,"price":1}

Response :
{"error_code":"0000","error_msg":"success"}

Desc :
API for buy some product with validation, if validation equal true email invoice will be sended to email that username has registered
