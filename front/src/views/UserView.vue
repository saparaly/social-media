<template>
    <div class="space"></div>
    <div class="container">
      <div class="users">
        <p>{{ mes }}</p>
        <div class="user" v-for="user in users">
          <button @click="addUser(user.username)">follow</button>
          <button @click="removeUser(user.username)">unfollow</button>
          <div class="user">@{{ user.username }}</div>
          <div class="user">{{ user.email }}</div>
          <div class="user">{{ user.role }}</div>
          <div class="user">followers: {{ user.followers}}</div>
          <div class="user">following: {{ user.following }}</div>
        </div>
        {{ userFolloing }}
      </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Users',
    data() {
        return {
            users: [],
            userFolloing:[],
            mes: '',
            isFolloing: false
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
        this.users = usersResponse.data.users;
        this.userFolloing = usersResponse.data.users.following
        console.log(usersResponse.data.users, " users")
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
      }
    }
}
</script>

<style scoped>
.user {
  width: 100%;
  display: flex;
  align-items: center;
  margin: 10px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
}
.user button {
  padding: 5px 10px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.user .user {
  font-weight: bold;
  margin-bottom: 5px;
}

.user .user:not(:last-child) {
  margin-bottom: 3px;
  color: #666;
}

.user .user:last-child {
  margin-top: 5px;
  font-size: 12px;
  color: #888;
}
</style>