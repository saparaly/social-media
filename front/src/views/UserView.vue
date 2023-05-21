<template>
  <!-- <div class="space"></div> -->
  <div class="container">
    <div class="search">
      <input type="text">
      <button>search</button>
    </div>
    <div class="container__two">
      <div class="users">
        <p>{{ mes }}</p>
        <div class="user" v-for="user in users" :key="user.id">

          <div v-if="currectUser.following === null">
            <button @click="addUser(user.username)">follow</button>
          </div>
          <div v-else>
            <button v-if="currectUser.following.includes(user.id)" @click="removeUser(user.username)">unfollow</button>
            <button v-else @click="addUser(user.username)">follow</button>
          </div>


          <div class="username">@{{ user.username }} <span> {{ user.role }}</span></div>
          <div class="username">{{ user.email }}</div>
          <!-- <div class="username">
            followers: {{ user.followers}}

          </div> -->
          <div class="follow click" @click.prevent="getFolloingUsers(user.followers, user.username)">
            followers:
            {{ user.followers && user.followers.length > 0 ? (user.followers[0] == 0 ? user.followers.length - 1 : user.followers.length) : 0 }}
          </div>
          <div class="follow click" @click.prevent="getFolloingUsers(user.following, user.username)">
            following:
            {{ user.following && user.following.length > 0 ? (user.following[0] == 0 ? user.following.length - 1 : user.following.length) : 0 }}
          </div>


        </div>
      </div>
      <div class="list">
        <div class="modal" v-if="followedUsers && followedUsers.length > 0">
      <h2>{{ username }} follows/follower: </h2>
      <ul class="user-list">
        <li v-for="user in followedUsers" :key="user.id">
          <p>@{{ user.username }} <span>{{ user.role }}</span></p>
          <div class="follow click" @click.prevent="getFolloingUsers(user.followers, user.username)">
          followers:
          {{ user.followers && user.followers.length > 0 ? (user.followers[0] == 0 ? user.followers.length - 1 : user.followers.length) : 0 }}
        </div>
        <div class="follow click" @click.prevent="getFolloingUsers(user.following, user.username)">
          following:
          {{ user.following && user.following.length > 0 ? (user.following[0] == 0 ? user.following.length - 1 : user.following.length) : 0 }}
        </div>
        </li>
      </ul>
    </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Users',
  data() {
      return { 
        currectUserId: null,
        currectUser: {},
        users: [],
        userFollowing: [],
        mes: '',
        isFollowing: false,
        followedUsers: [],
        username: '',
      }
  },
  async mounted() {
    try {
      const usersResponse = await axios.get(`http://localhost:8000/users`, {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json'
        }
      });
      console.log(usersResponse.data, " users")
      this.users = usersResponse.data.users;
      this.currectUserId = usersResponse.data.id
      this.currectUser = this.users.find(user => user.id === this.currectUserId);
    } catch (error) {
      console.log(error); 
      this.errorMessage = "Error loading users";
    }
  },
  methods: {
    async addUser(username) {
      try {
        const response = await axios.post('http://localhost:8000/user-add', {
          Username: username
        },{
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        })
        console.log(response.data.valid)
        console.log(response.data)
        if (response.status === 200 && response.data.valid) {
          this.mes = `now you follow ${username}`
          
        } else{
          this.errorMess = response.data.err
        }
      } catch (error) {
        console.log(error)
      }
    },
    async removeUser(username) {
      try {
        const response = await axios.post('http://localhost:8000/user-remove', {
          Username: username
        },{
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        })
        console.log(response.data.valid)
        console.log(response.data)
        if (response.status === 200 && response.data.valid) {
          this.mes = `now you unfollow ${username}`
          
        } else{
          this.errorMess = response.data.err
        }
      } catch (error) {
        console.log(error)
      }
    },
    async getFolloingUsers(userFollowing, username) {
      try {
        const response = await axios.post('http://localhost:8000/followed-users', {
          Following: userFollowing
        },{
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        })
        // console.log(response.data.valid)
        console.log(response.data, " folloewd users")
        if (response.status === 200 && response.data.valid) {
          this.followedUsers =  response.data.users   
          this.username =   username      
        } else{
          this.errorMess = response.data.err
        }
      } catch (error) {
        console.log(error)
      }
    }
  }
}
</script>

<style scoped>
.list li {
  color: #38416f;
  background: #fff;
  border-radius: 10px;
  padding: 10px;
  margin-bottom: 10px;
}
h2{
  margin-bottom: 10px;
}
.search{
  display: flex;
  width: 100%;
  margin-bottom: 30px;
  background: #fff;
  padding: 10px;
  border-radius: 10px;
  color: #38416f;
}
.search button{
  color: #38416f;
  border: 1px solid;
}
input{
  border: none;
  width: 100%;
}
.container__two {
  display: flex;
  justify-content: space-between;
}
.list {
  width: 300px;
  border-radius: 10px;
  background: #38416f;
  color: #fff;
  padding: 10px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04);
  border: 1px solid;
  margin-left: 50px;
}
.user{
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  border-radius: 10px;
  background: #38416f;
  color: #fff;
  padding: 10px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04);
  border: 1px solid;
  width: 100%;
}
.users {
  width: 100%;
}
.user:last-child{
  margin-bottom: 0;
}
.username span{
font-weight: bold;
color: #fff;
}
button {
width: 80px;
padding: 5px;
background: transparent;
border: 1px solid #fff;
border-radius: 10px;
cursor: pointer;
transition: 0.3s;
color: #fff;
}
button:hover {
transform: scale(1.1);
}
.none {
display: none;
}
.click {
cursor: pointer;
}
</style>