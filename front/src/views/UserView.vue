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


          <div class="user">@{{ user.username }}</div>
          <div class="user">{{ user.email }}</div>
          <div class="user">{{ user.role }}</div>
          <div class="user none">followers: {{ user.followers}}</div>
          <div class="user click" @click.prevent="getFolloingUsers(user.following, user.username)">
            following:
            {{ user.following.length > 0 ? (user.following[0] == 0 ? user.following.length - 1 : user.following.length) : 0 }}
          </div>

        </div>
      </div>
      <!-- {{ users }} -->
      <!-- {{ currectUser.following }} -->
      <div class="modal" v-if="followedUsers && followedUsers.length > 0">
        <h2>{{ username }} follows: </h2>
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
        // console.log(usersResponse.data, " users")
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
.modal {
  position: fixed;
  top: 30%;
  right: 17%;
    width: 400px;
    height: 300px;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    flex-direction: column;
    overflow: auto;
    padding: 20px;
    border-radius: 20px;
  }

  .user-list {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .user-list li {
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 10px;
  }

  .user-list li h3 {
    margin: 0;
  }
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

.none {
  display: none;
}
.click {
  cursor: pointer;
}
</style>