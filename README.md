# lockdown
The lockdown app to help administration
##
```aidl
go build
ADMIN=admin PASSWORD=password ./lockdown
```
### httpie way
````
http -j POST :8080/open/register-user-details  < request.json
````
here request.json is POST body
### curl way
```
curl -X POST \
  http://localhost:8080/open/register-user-details \
  -H 'Accept: application/json' \
  -d '{
        "deliveryLocation": {
          "area": "Khana Galli",
          "city": "Jaipur"
        },
        "shopDetails": {
          "name": "Ramlal mitaiwaala",
          "address": "RustomJee Area, Kalakand",
          "ownerMobile": "90881910",
          "email": "jackson@gmail.com",
          "type": "Retail"
        },
        "homeDeliveryInfo": {
          "homeDeliveryNumber": "98001010101",
          "agentInfo": {
            "agentName": "Ramchandani",
            "agentAge": 45,
            "agentMobile": "99092029292"
          },
          "vehicleInfo": {
            "vechicleType": "Car",
            "VehicleNumber": "MH091111"
          }
        }
      }
'
```

You will get the following response
```
{
{
  "City": "VijayWada",
  "DealerType": "Retail",
  "DeliveryLocation": "Jaipur",
  "Mobile": "89289211",
  "ShopDetails": {
    "name": "Ramlal mitaiwaala",
    "address": "RustomJee Area, Kalakand",
    "ownerMobile": "90881910",
    "email": "jackson@gmail.com",
    "type": "Retail"
  },
  "HomeDeliveryInfo": {
    "homeDeliveryNumber": "98001010101",
    "agentInfo": {
      "agentName": "Ramchandani",
      "agentAge": 45,
      "agentMobile": "99092029292"
    },
    "vehicleInfo": {
      "type": "Car",
      "number": "MH091111"
    }
  },
  "RegistrationDate": "Monday, 06-Apr-20 23:12:19 IST",
  "Id": 10
}
```


To download the csv
```
http --download GET :8080/secure/download-csv http --download GET :8080/secure/download-csv -a admin:pwd
```

To add a secret if deploying in cloud
```
kubectl create secret generic lockdownsecret --from-literal user=Bruce --from-literal password=verystrongpassword
```

Remaining  Tasks
1. Make GenerateId as singleton
2. Refactor and move other handler methods to use repo methods
3. Enhance request to include more fields
4. Mock Dates in csv
