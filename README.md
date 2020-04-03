# lockdown
The lockdown app to help administration
##
```aidl
go build
./lockdown
```
###install httpie
````
http -j POST :8080/open/register-user-details  tehsil=jaipur dealerType=sa deliveryLocation=jabalpur mobile=9221212121212121
````
You will get the following response
```
{
    "dealerType": "sa",
    "deliveryLocation": "jabalpur",
    "mobile": "9221212121212121",
    "tehsil": "jaipur"
}
```


To download the csv
```
http --download GET :8080/download-csv
```