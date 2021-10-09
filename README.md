# ibe

DOCUMENTATION>

Structures:
defined in ibe/models

Functions:

Create user:
via localhost:3000/user and ThunderPost client,you can POST an Instagram user account with given data of json structure.

Get User by ID:
via localhost:3000/user/{id}, you can GET all JSOn details stored in mongo db of the user with USER ID Given.

Create post:
via localhost:3000/posts and ThunderPost client,you can POST an Instagram post of json structure.

Get Post by ID:
via localhost:3000/posts/{id}, you can GET all JSOn details stored in mongo db of the Instagram Post with Posr ID Given.

Get All Posts of a given user.
via localhost:3000/posts/users/{id}, you can get JSON info sotred of all posts made by the user of USer ID given in link.

REST Api port= 3000
Mongo db port= 27017
Mongo db Database name= APPOINTY_19BCE1259_db


HOW TO USE:

Create user:
POST the user data in json format provided in ibe/models.
![image](https://user-images.githubusercontent.com/84318539/136667055-f6aa8436-15bd-42b7-994c-9bd0efb60bb9.png)

![image](https://user-images.githubusercontent.com/84318539/136667071-1d6ca35c-aa1b-4d28-8474-ad4a13702ec4.png)
![image](https://user-images.githubusercontent.com/84318539/136667079-b22e1ab4-7fc8-473b-83e5-3b1d21cb79ae.png)

Find user by User ID
Use Get in Thunder Client or post the format of link localhost:3000/user/{id} in browser after running the go file.
![image](https://user-images.githubusercontent.com/84318539/136667419-9f93995a-8007-4c21-9b0a-bc4c5a5e1e99.png)

Create post
Same as Create user, instead follow format of localhost:3000/posts in Thunder client
![image](https://user-images.githubusercontent.com/84318539/136667452-200f1c0d-9948-4400-aa6a-65c59fedceeb.png)
![image](https://user-images.githubusercontent.com/84318539/136667459-3f1024d0-f3b1-4ec4-8693-e86417256cfb.png)

Get Post by post id
same as Find user by user id, instead follow localhost:3000/posts/{id} format.

![image](https://user-images.githubusercontent.com/84318539/136667491-99762daf-a758-4d0f-874a-70ee9c9b2dfb.png)


get all posts of a user
follow format of localhost:3000/posts/users/{id} where id is requested user id whose posts are to be extracted.

![image](https://user-images.githubusercontent.com/84318539/136667540-b7a204d5-bc91-448c-b475-3c5f08bb5530.png)


