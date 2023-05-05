<template>
<div class="space"></div>
<div class="users__posts">
    <div class="container">
        <div class="posts">
            <div class="post" v-for="post in posts" :key="post.id">

                <p> title: {{  post.title }}</p>
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

<style>
</style>