<template>
<div :class="['nav', { 'hidden': shouldHideBackground }]">
  <!-- <div class="bg"></div> -->
  <div :class="['bg', { 'hidden': shouldHideBackground }]"></div>
    <div class="container">
    <div class="nav__container">
      <RouterLink to="/" class="logo">AITUSTACK</RouterLink>
      <div v-if="isAuthenticated">
        <RouterLink to="/" exact-active-class="active">Home</RouterLink>
        <RouterLink to="/users" exact-active-class="active">Users</RouterLink>
        <RouterLink to="/userProf" exact-active-class="active">Profile</RouterLink>
        <RouterLink to="/chat">Chat</RouterLink>
        <RouterLink class="last" to="/logout" @click.prevent="logout">Logout</RouterLink>
      </div>

        <div v-else>
          <RouterLink to="/signup" exact-active-class="active">Signup</RouterLink>
          <RouterLink to="/signin" exact-active-class="active">Signin</RouterLink>
        </div>
    </div>
  </div>
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
    },
    shouldHideBackground() {
      const currentRoute = this.$route.path;
      const pagesToHideBackground = ['/signup', '/signin', '/create-post', '/chat'];
      return pagesToHideBackground.includes(currentRoute);
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
/* .nav.hidden a{
  color: #3e404b;
} */
  .bg.hidden {
    /* display: none; */
    height: 132px;
  }
.active {
  font-weight: bold;
  position: relative;
}
.active::before{
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #fff;
}

.nav.hidden a.active::before{
  background: #3e404b;
}
.logo {
  display: inline;
  color: #fff;
  font-size: 48px;
  /* animation: anime2 .5s alternate infinite ease-in */
}
@keyframes anime {
  0% {
    transform: translateY(0px);
  }
  30% {
    transform: translateY(-40px);
  }
  60% {
    transform: translateY(0px);
  }
  80% {
    transform: scaleX(1.5) scaleY(0.7) translateY(40px);
  }
  100% {
    transform: scaleX(1) scaleY(1) translanteY(0px);
  }
}

@keyframes anime2 {
  0% {
    transform: scaleX(1.5) scaleY(0.7) translateY(40px);
  }
  30% {
    transform: scaleX(1.2) scaleY(0.8) translateY(20px);
  }
  80% {
    transform: scaleX(1) scaleY(1) translateY(-30px);
  }
  100% {
    transform: scaleX(1) scaleY(1) translateY(-40px);
  }
}
.logo {
  margin-right: auto;
  font-weight: bold;
  font-size: 25px;
}
.bg {
  background: var( --dark);
height: 400px;
position: absolute;
top: 0;
left: 0;
width: 100%;
z-index: -1;
}
.nav {
  /* display: none; */
  /* position: fixed;
  top: 20%;
  left: 20%;
  width: 200px;
  background: var(--primary-color);
  padding: 15px 0;
  color: #fff;
  border-radius: 20px; */
  width: 100%;
  /* background: #333; */
  margin-bottom: 50px;
  margin-top: 30px;
  padding: 10px 0;
}
.nav__container div, .nav__container{
  display: flex;
  align-items: center;
  justify-content: right;
  gap: 40px;
}
.last {
  margin-top: auto;
}

a{
  transition: 0.3s;
  color: #fff;
}
a:hover {
  color: #646b95;
  }
</style>