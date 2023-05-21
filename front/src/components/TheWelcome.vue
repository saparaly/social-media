<template>
<!-- <div class="space"></div> -->
<div class="users__posts">
    <div class="container">
      <div class="users__posts-container">
        <div class="left wow slideInLeft">
        <div class="users">
          <p>recommended users:</p>
          <div class="user">
            <RouterLink to="/">Assel - <span>vue.js, golang</span></RouterLink>
            <RouterLink to="/">Sanzhar - <span>golang, python</span></RouterLink>
            <RouterLink to="/">Ruslan - <span>sql, flutter</span></RouterLink>
          </div>
          <p>search for user</p>
          <div class="group serch">
            <input type="text" name="" id="">
            <button>search</button>
          </div>
          <div class="user">
            <p>here you will see result</p>
          </div>
        </div>
        <div class="users fit">
          <RouterLink to="/create-post">create post</RouterLink>
          <RouterLink to="/userprof">my posts</RouterLink>
          <RouterLink to="/userprof">liked posts</RouterLink>
        </div>
        </div>
        <div class="posts" v-if="posts">
            <div  class="post" v-for="post in posts" :key="post.id">
                <div class="small">
                  created at {{ formatDate(post.created) }} by {{  post.postuserrole }} 
                  <RouterLink :to="'/user/' + post.userId" class="click">
                    <span>@{{ post.postusername }}</span>
                  </RouterLink>
                </div>
              <div class="img" v-if="post.img"><img :src="post.img" alt=""></div>
                <RouterLink :to="'/post/' + post.id" class="title"> {{ post.title }}</RouterLink>
                <p class="desc"> {{ post.description }}</p>
                  <div class="tags" v-if="post.tags.length > 1">
                    <div class="tag" v-for="tag in post.tags">
                      {{ tag }}
                    </div>
                  </div>
                  <div class="rightSide">
                  <p class="small flex" v-if=" post.date || post.location">
                    <span>
                      <svg version="1.0" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="800px" height="800px" viewBox="0 0 64 64" enable-background="new 0 0 64 64" xml:space="preserve">
                        <g>
                          <path fill="#394240" d="M31.998,2.478c0.279,0,0.463,0.509,0.463,0.509l8.806,18.759l20.729,3.167l-14.999,15.38l3.541,21.701   l-18.54-10.254l-18.54,10.254l3.541-21.701L2,24.912l20.729-3.167l8.798-18.743C31.527,3.002,31.719,2.478,31.998,2.478 M31.998,0   c-0.775,0-1.48,0.448-1.811,1.15l-8.815,18.778L1.698,22.935c-0.741,0.113-1.356,0.632-1.595,1.343   c-0.238,0.71-0.059,1.494,0.465,2.031l14.294,14.657L11.484,61.67c-0.124,0.756,0.195,1.517,0.822,1.957   c0.344,0.243,0.747,0.366,1.151,0.366c0.332,0,0.666-0.084,0.968-0.25l17.572-9.719l17.572,9.719   c0.302,0.166,0.636,0.25,0.968,0.25c0.404,0,0.808-0.123,1.151-0.366c0.627-0.44,0.946-1.201,0.822-1.957l-3.378-20.704   l14.294-14.657c0.523-0.537,0.703-1.321,0.465-2.031c-0.238-0.711-0.854-1.229-1.595-1.343l-19.674-3.006L33.809,1.15   C33.479,0.448,32.773,0,31.998,0L31.998,0z"/>
                          <path fill="#F9EBB2" d="M31.998,2.478c0.279,0,0.463,0.509,0.463,0.509l8.806,18.759l20.729,3.167l-14.999,15.38l3.541,21.701   l-18.54-10.254l-18.54,10.254l3.541-21.701L2,24.912l20.729-3.167l8.798-18.743C31.527,3.002,31.719,2.478,31.998,2.478"/>
                        </g>
                      </svg>
                    </span>
                    data: {{ post.date }}, location: {{ post.location }}
                  </p>
                  <div class="reaction">
                    <div class="react" :class="{liked: post.isliked }">
                  <form action="">
                    <button @click.prevent="likePost(post.id)">
                      <svg fill="#000000" height="800px" width="800px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 297.221 297.221" xml:space="preserve"><g><path d="M283.762,32.835c2.705-1.913,3.346-5.658,1.432-8.363c-1.914-2.705-5.657-3.347-8.363-1.432l-14.984,10.602   c-2.705,1.913-3.346,5.658-1.432,8.363c1.169,1.652,3.022,2.535,4.902,2.535c1.198,0,2.408-0.358,3.461-1.104L283.762,32.835z"/><path d="M244.064,29.387c0.695,0.262,1.409,0.386,2.11,0.386c2.428,0,4.713-1.484,5.617-3.891l6.46-17.182   c1.166-3.101-0.403-6.561-3.505-7.727c-3.101-1.167-6.562,0.404-7.728,3.505l-6.46,17.182   C239.393,24.761,240.962,28.221,244.064,29.387z"/><path d="M291.223,55.611c-0.041,0-0.082,0-0.124,0l-18.351,0.154c-3.313,0.067-5.944,2.605-5.877,5.918   c0.066,3.271,2.739,5.928,5.997,5.928c0.041,0,0.082,0,0.124,0l18.351-0.313c3.313-0.068,5.944-2.732,5.877-6.045   C297.154,57.982,294.481,55.611,291.223,55.611z"/><path d="M254.2,147.154c-3.073-2.433-6.711-4.089-10.557-4.867c0.254-0.46,0.491-0.928,0.715-1.403l2.408-2.408   c9.274-9.275,10.248-23.874,2.264-33.961c-3.769-4.761-9.001-7.925-14.812-9.106c0.415-0.764,0.783-1.545,1.117-2.338   c6.316-9.149,6.213-21.445-0.782-30.283c-3.77-4.764-9.004-7.938-14.818-9.117c4.8-8.826,4.187-19.826-2.225-27.925   c-4.848-6.125-12.109-9.639-19.923-9.639c-6.257,0-12.16,2.236-16.792,6.33c-0.701-3.979-2.363-7.822-5.012-11.169   c-4.849-6.125-12.11-9.638-19.924-9.639l0,0c-6.79,0-13.164,2.635-17.947,7.418l-60.84,60.84l-0.232-8.12   c-0.107-13.83-11.392-25.049-25.247-25.049c-13.604,0-24.729,10.815-25.229,24.298l-12.146,88.306l-9.983,11.604   c-5.983,6.957-5.582,17.481,0.915,23.962L19.987,199.7c-4.574,6.881-3.773,16.266,2.206,22.23l69.667,69.557   c3.329,3.321,7.748,5.148,12.446,5.148c3.857,0,7.668-1.295,10.729-3.645l14.544-11.168c13.991-3.305,29.416-10.813,45.874-22.33   c14.371-10.058,29.962-23.46,46.337-39.836l34.631-34.631c5.107-5.107,7.795-12.188,7.375-19.427   C263.376,158.371,259.879,151.649,254.2,147.154z M188.124,32.009c2.603-2.602,6.032-3.903,9.462-3.903   c3.915,0,7.831,1.695,10.515,5.086c4.256,5.377,3.51,13.18-1.339,18.028l-6.177,6.176c-0.952,0.635-1.879,1.314-2.747,2.083   c-0.701-3.98-2.364-7.823-5.013-11.169c-3.257-4.114-7.604-7.043-12.475-8.527L188.124,32.009z M146.397,17.532   c2.602-2.602,6.032-3.903,9.462-3.903c3.916,0.001,7.831,1.696,10.515,5.087c4.256,5.377,3.51,13.179-1.339,18.027l-70.919,70.186   l-0.233-8.119c-0.061-7.825-3.7-14.812-9.356-19.405L146.397,17.532z M13.624,176.391c-2.082-2.078-2.209-5.41-0.291-7.64   l12.281-14.277c0.006-0.007,0.011-0.017,0.012-0.026l12.72-92.483c0-7.286,5.961-13.247,13.247-13.247   c7.286,0,13.248,5.961,13.248,13.247L65.186,74c-11.988,1.646-21.322,11.733-21.78,24.057l-12.145,88.307l-3.533,4.108   L13.624,176.391z M247.935,176.539l-34.63,34.631c-29.577,29.577-60.494,53.318-87.653,59.237   c-0.825,0.181-1.601,0.528-2.271,1.043l-15.655,12.022c-1.014,0.779-2.219,1.162-3.419,1.162c-1.443,0-2.88-0.555-3.968-1.641   l-69.671-69.56c-2.083-2.077-2.21-5.409-0.291-7.64l12.28-14.276c0.007-0.008,0.011-0.017,0.013-0.026l12.719-92.483   c0-7.286,5.962-13.248,13.248-13.248c7.286,0,13.247,5.962,13.247,13.248l0.626,21.824c0.104,3.626,3.087,5.987,6.191,5.987   c1.514,0,3.058-0.563,4.309-1.813l70.431-70.431c2.603-2.603,6.031-3.903,9.462-3.903c3.915,0,7.831,1.695,10.515,5.086   c4.256,5.377,3.509,13.18-1.34,18.028l-48.518,48.518c-2.519,2.52-2.519,6.603,0,9.121l0,0c1.275,1.275,2.946,1.913,4.617,1.913   s3.343-0.638,4.617-1.913l62.374-62.373c2.602-2.603,6.031-3.903,9.462-3.903c3.915,0.001,7.831,1.696,10.515,5.087   c4.256,5.376,3.509,13.179-1.34,18.027l-62.081,62.081c-2.553,2.554-2.553,6.692,0,9.246c1.258,1.258,2.906,1.887,4.556,1.887   c1.648,0,3.297-0.629,4.555-1.887l48.811-48.81c2.603-2.603,6.032-3.903,9.462-3.903c3.915,0,7.831,1.695,10.515,5.087   c4.256,5.376,3.509,13.179-1.34,18.027l-48.349,48.35c-2.612,2.611-2.612,6.847,0,9.458l0.078,0.079   c1.207,1.207,2.789,1.81,4.37,1.81c1.582,0,3.164-0.603,4.37-1.81l29.974-29.974c2.701-2.701,6.317-4.129,9.921-4.129   c2.867,0,5.726,0.904,8.107,2.789C253.114,161.598,253.508,170.967,247.935,176.539z"/></g></svg>
                    </button>
                  </form>
                  {{post.like}} 
                </div>
                <div class="react" :class="{liked: post.isdisliked }">
                  <form action="">
                    <button @click.prevent="dislikePost(post.id)">
                      <svg fill="#000000" height="800px" width="800px" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512" xml:space="preserve"><g><g><g><path d="M256,0C114.617,0,0,114.617,0,256s114.617,256,256,256s256-114.617,256-256S397.383,0,256,0z M256,494.933     C124.043,494.933,17.067,387.957,17.067,256S124.043,17.067,256,17.067S494.933,124.043,494.933,256S387.957,494.933,256,494.933     z"/><path d="M332.8,324.267H179.2c-4.713,0-8.533,3.821-8.533,8.533c0,4.713,3.82,8.533,8.533,8.533h153.6     c4.713,0,8.533-3.821,8.533-8.533C341.333,328.087,337.513,324.267,332.8,324.267z"/><path d="M196.267,213.333c0-9.606-3.176-18.469-8.534-25.6H204.8c4.713,0,8.533-3.82,8.533-8.533s-3.821-8.533-8.533-8.533H102.4     c-4.713,0-8.533,3.82-8.533,8.533s3.82,8.533,8.533,8.533h17.067c-5.357,7.131-8.534,15.995-8.534,25.6     c0,23.563,19.104,42.667,42.667,42.667C177.163,256,196.267,236.896,196.267,213.333z M153.6,238.933     c-14.137,0-25.6-11.463-25.6-25.6c0-14.137,11.463-25.6,25.6-25.6c14.137,0,25.6,11.463,25.6,25.6     C179.2,227.471,167.737,238.933,153.6,238.933z"/><path d="M409.6,170.667H307.2c-4.713,0-8.533,3.82-8.533,8.533s3.82,8.533,8.533,8.533h17.067     c-5.357,7.131-8.534,15.995-8.534,25.6c0,23.563,19.104,42.667,42.667,42.667c23.563,0,42.667-19.104,42.667-42.667     c0-9.606-3.176-18.469-8.534-25.6H409.6c4.713,0,8.533-3.82,8.533-8.533S414.313,170.667,409.6,170.667z M384,213.333     c0,14.137-11.463,25.6-25.6,25.6s-25.6-11.463-25.6-25.6c0-14.137,11.463-25.6,25.6-25.6S384,199.196,384,213.333z"/></g></g></g></svg>                
                      </button>
                  </form>
                  {{post.dislike}}
                </div>
                  </div>
                </div>
            </div>
        </div>
        <div class="posts" v-else>
          no post
        </div>
        <div class="right">
        <div class="filters">
          <div>
            <p>filter posts by popular tag</p>
            <div class="items">
              <div class="group">aitu</div>
              <div class="group">prototypes</div>
              <div class="group">javascript</div>
              <div class="group">golang</div>
              <div class="group">typescript</div>
            </div>
          </div>

          <div>
            <p>search by tag</p>
            <div class="group serch">
              <input type="text" name="" id="">
              <button>search</button>
            </div>
          </div>

          <div>
            <p>filter post by date</p>
            <div class="items">
              <div class="group">lates</div>
              <div class="group">oldest</div>
            </div>
          </div>

          <div>
            <p>filter post by star</p>
            <div class="group">post with start</div>
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
    name: 'Home',
    data() {
        return {
            posts: [],
            id: 0
        }
    },
    async mounted() {
      try {
        const userResponse = await axios.get(`http://localhost:8000/home`, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        });
        this.posts = userResponse.data;
        console.log(userResponse.data, " users")
      } catch (error) {
        console.log(error); 
        this.errorMessage = "Error loading users post";
      }
    },
    methods: {
      async postReaction(action, postid,) {
        try {
          const response = await axios.post('http://localhost:8000/post-reaction', {
            Id: postid,
            UserId: this.userid,
            ReactionType: action,
          },{
            withCredentials: true,
            headers: {
              'Content-Type': 'application/json'
            }
          })
          if (response.status === 200 && response.data.valid) {
          } else{
            this.errorMess = response.data.err
          }
        } catch (error) {
          console.log(error)
        }
      },
      likePost(postid) {
      this.postReaction('like', postid)
      },
      dislikePost(postid) {
        this.postReaction('dislike', postid)
      },
      formatDate(dateString) {
        const date = new Date(dateString);
        const day = date.getDate().toString().padStart(2, '0');
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const year = date.getFullYear().toString().slice(2);
        const hours = date.getHours().toString().padStart(2, '0');
        const minutes = date.getMinutes().toString().padStart(2, '0');
        const seconds = date.getSeconds().toString().padStart(2, '0');

        return `${day}.${month}.${year} ${hours}:${minutes}:${seconds}`;
      }
    }
}
</script>

<style scoped>

.filters {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.items {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-gap: 10px;
  margin-bottom: 20px;
}
.items .group:nth-child(3n) {
grid-column: span 2 / auto;
}
.filters .group {
  /* background: var(--primary-color);
  color: #fff; */
  background: var(--bg);
  color: var(--text-color);
  border: 1px solid  var(--text-color);
  padding: 6px;
  border-radius: 10px;
  cursor: pointer;
  transition: 0.2s;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
}
.filters .group.serch{
  background: none;
  padding: 0;
  margin-bottom: 20px;
}
.user p {
  margin-top: 20px;
}
.group {
  display: flex;
  align-items: center;
  width: 100%;
}
.serch{
  border: 1px solid #676767;
  border-radius: 10px;
  font-size: 18px;
  letter-spacing: 0.09em;
  color: #4F555A;
  background: transparent;
  width: 100%;
}
.serch input{
  width: 100%;
  border: none;
  padding: 5px 10px;
  background: transparent;
}
.serch button {
  background: var(--primary-color);
  color: #fff;
  font-weight: 600;
  padding: 6px 10px;
  border: none;
  border-radius: 10px;
  margin-top: -1px;
  margin-right: -1px;
  margin-bottom: -1px;
  cursor: pointer;
  transition: 0.2s;
  border: 1px solid var(--primary-color);
}
.serch button:hover {
  background: #fff;
  color: var(--primary-color);
  border: 1px solid #676767;
}
.users p, .filters p{
  font-size: 16px;
  letter-spacing: 0.09em;
  color: #4F555A;
  margin-bottom: 20px;
}
.user {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 20px;
}
.user a {
  color: var( --dark);

  font-weight: bold;
  transition: 0.2s;
}
.user a:hover {
  color: var(--pink);
}
.user a span {
  font-weight: 500;
}
.group {

}
.users__posts-container{
  display: flex;
  justify-content: space-between;
  gap: 50px;
}
.filters, .users {
  width: 300px;
  flex-shrink: 0;
  background: #fff;
  border-radius: 20px;
  padding: 30px;
  height: 500px;
  margin-bottom: 20px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04);
}
.users {
  height: 400px;
}
.users.fit a{
  display: inline-block;
  padding: 10px 20px;
  border-radius:10px ;
  text-align: center;
  /* background: var(--primary-color);
  color: #fff; */
  background: var(--bg);
color: var(--text-color);
border: 1px solid var(--text-color);
}
.users.fit {
  display: flex;
  flex-direction: column;
  gap: 20px;
}
.fit{

  height: fit-content;
}
.users__posts {
  margin-bottom: 100px;
}
.click {
  cursor: pointer;
}
/* .posts {
  display: grid;
  grid-template-columns: repeat(2,1fr);
  grid-gap: 20px;
} */
.post {
  background: #fff;
  padding: 20px;
  border-radius: 20px;
  margin-bottom: 20px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04);
}

.img{
  width: 100%;
  height: 400px;
  border-radius: 20px;
  overflow: hidden;
  flex-shrink: 0;
  margin: 10px 0;
}
.img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.rightSide{
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.small {

  color: #4F555A;
  font-size: 14px;
}
.small span {
  font-weight: bold;
  color: var(--dark);
}
.flex {
  display: flex;
  align-items: center;
  gap: 5px;
}
.small svg {
  width: 20px;
  height: 20px;
}
.title{
  font-size: 40px;
  text-transform: capitalize;
  margin-bottom: 15px;
  transition: all 0.15s;
  display: inline-block;
}
.title:hover {
  color: var(--primary-color);
}

.desc {
  font-size: 16px;
  margin-bottom: 20px;
}
.tags {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 20px;
  margin-bottom: 20px;
}
.tag{
  background: var(--bg);
  color: var(--text-color);
  border-radius: 10px;
  padding: 5px 10px;
}
.reaction{
  display: flex;
  align-items: center;
  gap: 30px;
  margin-left: auto;
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
.react:hover svg path, .liked  svg path{
  fill: var(--dark);
}
.react:hover, .liked {
  color: var(--dark);
}
.react svg {
  width: 100%;
  height: 100%;
  object-fit: contain;
  flex-shrink: 0;
}
</style>