<template>
  <!-- <div class="space"></div> -->
  <div class="container">

    <div class="user flex">
      <p class="flex">
        <div class="ava">
          <img :src="user.avaimg" alt="">
        </div>
        username: @{{ user.username }} <span class="bold">{{ user.role }}</span></p>
      <p> {{ user.email }}</p>
    </div>

    <div class="btns">
    <div class="follow click">
          followers:
          {{ user.followers && user.followers.length > 0 ? (user.followers[0] == 0 ? user.followers.length - 1 : user.followers.length) : 0 }}
    </div>
        <div class="follow click">
          following:
          {{ user.following && user.following.length > 0 ? (user.following[0] == 0 ? user.following.length - 1 : user.following.length) : 0 }}
        </div>
      <div class="btn"  v-if="user.role === 'student'" @click="showGraph" :class="{active : isgraph}">graph</div>
      <div class="btn" @click="showCreatedpost" :class="{active : iscreated}">created posts</div>
      <div class="btn" @click="showLikedPost" :class="{active : isliked}">likes posts</div>
    </div>
    
      
    <div class="posts" v-if="iscreated" >
      <div class="post" v-for="post in posts" :key="post.id">
              <div class="img" v-if="post.img"><img :src="post.img" alt=""></div>
                <RouterLink :to="'/post/' + post.id" class="title"> {{ post.title }}</RouterLink>
                <p class="desc"> {{ post.description }}</p>
                  <div class="tags" v-if="post.tags.length > 1">
                    <div class="tag" v-for="tag in post.tags">
                      {{ tag }}
                    </div>
                  </div>
            </div>
    </div>
      
    <div class="posts" v-if="isliked" >
      <div class="post" v-for="post in likedposts" :key="post.id">
              <div class="img" v-if="post.img"><img :src="post.img" alt=""></div>
                <RouterLink :to="'/post/' + post.id" class="title"> {{ post.title }}</RouterLink>
                <p class="desc"> {{ post.description }}</p>
                  <div class="tags" v-if="post.tags.length > 1">
                    <div class="tag" v-for="tag in post.tags">
                      {{ tag }}
                    </div>
                  </div>
            </div>
    </div>

    <div class="graph" v-if="isgraph" >
      <GraphComponent/>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import GraphComponent from '../components/GraphComponent.vue';

export default {
    name: 'Users',
    components: {GraphComponent},
    data() {
        return { 
          user: {},
          posts: [],
          likedposts: [],
          iscreated: true,
          isliked: false,
          isgraph: false,

        }
    },
    async mounted() {
      try {
        const usersResponse = await axios.get(`http://localhost:8000/profile`, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        });
        this.user = usersResponse.data.user
        this.posts = usersResponse.data.createdpost
        this.likedposts = usersResponse.data.likedpost
      } catch (error) {
        console.log(error); 
        this.errorMessage = "Error loading users";
      }
    },
    methods: {
      showCreatedpost() {
        this.iscreated = true
        this.isliked = false
        this.isgraph = false

      },
      showLikedPost() {
        this.isliked = true
        this.iscreated = false
        this.isgraph = false
      },
      showGraph() {
        this.isgraph = true
        this.iscreated = false
        this.isliked = false

      }
    }
}
</script>

<style scoped>
.ava {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
}
.ava img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user {
  color: #fff;
  display: flex;
  justify-content: space-between;
}
.bold {
  font-weight: bold;
}
.btns{
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 30px 0 200px 0;
  color: #fff;
  border-radius: 20px;
  border: 1px solid;
  padding: 10px 20px;
}
.btn.active {
/* border: 1px solid #fff; */
color: var(--dark);
background:#fff;
font-weight: 500;
padding: 6px;
border-radius: 10px;
}
.btn {
  cursor: pointer;
}
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