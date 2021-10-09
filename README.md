    DOCUMENTATION
    
**Create User(POST Method)**

format: localhost:3000/user and Thunder client to post the JSON of user data in structure :
type User struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	Email string        `json:"email" bson:"email,omitempty"`
	Pass  string        `json:"pass" bson:"pass,omitempty"`
}

![image](https://user-images.githubusercontent.com/84318539/136668392-8281dac4-6fbe-4562-b91a-791c0f00aa47.png)
![image](https://user-images.githubusercontent.com/84318539/136668396-d557aa7f-65fa-49ef-aa45-d2a68481cddb.png)

**Create Post(POST Method)**

format: localhost:3000/posts and Thunder client to post the JSON of Post data in structure :
type Posts struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Captions  string        `json:"captions,omitempty" bson:"captions,omitempty"`
	Imageurl  string        `json:"imageurl" bson:"imageurl,omitempty"`
	Timestamp string        `json:"timestamp" bson:"timestamp,omitempty"`
	UserDat   string        `json:"user" bson:"user"`
}

![image](https://user-images.githubusercontent.com/84318539/136668449-d7cb2cc3-c74b-4900-824d-ed1057f9b009.png)
![image](https://user-images.githubusercontent.com/84318539/136668455-c545dbe0-fa38-4f1d-99cf-c0dcac27b0a2.png)

**Get user by user ID(GET Method)**
format: use gormat localhost:3000/user/{id} where {id} is the user id whose data is to be extracted and post it in browser or thunderclient GET method.
![image](https://user-images.githubusercontent.com/84318539/136668505-0344ab73-77be-4e74-9513-81fdadb65346.png)

**Get Post by post ID(GET Method)**
format: use gormat localhost:3000/posts/{id} where {id} is the post id whose data is to be extracted and post it in browser or thunderclient GET method.
![image](https://user-images.githubusercontent.com/84318539/136668529-4a799dd6-d5c6-4b87-b237-bef1355f9f4e.png)

**Get All post data of given user ID(GET Method)**
format: use gormat localhost:3000/posts/user/{id} where {id} is the user id whose all posts data is to be extracted and post it in browser or thunderclient GET method.
![image](https://user-images.githubusercontent.com/84318539/136668571-5e4b9d78-9ed7-4588-ac66-b5ed57b2e1d5.png)

		Other Technologies
*Mutex locks and unlocks have been used from sync imports to provide thread failsafe and avoid deadlocks

*Passwords of User are hashed and only hashed data is shown to developer to provide better privacy and security

_Mongodb port: 27017_

_Mongodb databast name APPOINTY_19BCE1259_db_

_HTTP Web server port: 3000_
