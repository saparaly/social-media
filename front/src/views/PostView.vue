<template>
    <div class="space"></div>
    <div class="post">
      <div class="container">
        <p>{{ errorMessage }}</p>
        <div class="post__container" v-if="!errorMessage">
          <div class="post__header">
            <div class="post__header-container">
              <h1 class="title">{{ post.title }}</h1>
              <!-- edit -->
              <div class="edit__container">
                <div class="edit" @click="edit" >
                  <svg width="800px" height="800px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title/><g id="Complete"><g id="edit"><g><path d="M20,16v4a2,2,0,0,1-2,2H4a2,2,0,0,1-2-2V6A2,2,0,0,1,4,4H8" fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><polygon fill="none" points="12.5 15.8 22 6.2 17.8 2 8.3 11.5 8 16 12.5 15.8" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g></g></g></svg>
                </div>
                <div class="edit" @click="deletePost">
                  <svg width="800px" height="800px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M5.16565 10.1534C5.07629 8.99181 5.99473 8 7.15975 8H16.8402C18.0053 8 18.9237 8.9918 18.8344 10.1534L18.142 19.1534C18.0619 20.1954 17.193 21 16.1479 21H7.85206C6.80699 21 5.93811 20.1954 5.85795 19.1534L5.16565 10.1534Z" stroke="#000000" stroke-width="2"/><path d="M19.5 5H4.5" stroke="#000000" stroke-width="2" stroke-linecap="round"/><path d="M10 3C10 2.44772 10.4477 2 11 2H13C13.5523 2 14 2.44772 14 3V5H10V3Z" stroke="#000000" stroke-width="2"/></svg>
                </div>
              </div>
            </div>
          </div>
          <div class="post__subconainer">
            <div class="add__todo" v-if="post.location">
              <div class="icon" >
                <svg width="800px" height="800px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M4 11h12" stroke="#000000" stroke-width="1.5" stroke-miterlimit="10" stroke-linecap="round"/><path d="m15 15 2 2 4-4" stroke="#000000" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 7h12M4 15h8" stroke="#000000" stroke-width="1.5" stroke-miterlimit="10" stroke-linecap="round"/></svg>
              </div>
              <div class="p">{{ post.location }}, {{ post.date }}</div>
            </div>
            <div class="small">{{ post.created }}</div>
          </div>

          <div class="img" v-if="post.img">
            <img :src="post.img" alt="">
          </div>
          <div class="desc"> {{ post.description }} </div>
          <div class="tags" v-if="post.tags">
            <span v-for="tag in post.tags">
              {{ tag }}
            </span>
          </div>
          <div class="username">created by <span>{{postuserrole}}</span> @{{ usernamepost }}</div>

          <div class="reaction">
            <div class="react">
              <form action="">
                <button @click.prevent="likePost">
                  <svg fill="#000000" height="800px" width="800px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 297.221 297.221" xml:space="preserve"><g><path d="M283.762,32.835c2.705-1.913,3.346-5.658,1.432-8.363c-1.914-2.705-5.657-3.347-8.363-1.432l-14.984,10.602   c-2.705,1.913-3.346,5.658-1.432,8.363c1.169,1.652,3.022,2.535,4.902,2.535c1.198,0,2.408-0.358,3.461-1.104L283.762,32.835z"/><path d="M244.064,29.387c0.695,0.262,1.409,0.386,2.11,0.386c2.428,0,4.713-1.484,5.617-3.891l6.46-17.182   c1.166-3.101-0.403-6.561-3.505-7.727c-3.101-1.167-6.562,0.404-7.728,3.505l-6.46,17.182   C239.393,24.761,240.962,28.221,244.064,29.387z"/><path d="M291.223,55.611c-0.041,0-0.082,0-0.124,0l-18.351,0.154c-3.313,0.067-5.944,2.605-5.877,5.918   c0.066,3.271,2.739,5.928,5.997,5.928c0.041,0,0.082,0,0.124,0l18.351-0.313c3.313-0.068,5.944-2.732,5.877-6.045   C297.154,57.982,294.481,55.611,291.223,55.611z"/><path d="M254.2,147.154c-3.073-2.433-6.711-4.089-10.557-4.867c0.254-0.46,0.491-0.928,0.715-1.403l2.408-2.408   c9.274-9.275,10.248-23.874,2.264-33.961c-3.769-4.761-9.001-7.925-14.812-9.106c0.415-0.764,0.783-1.545,1.117-2.338   c6.316-9.149,6.213-21.445-0.782-30.283c-3.77-4.764-9.004-7.938-14.818-9.117c4.8-8.826,4.187-19.826-2.225-27.925   c-4.848-6.125-12.109-9.639-19.923-9.639c-6.257,0-12.16,2.236-16.792,6.33c-0.701-3.979-2.363-7.822-5.012-11.169   c-4.849-6.125-12.11-9.638-19.924-9.639l0,0c-6.79,0-13.164,2.635-17.947,7.418l-60.84,60.84l-0.232-8.12   c-0.107-13.83-11.392-25.049-25.247-25.049c-13.604,0-24.729,10.815-25.229,24.298l-12.146,88.306l-9.983,11.604   c-5.983,6.957-5.582,17.481,0.915,23.962L19.987,199.7c-4.574,6.881-3.773,16.266,2.206,22.23l69.667,69.557   c3.329,3.321,7.748,5.148,12.446,5.148c3.857,0,7.668-1.295,10.729-3.645l14.544-11.168c13.991-3.305,29.416-10.813,45.874-22.33   c14.371-10.058,29.962-23.46,46.337-39.836l34.631-34.631c5.107-5.107,7.795-12.188,7.375-19.427   C263.376,158.371,259.879,151.649,254.2,147.154z M188.124,32.009c2.603-2.602,6.032-3.903,9.462-3.903   c3.915,0,7.831,1.695,10.515,5.086c4.256,5.377,3.51,13.18-1.339,18.028l-6.177,6.176c-0.952,0.635-1.879,1.314-2.747,2.083   c-0.701-3.98-2.364-7.823-5.013-11.169c-3.257-4.114-7.604-7.043-12.475-8.527L188.124,32.009z M146.397,17.532   c2.602-2.602,6.032-3.903,9.462-3.903c3.916,0.001,7.831,1.696,10.515,5.087c4.256,5.377,3.51,13.179-1.339,18.027l-70.919,70.186   l-0.233-8.119c-0.061-7.825-3.7-14.812-9.356-19.405L146.397,17.532z M13.624,176.391c-2.082-2.078-2.209-5.41-0.291-7.64   l12.281-14.277c0.006-0.007,0.011-0.017,0.012-0.026l12.72-92.483c0-7.286,5.961-13.247,13.247-13.247   c7.286,0,13.248,5.961,13.248,13.247L65.186,74c-11.988,1.646-21.322,11.733-21.78,24.057l-12.145,88.307l-3.533,4.108   L13.624,176.391z M247.935,176.539l-34.63,34.631c-29.577,29.577-60.494,53.318-87.653,59.237   c-0.825,0.181-1.601,0.528-2.271,1.043l-15.655,12.022c-1.014,0.779-2.219,1.162-3.419,1.162c-1.443,0-2.88-0.555-3.968-1.641   l-69.671-69.56c-2.083-2.077-2.21-5.409-0.291-7.64l12.28-14.276c0.007-0.008,0.011-0.017,0.013-0.026l12.719-92.483   c0-7.286,5.962-13.248,13.248-13.248c7.286,0,13.247,5.962,13.247,13.248l0.626,21.824c0.104,3.626,3.087,5.987,6.191,5.987   c1.514,0,3.058-0.563,4.309-1.813l70.431-70.431c2.603-2.603,6.031-3.903,9.462-3.903c3.915,0,7.831,1.695,10.515,5.086   c4.256,5.377,3.509,13.18-1.34,18.028l-48.518,48.518c-2.519,2.52-2.519,6.603,0,9.121l0,0c1.275,1.275,2.946,1.913,4.617,1.913   s3.343-0.638,4.617-1.913l62.374-62.373c2.602-2.603,6.031-3.903,9.462-3.903c3.915,0.001,7.831,1.696,10.515,5.087   c4.256,5.376,3.509,13.179-1.34,18.027l-62.081,62.081c-2.553,2.554-2.553,6.692,0,9.246c1.258,1.258,2.906,1.887,4.556,1.887   c1.648,0,3.297-0.629,4.555-1.887l48.811-48.81c2.603-2.603,6.032-3.903,9.462-3.903c3.915,0,7.831,1.695,10.515,5.087   c4.256,5.376,3.509,13.179-1.34,18.027l-48.349,48.35c-2.612,2.611-2.612,6.847,0,9.458l0.078,0.079   c1.207,1.207,2.789,1.81,4.37,1.81c1.582,0,3.164-0.603,4.37-1.81l29.974-29.974c2.701-2.701,6.317-4.129,9.921-4.129   c2.867,0,5.726,0.904,8.107,2.789C253.114,161.598,253.508,170.967,247.935,176.539z"/></g></svg>
                </button>
              </form>
              {{likes}} 
            </div>
            <div class="react">
              <form action="">
                <input type="text" value="dislike" name="reaction" hidden>
                <button @click.prevent="dislikePost">
                  <svg fill="#000000" height="800px" width="800px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 58.036 58.036" xml:space="preserve"><g><g><path d="M53,45.518v-16.5c0-10.201-8.301-18.5-18.505-18.5H23.505C13.301,10.518,5,18.817,5,29.018v16.5H0v2h6h46h6.036v-2H53z     M7,45.518v-16.5c0-9.098,7.404-16.5,16.505-16.5h10.989c9.101,0,16.505,7.402,16.505,16.5v16.5H7z"/><polygon points="12,33.518 21,33.518 21,36.518 23,36.518 23,33.518 24,33.518 24,31.518 12,31.518   "/><polygon points="34,33.518 43,33.518 43,36.518 45,36.518 45,33.518 46,33.518 46,31.518 34,31.518   "/></g></g></svg>
                </button>
              </form>
              {{dislikes}}
            </div>
            <div class="react">
              <form action="">
                <input type="text" value="comment" name="reaction" hidden>
                <input type="text" :value="postid" hidden>
                <button>
                  <svg fill="#000000" width="800px" height="800px" viewBox="0 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg"><title>comment-right</title><path d="M30.535 28.373c-1.809-1.73-3.119-3.968-3.7-6.485l-0.017-0.088c1.782-1.921 2.888-4.489 2.932-7.315l0-0.009c0-6.686-6.393-12.124-14.25-12.124s-14.25 5.438-14.25 12.124 6.393 12.125 14.25 12.125c0.004 0 0.009 0 0.014 0 1.883 0 3.691-0.319 5.374-0.906l-0.114 0.035c2.528 1.962 5.604 3.343 8.96 3.887l0.115 0.015c0.046 0.010 0.098 0.015 0.152 0.016h0c0.414-0 0.75-0.336 0.75-0.75 0-0.205-0.082-0.39-0.215-0.526l0 0zM21.426 24.348c-0.010-0.009-0.025-0.004-0.035-0.013-0.128-0.11-0.296-0.177-0.479-0.177h-0c-0.022 0-0.039 0.007-0.061 0.009-0.070 0.001-0.137 0.011-0.2 0.030l0.005-0.001c-1.516 0.574-3.269 0.906-5.099 0.906-0.020 0-0.040-0-0.060-0h0.003c-7.030 0-12.75-4.766-12.75-10.625s5.72-10.624 12.75-10.624c7.031 0 12.75 4.766 12.75 10.624-0.024 2.593-1.087 4.934-2.791 6.63l-0 0-0.010 0.026c-0.111 0.124-0.18 0.288-0.183 0.468v0.001c-0.001 0.015-0.002 0.033-0.002 0.050 0 0.007 0 0.014 0 0.022l-0-0.001c-0.002 0.017-0.002 0.037-0.002 0.058 0 0.008 0 0.016 0 0.024l-0-0.001c0.415 2.246 1.34 4.227 2.652 5.887l-0.021-0.028c-2.496-0.614-4.669-1.747-6.49-3.285l0.024 0.019z"/></svg>
                </button>
              </form>
              {{ comments.length }}
            </div>
          </div>

          <p v-if="isUpdated" class="edited">your post was successfully updated</p>
          <div class="overlay" v-if="edited"  @click.stop="handleClick"></div>
            <form action="" class="edited" v-if="edited">
              <h3 class="title">Edit Post</h3>
              <input type="text" placeholder="title" v-model="title" name="title" id="">
              <textarea laceholder="decription" v-model="description" name="decription" cols="30" rows="10"></textarea>
              <input type="text" placeholder="location" v-model="location" name="location">
              <input type="text" placeholder="date" v-model="date" name="date">
              <input type="button" value="edit" @click="updatePost"/>
            </form>
          <div action="" v-if="isdeleted" class="edited">
            <p>your post was successfully deleted</p>
            <RouterLink to="/">ok</RouterLink>
          </div>

          <div class="create-comment">
            <form>
              <textarea v-model="commentContent" name="text" cols="30" rows="10" placeholder="Write your comment here"></textarea>
              <button @click.prevent="createComment">Post Comment</button>
            </form>
          </div>
          <div class="comments">
            <div v-for="comment in comments" :key="comment.id" class="comment">
              {{ comment.id }}
              <div class="username">commented by <span>{{ comment.commentuserrole}}</span> @{{ comment.commentusername}}</div>
              <p class="text">{{ comment.text }}</p>
              <div class="reactions-to-comment">
                <button @click.prevent="likeComment(comment.id)" class="like-btn">{{ comment.like }} 👍</button>
                <button @click.prevent="dislikeComment(comment.id)" class="dislike-btn">{{ comment.dislike }} 👎</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        isLiked: false,
        isUser: false,
        post: {},
        comments: {},
        commentContent: '',
        errorMessage: '',
        isdeleted: false,
        edited: false,
        isUpdated: false,
        title: '',
        description: '',
        date: '',
        location: '',
        postid: null,
        userid: null,
        likes: null,
        dislikes: null,
        dislikeC: null,
        likeC: null,
        usernamepost: null,
        postuserrole: null
      };
    },
    async mounted() {
      const postId = this.$route.params.id;
      try {
        const postResponse = await axios.get(`http://localhost:8000/post?id=${postId}`);
        const commentsResponse = await axios.get(`http://localhost:8000/comments?postId=${postId}`);
        this.post = postResponse.data;
        this.comments = commentsResponse.data;
        // for post /************************************************************************ */
        this.postid = postResponse.data.id
        this.userid = postResponse.data.userId
        this.title = postResponse.data.title
        this.description = postResponse.data.description
        this.location = postResponse.data.location
        this.date = postResponse.data.date
        this.likes = postResponse.data.like
        this.dislikes = postResponse.data.dislike
        this.usernamepost = postResponse.data.postusername
        this.postuserrole = postResponse.data.postuserrole
        // for comment /************************************************************************ */
        // console.log(commentsResponse.data, " comment")
        console.log(postResponse.data, " post")
      } catch (error) {
        console.log(error);
        this.errorMessage = "Error loading post";
      }
    },
    methods: {
      async updatePost() {
        try {
          const response = await axios.post('http://localhost:8000/update-post', {
            Id: this.postid,
            UserId: this.userid,
            Title: this.title,
            Description: this.description,
            Date: this.date,
            Location: this.location
          },{
            withCredentials: true,
            headers: {
              'Content-Type': 'application/json'
            }
          })
          console.log(response.data.valid, " valid")
          console.log(response.data, " updated")
          console.log(response.data.isuser, " response.data.isUser")
          this.isUser = response.data.isuser
          this.isUpdated = response.data.isupdated

          if (response.status === 200 && response.data.valid) {
            // this.$router.push(`/post/${response.data.postid}`)
          } else{
            this.errorMess = response.data.err
          }
        } catch (error) {
          console.log(error)
        }
        this.edited = !this.edited
      },
      async deletePost() {
        try {
          const response = await axios.post('http://localhost:8000/delete-post', {
            Id: this.postid,
            UserId: this.userid
          }, {
            withCredentials: true,
            headers: {
              'Content-Type': 'application/json'
            }
          })
          console.log(response.data.deletedPost, " deleted post")
          this.isdeleted = response.data.isdeleted
          this.isUser = response.data.isuser
          if (response.status === 200 && response.data.valid) {
            // this.$router.push(`/post/${response.data.postid}`)
          } else{
            this.errorMess = response.data.err
          }
        } catch (error) {
          console.log(error)
        }
      },
      async postReaction(action) {
        try {
          const response = await axios.post('http://localhost:8000/post-reaction', {
            Id: this.postid,
            UserId: this.userid,
            ReactionType: action,
          },{
            withCredentials: true,
            headers: {
              'Content-Type': 'application/json'
            }
          })
          // console.log(response.data)
          // this.likePost = response.data.isliked
          if (response.status === 200 && response.data.valid) {
          } else{
            this.errorMess = response.data.err
          }
        } catch (error) {
          console.log(error)
        }
      },
      async createComment() {
        try { 
          const response = await axios.post('http://localhost:8000/create-comment', {
          Text: this.commentContent,
          UserId: this.userid,
          PostId: this.postid,
        },{
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        })
        console.log(response.data.valid)
        console.log(response.data)
        if (response.status === 200 && response.data.valid) {
          // this.$router.push(`/post/${response.data.postid}`)
        } else{
          this.errorMess = response.data.err
        }
      } catch (error) {
        console.log(error)
      }
      },
      async commentReaction(commentId, action) {
        try {
          const response = await axios.post('http://localhost:8000/comment-reaction', {
            Id: commentId,
            UserId: this.userid,
            ReactionType: action,
          },{
            withCredentials: true,
            headers: {
              'Content-Type': 'application/json'
            }
          })
          // console.log(response.data)
          // this.likePost = response.data.isliked
          if (response.status === 200 && response.data.valid) {
            
          } else{
            this.errorMess = response.data.err
          }
        } catch (error) {
          console.log(error)
        }
      },
      edit() {
        this.edited = !this.edited
      },
      handleClick() {
        this.edited = !this.edited
      },
      likePost() {
      this.postReaction('like')
      },
      dislikePost() {
        this.postReaction('dislike')
      },
      likeComment(commentId) {
      this.commentReaction(commentId,'like')
      },
      dislikeComment(commentId) {
        this.commentReaction(commentId,'dislike')
      },
    }
  };
  </script>

<style scoped>
.comments {
  margin-top: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.comment {
  margin-top: 20px;
  padding: 20px;
  background-color: #fff;
  border-radius: 10px;
  width: 100%;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
}

.comment:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.username {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}
.username span {
  color: #8593bd;}
.text {
  font-size: 14px;
  margin-bottom: 20px;
  line-height: 1.5;
}

.reactions-to-comment {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100px;
}

.like-btn, .dislike-btn {
  border: none;
  background-color: transparent;
  font-size: 14px;
  cursor: pointer;
  color: #888;
  transition: transform 0.2s, color 0.2s;
}

.like-btn:hover, .dislike-btn:hover {
  color: #555;
  transform: scale(1.1);
}

/*  */
.create-comment {
  margin-top: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.create-comment form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.create-comment textarea {
  border: none;
  padding: 15px;
  background-color: #f5f5f5;
  font-family: 'Roboto', sans-serif;
  font-size: 16px;
  border-radius: 5px;
  resize: none;
  height: 150px;
}

.create-comment textarea:focus {
  outline: none;
  box-shadow: 0 0 0 2px #2196f3;
}

.create-comment button {
  border: none;
  background-color: #2196f3;
  color: white;
  padding: 10px 16px;
  text-align: center;
  text-decoration: none;
  font-size: 16px;
  font-weight: bold;
  border-radius: 5px;
  cursor: pointer;
}

.create-comment button:hover {
  background-color: #1976d2;
}

/*  */

.reaction{
  display: flex;
  align-items: center;
  gap: 50px;
}
.react {  
  display: flex;
  align-items: center;
  transition: all 0.15s ease;
}
.react button {
  width: 30px;
  height: 30px;
  cursor: pointer;
  border: none;
  background: none;
}
.react svg path {
  transition: all 0.15s ease;
}
.react:hover svg path {
  fill: #5d54a4;
}
.react:hover {
  color: #5d54a4;
}
.react svg {
  width: 100%;
  height: 100%;
  object-fit: contain;
  flex-shrink: 0;
}
.edited {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 400px;
  padding: 30px;
  background: #8593bd;
  border-radius: 20px;
  text-align: center;
  z-index: 100;
}
input, textarea {
  background: transparent;
  border-radius: 10px;
  padding: 10px;
  border: 1px solid #0f2051;
  resize: vertical;
}
.overlay {
  width: 100vw;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;
  backdrop-filter: blur(10px);
}
.post__subconainer {
  display: flex;
  align-items: center;
  flex-direction: row-reverse;
  justify-content: space-between;
  margin: 20px 0;
}
.post__header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.edit__container {
  margin-left: 20px;
}
.edit__container .edit {
  margin-bottom: 30px;
}
.title {
  margin-bottom: 0 ;
}
.post__header {
  display: flex;
  flex-direction: column;
  /* align-items: center; */
  /* justify-content: space-between; */
}
.add__todo{
  display: flex;
  align-items: center;
}
.post__header {
  flex-wrap: wrap;
}
.icon, .edit {
  width: 30px;
  height: 30px;
  margin-right: 20px;
  cursor: pointer;
}
.icon svg, .icon img , .edit svg{
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.small {
  font-size: 14px;
  /* margin-bottom: 10px; */
}
.username {
  font-size: 14px;
  margin-top: 10px;
  text-align: right;
}
.img {
  width: 100%;
  height: 400px;
  border: 5px solid;

}
.img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.desc, .tags {
  margin-top: 20px;
}
.tags{
  display: flex;
  align-items: center;
  gap: 10px;
}
.tags span{
  border-radius: 10px;
  background: #8593bd;
  color: #0f2051;
  padding: 5px 10px;
}
.post {
  margin-bottom: 100px;
}
</style>
  