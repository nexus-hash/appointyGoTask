# AppointyTask API

### How to Compile

- Export All the required environment variables
<pre><code>
  MONGO_URI = Database URI
  PORT = Port to run the server on
  DATABASE_NAME = Database name
</code></pre>

- Navigate to api directory and Run the command
<pre><code>
  go run main.go
</code></pre>

- The server will be running on the port specified using the environment variable PORT
- The API will through an error if the environment variables are not set


### Details About the Endpoints

- [POST] /users: To create a new user   {Implemented and Available}  
Checks if the user already exists in the database.
If the user does not exist, creates a new user.
<br></br>
Request Body:
<pre><code>
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "name"      : string,
    "username"  : string,
    "password"  : string,
  }
</code></pre>


- GET /users/{id}: To get a user by id  {Implemented and Available}
<pre><code>
  Returns the user with the given id.

  Request Body:
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "name"      : string,
    "username"  : string,
    "password"  : nil,
  }
</code></pre>
- POST /posts: To create a new post    {Implemented and Available}  
<pre><code>
  Creates a new post.

  Response :
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObecjId,     
  }
</code></pre>
- GET /posts/{id}: To get a post by id    {Implemented and Available}
<pre><code>
  Returns the post with the given id.

  Response :
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObecjId,     
  }
</code></pre>
- GET /posts/users/{uid}: To get all posts by userid  {Implemented and Available}
<pre><code>
  For the specified userid, returns the posts of the user based on Page No. and Page Limit.

  Request Body:
  {
    "page"      : int,
    "limit"     : int,
  }

  Response :
  [{
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObecjId,     
  }, Upto the Page Limit or till the end of the posts]
</code></pre>


