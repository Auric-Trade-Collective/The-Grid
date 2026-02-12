# Api Spec
This is an API spec meant to keep the frontend and backend matching whilst
Nox is in development.

**Note this spec covers how objects look in the DB as well as what is exposed 
by the API**

## User

**Full User Object**
```json
{
    "username": "",
    "hash": "",

    "email": "" | null,
    "recoveryCode": "" | null,

    "bannerLocation": "",
    "pfpLocation": "",
}
```

**API Exposed Object**
```json
{
    "username": "",
    "email": "" | null,
}
