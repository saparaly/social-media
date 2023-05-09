<template>
    <div class="space"></div>
    <div class="container">
      <div class="users">
        <p>{{ mes }}</p>
        <div class="user" v-for="user in users">

          <div v-if="currectUser.following === null">
            <button @click="addUser(user.username)">follow</button>
          </div>
          <div v-else>
            <button v-if="currectUser.following.includes(user.id)" @click="removeUser(user.username)">unfollow</button>
            <button v-else @click="addUser(user.username)">follow</button>
          </div>


          <div class="username">@{{ user.username }} <span> {{ user.role }}</span></div>
          <div class="username">{{ user.email }}</div>
          <div class="username">followers: {{ user.followers}}</div>
          <div class="follow click" @click.prevent="getFolloingUsers(user.following, user.username)">
            following:
            {{ user.following.length > 0 ? (user.following[0] == 0 ? user.following.length - 1 : user.following.length) : 0 }}
          </div>

        </div>
      </div>
      <!-- {{ users }} -->
      <!-- {{ currectUser.following }} -->
      <div class="modal" v-if="followedUsers && followedUsers.length > 0">
        <h2>{{ username }} Follows </h2>
        <ul class="user-list">
          <li v-for="user in followedUsers" :key="user.id">
            <h3>{{ user.username }}</h3>
            <p>Role: {{ user.role }}</p>
            <p>Following: {{ user.following.length }}</p>
          </li>
        </ul>
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
.user{
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  border-radius: 10px;
  background: #8593bd;
  color: #5d54a4;
  padding: 10px;
}
.username span{
  font-weight: bold;
  color: #fff;
}
button {
  width: 80px;
  padding: 5px;
  background: transparent;
  border: 1px solid #5d54a4;
  border-radius: 10px;
  cursor: pointer;
  transition: 0.3s;
  color: #5d54a4;
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