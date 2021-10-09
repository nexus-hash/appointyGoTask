# AppointyTask API


### Details About the Endpoints

- POST /users: To create a new user  
Checks if the user already exists in the database.  
If the user does not exist, creates a new user.  

Request Body:  
<pre><code>
  {
    "_id"        : string,     //Later Converted to Mongo ObjectId
    "name"      : string,
    "username"  : string,
    "password"  : string,
  }
</code></pre>


- GET /users/{id}: To get a user by id
- POST /posts: To create a new post  

- GET /posts/{id}: To get a post by id
- GET /posts/users/{id}: To get all posts by user id


