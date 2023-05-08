<template>
<div class="space"></div>
<div class="users__posts">
    <div class="container">
        <div class="posts">
            <div class="post" v-for="post in posts" :key="post.id">

                <!-- <p> title: {{  post.title }}</p> -->
                <a :href="'/post/' + post.id">
                  <p>title: {{ post.title }}</p>
                </a>
                <img :src="post.img" alt="">
                <p>created: {{ post.created }}</p>
                <p>desc: {{ post.description }}</p>
                <p>@{{ post.postusername }}</p>
                <p>role: {{  post.postuserrole }}</p>
                tags:
                <div class="tags" v-for="tag in post.tags">
                    {{ tag }}
                </div>
                <p>data: {{ post.date }}, location: {{ post.location }}</p>
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
            posts: []
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
}
</script>

<style scoped>
.post {
  border: 1px solid black;
  margin: 10px;
  padding: 10px;
}

.post img {
  max-width: 100%;
  height: auto;
  margin: 10px 0;
}

.tags {
  display: inline-block;
  margin-right: 5px;
  margin-bottom: 5px;
  padding: 2px 5px;
  background-color: lightgray;
  border-radius: 5px;
  font-size: 12px;
  font-weight: bold;
}

</style>