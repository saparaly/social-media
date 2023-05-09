<template>
<div class="nav">
  <!-- <div class="container"> -->
    <div class="nav__container">
      <RouterLink to="/">Home</RouterLink>

      <div v-if="isAuthenticated">
        <RouterLink to="/create-post">create post</RouterLink>
        <RouterLink to="/users">user list</RouterLink>
          <RouterLink to="/userProf">profile</RouterLink>
          <!-- <RouterLink to="/todos">todos</RouterLink>
          <RouterLink to="/chat">chat</RouterLink>
          <RouterLink to="/post/">my post</RouterLink> -->
          <RouterLink class="last" to="/logout" @click.prevent="logout">logout</RouterLink>
        </div>

        <div v-else>
          <RouterLink to="/signup">signup</RouterLink>
          <RouterLink to="/signin">signin</RouterLink>
        </div>
    </div>
  <!-- </div> -->
</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'NavBar',

  computed: {
    isAuthenticated() {
      const token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
      return token !== "";
    }

  },
  methods: {
    async logout() {
      try {
        const response = await axios.post('http://localhost:8000/logout', null, {
          withCredentials: true,
        })
        console.log(response.data)
        this.$router.push('/')
      } catch (error) {
        console.log(error)
      }
    },
  },
}
</script>

<style scoped>
.nav {
  position: fixed;
  top: 20%;
  left: 20%;
  width: 200px;
  background: #8294C4;
  padding: 15px 0;
  color: #fff;
  border-radius: 20px;
}
.nav__container div, .nav__container{
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 40px;
  height: 500px;
}
.last {
  margin-top: auto;
}
</style>