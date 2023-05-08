<template>
  <div class="space"></div>
  <div class="container">
    <div class="userProg">userProf</div>
    user:
    {{ user }}
    post:
    {{ posts }}
    like:
    {{ likedposts }}
  </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Users',
    data() {
        return { 
          user: {},
          posts: [],
          likedposts: [],
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
    }
}
</script>

<style>

</style>