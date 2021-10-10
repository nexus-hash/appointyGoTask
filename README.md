# AppointyTask API

### How to Compile

- Export All the required environment variables.
<pre><code>
  MONGO_URI = Database URI
  PORT = Port to run the server on
  DATABASE_NAME = Database name
</code></pre>

- Navigate to api directory and Run the command.
<pre><code>
  go run main.go
</code></pre>

- The server will be running on the port specified using the environment variable PORT.
- The API will through an runtime error if the environment variables are not set.


### Details About the Endpoints

- [ POST ] /users: To create a new user.
<br></br>
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


- [ GET ] /users/{id}: To get a user by id.
<br></br>
Returns the user with the given id.
<br></br>
Request Body:
<pre><code>
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "name"      : string,
    "username"  : string,
    "password"  : nil,
  }
</code></pre>
- [ POST ] /posts: To create a new post.
<br></br>
Creates a new post.
<br></br>
Response :  
<pre><code>
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObectjId,     
  }
</code></pre>
- [ GET ] <b>/posts/{id}</b>: To get a post by id.
<br></br>
Returns the post with the given id.  
<br></br>
Response :
<pre><code>
  {
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObectjId,     
  }
</code></pre>
- [ GET ] /posts/users/{uid}: To get all posts by userid.
<br></br>
For the specified userid, returns the posts of the user based on Page No. and Page Limit.
<br></br>
Request Body:
<pre><code>
  {
    "page"      : int,
    "limit"     : int,
  }
</code></pre>
Response :
<pre><code>
  [{
    "_id"       : string,     //Later Converted to Mongo ObjectId
    "caption"   : string,
    "imageurl"  : string,
    "postedat"  : string,
    "userid"    : mongo.ObectjId,     
  }, Upto the Page Limit or till the end of the posts]
</code></pre>


